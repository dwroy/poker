package arbitrage

import (
    "github.com/roydong/gmvc"
    "fmt"
    "time"
    "github.com/gorilla/websocket"
    "strings"
    "sync"
)


type OKFutureWS struct {
    *gmvc.Event

    wsHost   string
    apiKey    string
    apiSecret string
    contractType string
    leverRate int

    conn *websocket.Conn

    lastTicker Ticker
    lastTrade Trade
    lastAsks [][]float64
    lastBids [][]float64

    lastBtc chan float64
    lastBtcLocker *sync.Mutex

    tradeLocker *sync.Mutex
    lastOrderId chan int64
    cancelOrderId chan int64

    tickerUpdated int64
    depthUpdated int64
    balanceUpdated int64
    msgUpdated int64

    currentOrders map[int64]Order

    subChannels []string

    dealAmount []float64
    totalPrice []float64
}

func NewOKFutureWS(contractType string) *OKFutureWS {
    conf := gmvc.Store.Tree("config.market.okfuture")
    ok := &OKFutureWS{}
    ok.Event = gmvc.NewEvent()
    ok.wsHost, _ = conf.String("ws_host")
    ok.apiKey, _ = conf.String("api_key")
    ok.apiSecret, _ = conf.String("api_secret")
    ok.contractType = contractType
    ok.leverRate = 10

    ok.lastTicker = Ticker{}
    ok.lastTrade = Trade{}

    ok.lastBtc = make(chan float64)
    ok.lastBtcLocker = &sync.Mutex{}

    ok.tradeLocker = &sync.Mutex{}
    ok.lastOrderId = make(chan int64)
    ok.cancelOrderId = make(chan int64)

    ok.currentOrders = make(map[int64]Order, 4)

    ok.subChannels = []string{
        //ticker订阅
        "ok_sub_futureusd_btc_ticker_" + ok.contractType,
        //最新深度订阅
        fmt.Sprintf("ok_sub_futureusd_btc_depth_%s_%d", ok.contractType, 20),
        //最新交易单订阅
        "ok_sub_futureusd_btc_trade_" + ok.contractType,
        //订单交易结果订阅
        "ok_sub_futureusd_trades",
    }

    ok.initCallbacks()
    ok.Connect()

    return ok
}

func (ok *OKFutureWS) Connect() {
    var err error
    ok.conn, _, err = (&websocket.Dialer{}).Dial(ok.wsHost, nil)
    if err != nil {
        gmvc.Logger.Fatalln(err)
    }

    go ok.readLoop()
    //ok.addChannel(ok.subChannels[0], nil)
    //ok.addChannel(ok.subChannels[1], nil)
    ok.addChannel(ok.subChannels[2], nil)
    ok.addChannel(ok.subChannels[3], map[string]interface{}{})
}

func (ok *OKFutureWS) Stop() {
    ok.removeChannels()
    ok.conn.Close()
}

func (ok *OKFutureWS) initCallbacks() {
    ok.AddSyncHandler(ok.subChannels[0], ok.syncTicker)
    ok.AddSyncHandler(ok.subChannels[1], ok.syncDepth)
    ok.AddSyncHandler(ok.subChannels[2], ok.syncTrade)
    ok.AddSyncHandler(ok.subChannels[3], ok.syncCurrentOrder)
    ok.AddSyncHandler("ok_futureusd_userinfo", ok.syncBalance)
    ok.AddSyncHandler("ok_futuresusd_trade", ok.syncTradeResult)
    ok.AddSyncHandler("ok_futureusd_cancel_order", ok.syncCancelResult)
}

func (ok *OKFutureWS) Name() string {
    return ok.contractType
}

func (ok *OKFutureWS) syncTicker(args ...interface{}) {
    rs, _ := args[0].(*gmvc.Tree)
    if rs == nil {
        return
    }

    t := Ticker{}
    t.High, _ = rs.Float("high")
    t.Low,  _ = rs.Float("low")
    t.Ask, _ = rs.Float("sell")
    t.Bid,  _ = rs.Float("buy")
    t.Last, _ = rs.Float("last")
    t.Vol,  _ = rs.Float("volume")
    t.Time    = time.Now().Unix()
    ok.tickerUpdated = t.Time
    ok.lastTicker = t
}

func (ok *OKFutureWS) LastTicker() Ticker {
    return ok.lastTicker
}

func (ok *OKFutureWS) syncTrade(args ...interface{}) {
    rs, _ := args[0].(*gmvc.Tree)
    if rs == nil {
        return
    }
    var trade Trade
    for i, l := 0, rs.NodeNum(""); i < l; i++ {
        trade = Trade{}
        trade.Id, _ = rs.Int64(fmt.Sprintf("%d.0", i))
        trade.Price, _ = rs.Float(fmt.Sprintf("%d.1", i))
        trade.Amount, _ = rs.Float(fmt.Sprintf("%d.2", i))
        trade.Time, _ = rs.Int64(fmt.Sprintf("%d.3", i))
        trade.Type, _ = rs.String(fmt.Sprintf("%d.4", i))
        ok.lastTrade = trade
        ok.Trigger("new_trade", trade)
    }

    ok.Trigger("last_trade", trade)
}

func (ok *OKFutureWS) LastTrade() Trade {
    return ok.lastTrade
}

func (ok *OKFutureWS) syncDepth(args ...interface{}) {
    rs, _ := args[0].(*gmvc.Tree)
    if rs == nil {
        gmvc.Logger.Println("depth data is nil")
        return
    }

    l := rs.NodeNum("asks")
    asks := make([][]float64, 0, l)
    for i := l - 1; i >= 0; i-- {
        price, _ := rs.Float(fmt.Sprintf("asks.%v.0", i))
        amount, _ := rs.Float(fmt.Sprintf("asks.%v.1", i))
        asks = append(asks, []float64{price, amount})
    }

    l = rs.NodeNum("bids")
    bids := make([][]float64, 0, l)
    for i := 0; i < l; i++ {
        price, _ := rs.Float(fmt.Sprintf("bids.%v.0", i))
        amount, _ := rs.Float(fmt.Sprintf("bids.%v.1", i))
        bids = append(bids, []float64{price, amount})
    }

    ok.lastAsks = asks
    ok.lastBids = bids
    ok.depthUpdated = time.Now().Unix()
    ok.Trigger("depth", asks, bids, ok.depthUpdated)
}

func (ok *OKFutureWS) GetDepth() ([][]float64, [][]float64) {
    return ok.lastAsks, ok.lastBids
}

func (ok *OKFutureWS) DepthUpdated() int64 {
    return ok.depthUpdated
}

func (ok *OKFutureWS) syncBalance(args ...interface{}) {
    code, _ := args[1].(int64)
    var btc float64
    if code > 0 {
        btc = 0
    } else {
        rs, _ := args[0].(*gmvc.Tree)
        btc, _ = rs.Float("info.btc.balance")
    }
    ok.balanceUpdated = time.Now().Unix()
    ok.lastBtc <-btc
}

func (ok *OKFutureWS) GetBalance() (float64, float64) {
    ok.lastBtcLocker.Lock()
    defer ok.lastBtcLocker.Unlock()
    ok.addChannel("ok_futureusd_userinfo", make(map[string]interface{}))
    r, _ :=  <-ok.lastBtc, 0
    return r, 0
}

func (ok *OKFutureWS) Trade(typ int, amount, price float64) int64 {
    ok.tradeLocker.Lock()
    defer ok.tradeLocker.Unlock()
    return ok.tradeNolock(typ, amount, price)
}

func (ok *OKFutureWS) tradeNolock(typ int, amount, price float64) int64 {
    params := map[string]interface{}{
        "symbol": "btc_usd",
        "contract_type": ok.contractType,
        "type": typ,
        "amount": amount,
        "price": price,
        "match_price": 0,
        "lever_rate": ok.leverRate,
    }
    if price == 0 {
        params["mathc_price"] = 1
    }
    ok.addChannel("ok_futuresusd_trade", params)
    return <-ok.lastOrderId
}


func (ok *OKFutureWS) syncTradeResult(args ...interface{}) {
    var id int64
    rs, _ := args[0].(*gmvc.Tree)
    if rs != nil {
        id, _ = rs.Int64("order_id")
    }
    ok.lastOrderId <-id
}

func (ok *OKFutureWS) Order(id int64) Order {
    return ok.currentOrders[id]
}

func (ok *OKFutureWS) CancelOrder(id int64) int64 {
    ok.tradeLocker.Lock()
    defer ok.tradeLocker.Unlock()
    return ok.cancelOrderNolock(id)
}

func (ok *OKFutureWS) cancelOrderNolock(id int64) int64 {
    params := map[string]interface{}{
        "symbol": "btc_usd",
        "contract_type": ok.contractType,
        "order_id": id,
    }
    ok.addChannel("ok_futureusd_cancel_order", params)
    return <-ok.cancelOrderId
}

func (ok *OKFutureWS) syncCancelResult(args ...interface{}) {
    var id int64
    rs, _ := args[0].(*gmvc.Tree)
    if rs != nil {
        id, _ = rs.Int64("order_id")
    }
    ok.cancelOrderId <-id
}

func (ok *OKFutureWS) syncCurrentOrder(args ...interface{}) {
    rs, _ := args[0].(*gmvc.Tree)
    if rs == nil {
        return
    }

    order := Order{}
    order.Id, _ = rs.Int64("orderid")
    order.Type, _ = rs.Int("type")
    order.Amount, _ = rs.Float("amount")
    order.Price, _ = rs.Float("price")
    order.DealAmount, _ = rs.Float("deal_amount")
    order.AvgPrice, _ = rs.Float("price_avg")
    order.Status, _ = rs.Int("status")
    order.Fee, _ = rs.Float("fee")

    if order.Status == OrderStatusCreated ||
            order.Status == OrderStatusCancel ||
            order.Status == OrderStatusCanceling {
        return
    }

    ok.currentOrders[order.Id] = order
    ok.dealAmount = make([]float64, 5)
    ok.totalPrice = make([]float64, 5)
    for _, order := range ok.currentOrders {
        ok.dealAmount[order.Type] += order.DealAmount
        ok.totalPrice[order.Type] += order.AvgPrice * order.DealAmount
    }
    ok.Trigger("order_change", order)
}

func (ok *OKFutureWS) ClearOrders() {
    ok.currentOrders = make(map[int64]Order, 10)
    ok.dealAmount = make([]float64, 5)
    ok.totalPrice = make([]float64, 5)
}

func (ok *OKFutureWS) GetAvgPrice(typ int) float64 {
    amount := ok.dealAmount[typ]
    if amount > 0 {
        return ok.totalPrice[typ] / amount
    }
    return 0
}

func (ok *OKFutureWS) signParams(params map[string]interface{}) map[string]interface{} {
    if params == nil {
        params = make(map[string]interface{}, 2)
    }
    params["api_key"] = ok.apiKey
    params["sign"] = strings.ToUpper(createSignature(params, ok.apiSecret))
    return params
}

func (ok *OKFutureWS) addChannel(name string, params map[string]interface{}) {
    msg := map[string]interface{} {
        "event": "addChannel",
        "channel": name,
    }
    if params != nil {
        msg["parameters"] = ok.signParams(params)
    }
    err := ok.conn.WriteJSON(msg)
    if err != nil {
        gmvc.Logger.Fatalln("okfuture add channel failed")
    }
}

func (ok *OKFutureWS) RemoveChannel(name string) {
    err := ok.conn.WriteMessage(websocket.TextMessage,
        []byte(fmt.Sprintf(`{"event":"removeChannel","channel":"%s"}`, name)))
    if err != nil {
        gmvc.Logger.Fatalln("okfuture remove channel failed " + name)
    }
}

func (ok *OKFutureWS) removeChannels() {
    query := make([]string, 0, len(ok.subChannels))
    for _, channel := range ok.subChannels{
        query = append(query, fmt.Sprintf(`{"event":"removeChannel","channel":"%s"}`, channel))
    }
    err := ok.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[%s]", strings.Join(query, ","))))
    if err != nil {
        gmvc.Logger.Fatalln("okfuture remove channels failed")
    }
}

func (ok *OKFutureWS) Ping() {
    ok.conn.WriteMessage(websocket.TextMessage, []byte(`{"event":"ping"}`))
}

func (ok *OKFutureWS) readLoop() {
   for {
       typ, raw, err := ok.conn.ReadMessage()
       if err != nil {
           gmvc.Logger.Println("okfuture ws read message error ", err.Error())
           return
       }
       if typ != websocket.TextMessage {
           gmvc.Logger.Println("okfuture ws not text message", err.Error())
           continue
       }
       rs := gmvc.NewTree()
       err = rs.LoadJson("", raw, true)
       if err != nil {
           continue
       }
       ok.msgUpdated = time.Now().Unix()
       for i, l := 0, rs.NodeNum(""); i < l; i++ {
           event := rs.Tree(fmt.Sprintf("%d", i))
           if event == nil {
               gmvc.Logger.Println("okfuture ws error " + string(raw))
               continue
           }
           if name, has := event.String("channel"); has && len(name) > 0 {
               data := event.Tree("data")
               code, _ := event.Int64("errorcode")
               if code > 0 {
                   gmvc.Logger.Println("okfuture ws error " + string(raw))
               }
               ok.Trigger(name, data, code)
           }
       }
   }
}


