package market

import (
    "time"
    "container/list"
    "math"
    "fmt"
    "github.com/roydong/gmvc"
    "log"
)


const (
    OrderStatusCreated   = 0
    OrderStatusPartial   = 1
    OrderStatusComplete  = 2
    OrderStatusCancel    = -1
    OrderStatusCanceling = 4

    StateOpen  = iota
    StateOpenPending
    StateClose
    StateClosePending
)

type HedgerWS struct {
    zuo *OKFutureWS
    you *OKFutureWS

    short *OKFutureWS
    long *OKFutureWS

    shortPrice, longPrice float64
    shortOrder, longOrder Order

    trade chan Trade
    lastTrade chan Trade
    lastOrder chan Order
    currentOrders map[int64]Order

    tradeAmount float64
    tradePriceOffset float64
    minTradeMargin float64

    minMargin     float64
    minMarginTime int64

    maxMargin     float64
    maxMarginTime int64

    totalMargin   float64
    avgMargin     float64

    tickerNum     int
    margins       map[int64]float64
    marginList    *list.List

    stoped        bool
    state         int
    pendingTime   int64

    started       time.Time
    tradeNum      int

    tcny        float64
    cny         float64
    btc         float64
}


func NewHedgerWS(zuo, you *OKFutureWS) *HedgerWS {
    hg := &HedgerWS{
        zuo: zuo,
        you: you,

        trade: make(chan Trade, 5),
        lastTrade: make(chan Trade, 5),
        lastOrder: make(chan Order, 10),
        currentOrders: make(map[int64]Order, 10),

        minMargin: math.Inf(1),
        maxMargin: math.Inf(-1),

        tickerNum: 200,
        margins: make(map[int64]float64),
        marginList: list.New(),

        state: StateClose,
        stoped: false,
    }

    conf := gmvc.Store.Tree("config.hedger")
    hg.minTradeMargin,   _ = conf.Float("min_trade_margin")
    hg.tradeAmount,      _ = conf.Float("trade_amount")
    hg.tradePriceOffset, _ = conf.Float("trade_price_offset")

    hg.zuo.AddHandler("trade", hg.syncTrade)
    hg.you.AddHandler("trade", hg.syncTrade)
    hg.zuo.AddHandler("lastTrade", hg.syncLastTrade)
    hg.you.AddHandler("lastTrade", hg.syncLastTrade)
    hg.zuo.AddHandler("order", hg.syncOrder)
    hg.you.AddHandler("order", hg.syncOrder)

    return hg
}

func (hg *HedgerWS) Start() {
    hg.stoped = false
    hg.tradeNum = 0
    hg.started = time.Now()

    go hg.updateMargins()
    go hg.arbitrage()
    //go hg.checkPending()

    gmvc.Logger.Println("started...")
}

func (hg *HedgerWS) syncOrder(args ...interface{}) {
    order, _ := args[0].(Order)
    hg.currentOrders[order.Id] = order
    hg.lastOrder <-order
}

func (hg *HedgerWS) syncTrade(args ...interface{}) {
    trade, _ := args[0].(Trade)
    if len(hg.trade) > cap(hg.trade) {
        <-hg.trade
    }
    hg.trade <-trade
}

func (hg *HedgerWS) syncLastTrade(args ...interface{}) {
    trade, _ := args[0].(Trade)
    if len(hg.lastTrade) > cap(hg.lastTrade) {
        <-hg.lastTrade
    }
    hg.lastTrade <-trade
}

func (hg *HedgerWS) Stop() {
    hg.stoped = true
}

func (hg *HedgerWS) updateMargins() {
    for !hg.stoped {
        trade := <-hg.trade
        idx := trade.No
        zuoPrice := hg.zuo.lastTrade.Price
        youPrice := hg.you.lastTrade.Price
        if zuoPrice <= 0 || youPrice <= 0 {
            continue
        }

        margin := youPrice - zuoPrice
        hg.totalMargin += margin
        hg.margins[idx] = margin
        hg.marginList.PushFront(idx)
        if hg.marginList.Len() > hg.tickerNum {
            el := hg.marginList.Back()
            hg.marginList.Remove(el)
            i, _ := el.Value.(int64)
            hg.totalMargin -= hg.margins[i]
            delete(hg.margins, i)

            if i == hg.minMarginTime {
                hg.minMarginTime, hg.minMargin = hg.getMinMargin()
            }
            if i == hg.maxMarginTime {
                hg.maxMarginTime, hg.maxMargin = hg.getMaxMargin()
            }
        }

        hg.avgMargin = hg.totalMargin / float64(hg.marginList.Len())
        if hg.minMargin > margin {
            hg.minMargin = margin
            hg.minMarginTime = idx
        }
        if hg.maxMargin < margin {
            hg.maxMargin = margin
            hg.maxMarginTime = idx
        }

        log.Println(fmt.Sprintf("%.3f <= %.3f(%.3f) => %.3f",
            hg.minMargin - hg.avgMargin, hg.avgMargin, margin, hg.maxMargin - hg.avgMargin))
    }
}

func (hg *HedgerWS) getMinMargin() (int64, float64) {
    min := math.Inf(1)
    var idx int64
    for el := hg.marginList.Back(); el != nil; el = el.Prev() {
        i, _ := el.Value.(int64)
        v := hg.margins[i]
        if v < min {
            idx = i
            min = v
        }
    }
    return idx, min
}

func (hg *HedgerWS) getMaxMargin() (int64, float64) {
    max := math.Inf(-1)
    var idx int64
    for el := hg.marginList.Back(); el != nil; el = el.Prev() {
        i, _ := el.Value.(int64)
        v := hg.margins[i]
        if v > max {
            idx = i
            max = v
        }
    }
    return idx, max
}


func (hg *HedgerWS) arbitrage() {
    for !hg.stoped {
        <-hg.lastTrade
        if hg.marginList.Len() < 50 {
            continue
        }
        if len(hg.zuo.lastAsks) == 0 {
            gmvc.Logger.Println(hg.zuo.Name() + " depth is empty")
            continue
        }
        if len(hg.you.lastAsks) == 0 {
            gmvc.Logger.Println(hg.you.Name() + " depth is empty")
            continue
        }

        zuoPrice := hg.zuo.lastTrade.Price
        youPrice := hg.you.lastTrade.Price
        margin := youPrice - zuoPrice
        log.Println(margin)
        if hg.state == StateClose {

            //满足最小差价条件,并且超过最大差价
            if margin - hg.avgMargin >= hg.minTradeMargin && margin >= hg.maxMargin {
                hg.state = StateOpenPending
                hg.pendingTime = time.Now().Unix()

                hg.short = hg.you
                hg.shortPrice = youPrice

                hg.long = hg.zuo
                hg.longPrice = zuoPrice

                go hg.short.Trade(TypeOpenShort, hg.tradeAmount, hg.shortPrice - hg.tradePriceOffset)
                go hg.long.Trade(TypeOpenLong, hg.tradeAmount, hg.longPrice + hg.tradePriceOffset)

                continue
            }

            //满足最小差价条件,并且低于最小差价
            if hg.avgMargin - margin >= hg.minTradeMargin && margin <= hg.minMargin {
                hg.state = StateOpenPending
                hg.pendingTime = time.Now().Unix()

                hg.short = hg.zuo
                hg.long = hg.you

                hg.shortPrice = zuoPrice
                hg.longPrice = youPrice

                go hg.short.Trade(TypeOpenShort, hg.tradeAmount, hg.shortPrice - hg.tradePriceOffset)
                go hg.long.Trade(TypeOpenLong, hg.tradeAmount, hg.longPrice + hg.tradePriceOffset)

                continue
            }

        } else if hg.state == StateOpen {

            //如果是右手做空
            if (hg.short.Name() == hg.you.Name()) {

                //差价低于平均差价即可平仓
                if margin <= hg.avgMargin {
                    hg.state = StateClosePending
                    hg.pendingTime = time.Now().Unix()

                    hg.shortPrice = youPrice
                    hg.longPrice = zuoPrice

                    go hg.short.Trade(TypeCloseShort, hg.tradeAmount, hg.shortPrice + hg.tradePriceOffset)
                    go hg.long.Trade(TypeCloseLong, hg.tradeAmount, hg.longPrice - hg.tradePriceOffset)
                }

            //如果是左手做空的
            } else {

                //差价高于平均差价即可平仓
                if margin >= hg.avgMargin {
                    hg.state = StateClosePending
                    hg.pendingTime = time.Now().Unix()

                    hg.shortPrice = zuoPrice
                    hg.longPrice = youPrice

                    go hg.short.Trade(TypeCloseShort, hg.tradeAmount, hg.shortPrice + hg.tradePriceOffset)
                    go hg.long.Trade(TypeCloseLong, hg.tradeAmount, hg.longPrice - hg.tradePriceOffset)
                }
            }
        }
    }
}

func (hg *HedgerWS) pending() {
    for !hg.stoped {
        order := <-hg.lastOrder
        t := time.Now().Unix()
        if hg.state == StateOpenPending {
            if order.Type == TypeOpenShort {
                hg.shortOrder = order
            } else if order.Type == TypeOpenLong {
                hg.longOrder = order
            }

            if t - hg.pendingTime < 2 {
                if hg.shortOrder.DealAmount == hg.longOrder.DealAmount {
                     
                }
            }

        } else if hg.state == StateClosePending {
            if order.Type == TypeCloseShort {
                hg.shortOrder = order
            } else if order.Type == TypeCloseLong {
                hg.longOrder = order
            }
        }
    }
}

func (hg *HedgerWS) checkPending() {
    for _ = range time.Tick(5 * time.Second) {
        if !hg.stoped {
            return
        }
        if hg.state == StateOpenPending {
        } else if hg.state == StateClosePending {
        }
    }
}

/*
func (hg *HedgerWS) openPosition(short *OKFutureWS, shortSellPrice float64, long *OKFutureWS, longBuyPrice float64) {



    //交易统计
    time.Sleep(2 * time.Second)
    sorder := short.OrderInfo(sid)
    if sorder == nil {
        return
    }
    if sorder.DealAmount <= 0 {
        sorder.DealAmount = hg.tradeAmount
    }
    if sorder.AvgPrice <= 0 {
        sorder.AvgPrice = shortSellPrice
    }

    lorder := long.OrderInfo(lid)
    if lorder == nil {
        return
    }
    if lorder.DealAmount <= 0 {
        lorder.DealAmount = hg.tradeAmount
    }
    if lorder.AvgPrice <= 0 {
        lorder.AvgPrice = longBuyPrice
    }

    hg.btc += lorder.DealAmount - hg.tradeAmount
    hg.cny += sorder.AvgPrice - lorder.AvgPrice
    hg.tcny += shortSellPrice - longBuyPrice

    gmvc.Logger.Println("open position:")
    gmvc.Logger.Println(fmt.Sprintf("   short: %v - %.2f(%.2f) btc, + %.2f(%.2f) cny",
        short.name, hg.tradeAmount, sorder.DealAmount, shortSellPrice, sorder.AvgPrice))
    gmvc.Logger.Println(fmt.Sprintf("   long: %v + %.2f(%.2f) btc, - %.2f(%.2f) cny",
        long.name, hg.tradeAmount, lorder.DealAmount, longBuyPrice, lorder.AvgPrice))
    gmvc.Logger.Println("")
}

func (hg *HedgerWS) openShort(short *OKFutureWS, sellPrice float64) int64 {
    id := short.Sell(hg.tradeAmount)
    hg.short = short
    return id
}

func (hg *HedgerWS) openLong(long *OKFutureWS, buyPrice float64) int64 {
    delta := 0.0;
    if long.name == "okcoin" {
        delta = 0.006
    }
    id := long.Buy((hg.tradeAmount + delta) * buyPrice)
    hg.long = long
    return id
}


func (hg *HedgerWS) closePosition(buyPrice, sellPrice float64) {
    var sid, lid int64
    if hg.short.name == "huobi" {
        sid = hg.closeShort(buyPrice)
        if sid == 0 {
            return
        }
        lid = hg.closeLong(sellPrice)
    } else {
        lid = hg.closeLong(sellPrice)
        if lid == 0 {
            return
        }
        sid = hg.closeShort(buyPrice)
    }
    hg.state = STATE_CLOSE
    hg.tradeNum++

    //交易统计
    time.Sleep(2 * time.Second)
    sorder := hg.short.OrderInfo(sid)
    if sorder == nil {
        return
    }

    if sorder.DealAmount <= 0 {
        sorder.DealAmount = hg.tradeAmount
    }
    if sorder.AvgPrice <= 0 {
        sorder.AvgPrice = buyPrice
    }

    lorder := hg.long.OrderInfo(lid)
    if lorder == nil {
        return
    }

    if lorder.DealAmount <= 0 {
        lorder.DealAmount = hg.tradeAmount
    }
    if lorder.AvgPrice <= 0 {
        lorder.AvgPrice = sellPrice
    }

    hg.btc += sorder.DealAmount - lorder.DealAmount
    hg.cny += lorder.AvgPrice - sorder.AvgPrice
    hg.tcny += sellPrice - buyPrice

    gmvc.Logger.Println("close position:")
    gmvc.Logger.Println(fmt.Sprintf("   short: %v + %.2f(%.2f) btc, - %.2f(%.2f) cny",
        hg.short.name, hg.tradeAmount, sorder.DealAmount, buyPrice, sorder.AvgPrice))
    gmvc.Logger.Println(fmt.Sprintf("   long: %v - %.2f(%.2f) btc, + %.2f(%.2f) cny",
        hg.long.name, hg.tradeAmount, lorder.DealAmount, sellPrice, lorder.AvgPrice))
    gmvc.Logger.Println("")

    now := time.Now()
    d := time.Unix(now.Unix() - hg.started.Unix(), 0)
    gmvc.Logger.Println(fmt.Sprintf("profit: %.4f btc, %.2f(%.2f) cny, %v min, %v round %v",
        hg.btc, hg.tcny * hg.tradeAmount, hg.cny * hg.tradeAmount,
        d.Format("15:04:05"), hg.tradeNum, now.Format("15:04:05")))
    gmvc.Logger.Println("")
}

func (hg *HedgerWS) closeShort(price float64) int64 {
    delta := 0.0;
    if hg.short.name == "okcoin" {
        delta = 0.006
    }
    return hg.short.Buy((hg.tradeAmount + delta) * price)
}

func (hg *HedgerWS) closeLong(price float64) int64 {
    return hg.long.Sell(hg.tradeAmount)
}

*/

