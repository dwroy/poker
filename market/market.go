package market

import (
    "dw/poker/protobuf/exsync"
    "dw/poker/context"
)

const (
    OkexWeek = "OkexWeek"
    OkexQuarter = "OkexQuarter"
    BitmexXbtusd = "BitmexXbtusd"
)

var exchanges = make(map[string]*Exchange, 0)

func AddExchange(ex *Exchange) {
    exchanges[ex.Name()] = ex
}

func GetExchange(name string) *Exchange {
    return exchanges[name]
}

func Init(conf *context.Config) {
    AddExchange(NewExchange(OkexQuarter, conf.Market.ExsyncHost))
    AddExchange(NewExchange(OkexWeek, conf.Market.ExsyncHost))

    /*
        bmconf := conf.Market.Bitmex
        bm, _ := bitmex.NewExchange(bmconf.HttpHost, bmconf.ApiKey, bmconf.ApiSecret, bmconf.Wss)
        AddExchange(NewExchange(bm))
    */
}




func GetProfit(pos *exsync.Position, price float64) float64 {
    profit := pos.Amount * (price - pos.AvgPrice)
    if pos.PType == exsync.PositionType_Long {
        return profit
    }
    return -profit
}

func GetROP(pos *exsync.Position, price float64) float64 {
    if pos.Money > 0 {
        return GetProfit(pos, price) / pos.Deposit
    }
    return 0
}
