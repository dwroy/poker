package controller

import (
    "github.com/roydong/gmvc"
    "github.com/roydong/poker/service/market"
)

func init() {
    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {

        name, _ := r.String("$1")



        sess := r.Session

        sess.Set("aaa", name)



        return r.HtmlResponse(name, "index/index")

    }, `/home/(\w+)`)


    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        huobi := market.NewMarket("huobi")
        okcoin := market.NewMarket("okcoin")
        haobtc := market.NewMarket("haobtc")

        oka, okb := okcoin.GetDepth()
        hba, hbb := huobi.GetDepth()

        hta, htb := haobtc.GetDepth()

        a := map[string][][]float64{
            "ok_ask": oka,
            "ok_bid": okb,

            "hb_ask": hba,
            "hb_bid": hbb,

            "ht_ask": hta,
            "ht_bid": htb,
        }

        return r.JsonResponse(a)
    }, "/depth")

    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        haobtc := market.NewMarket("haobtc")

        ticker := haobtc.LastTicker()

        return r.JsonResponse(ticker)
    }, "/ticker")

    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        okcoin := market.NewMarket("okcoin")

        b, c := okcoin.GetBalance()

        huobi := market.NewMarket("huobi")

        b1, c1 := huobi.GetBalance()

        haobtc := market.NewMarket("haobtc")

        b2, c2 := haobtc.GetBalance()

        return r.JsonResponse([]float64{b, c, b1, c1, b2, c2, c + c1 + c2})
    }, "/balance")

    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        huobi := market.NewMarket("huobi")

        //huobi.Sell(0.02)
        huobi.Buy(50)

        return r.TextResponse("done")

    }, "/trade_hb")

    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        haobtc := market.NewMarket("haobtc")

        haobtc.Sell(0.02)
        haobtc.Buy(50)

        return r.TextResponse("done")

    }, "/trade_ht")

    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {

        tree := gmvc.NewTree()
        tree.LoadJson("", []byte(`{"code":1,"msg":"服务器繁忙","message":"服务器繁忙"}`), false)

        gmvc.Logger.Println(tree.Float64("code"))

        return r.TextResponse("aa")

    }, "/trade_ok")


    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        okcoin := market.NewOKContact()

        v := okcoin.Ticker()

        i := okcoin.Index()

        p := (v.Last - i) / i

        return r.JsonResponse([]float64{v.Last, i, gmvc.Round(p, 4) * 100})

    }, "/okcoin_premium")

    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        okcoin := market.NewMarket("okcoin")
        huobi := market.NewMarket("huobi")

        hg := market.NewHedger(okcoin, huobi)
        hg.Start()

        return r.TextResponse("done")

    }, "/okcoin_huobi")

    gmvc.SetAction(func(r *gmvc.Request) *gmvc.Response {
        okcoin := market.NewMarket("okcoin")
        haobtc := market.NewMarket("haobtc")

        hg := market.NewHedger(okcoin, haobtc)
        hg.Start()

        return r.TextResponse("done")

    }, "/okcoin_haobtc")


    gmvc.WSActionMap["ws"] = func(wsm *gmvc.WSMessage) {
        val, _ := wsm.String("a")

        q := map[string]interface{}{
            "aaa": "sss",
            "value": val,
        }
        wsm.Send("name", q, nil)
    }
}

