package server

import (
    "dw/poker/context"
    "dw/poker/utils"
    "dw/poker/module/riskctrl"
    "net/http"
    "strings"
    "dw/poker/module/account"
)

type IModule interface {
    Init(conf *context.Config) error
    Run(ctx *context.Context) error
}

var routes = map[string][]IModule{
    "/account": {
        &account.Account{},
    },
    "/riskctrl": {
        &riskctrl.RiskCtrl{},
    },
}

func initRouter(conf *context.Config) error {
    for path, modules := range routes {
        for i := 0; i < len(modules); i++ {
            err := modules[i].Init(conf)
            if err != nil {
                utils.FatalLog.Write("init modules[%s/%d] fail. err[%s]", path, i, err.Error())
                return err
            }
        }
    }
    return nil
}

func initContext(route string) *context.Context {
    return &context.Context{}
}

var defaultRoute = "/"

func runHttpThread(req *http.Request) ([]byte, int) {
    path := strings.ToLower(strings.TrimSpace(req.URL.Path))
    modules := routes[path]
    if len(modules) <= 0 && len(defaultRoute) > 0 {
        modules = routes[defaultRoute]
    }
    if len(modules) <= 0 {
        utils.WarningLog.Write("http not found[%s]", req.URL.Path)
        return []byte("404 not found"), http.StatusNotFound
    }
    ctx := initContext(path)
    ctx.Request = req
    for i := 0; i < len(modules); i++ {
        err := modules[i].Run(ctx)
        if err != nil {
            utils.FatalLog.Write("run modules[%s/%d] fail. err[%s]", path, i, err.Error())
            return []byte(err.Error()), http.StatusInternalServerError
        }
    }
    return ctx.RespBody, http.StatusOK
}

