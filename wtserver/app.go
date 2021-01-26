package main

import (
	"fmt"

	"github.com/micro-plat/hydra"
	xapp "github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/wetoken/modules/const/sql"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/app"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/ticket"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/token"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/wxcode"
)

var App = hydra.NewApp(
	hydra.WithPlatName("wetoken", "微信token维护服务"),
	hydra.WithSystemName("wtserver"),
	hydra.WithServerTypes(http.API, cron.CRON),
	hydra.WithClusterName("prod"))

func init() {

	install()

	sql.Install()

	App.OnStarting(func(appConf xapp.IAPPConf) error {
		_, err := hydra.C.DB().GetDB()
		if err != nil {
			return fmt.Errorf("db数据库配置错误,err:%v", err)
		}
		return nil
	})

	//根据配置注册服务
	App.API("/:appid/wechat/token/get", token.NewGetTokenHandler)     //接口，获取access token
	App.API("/:appid/wechat/token/query", token.NewQueryTokenHandler) //接口，获取access token
	App.API("/:appid/wechat/ticket/get", ticket.NewGetTicketHandler)  //接口，获取jsapi ticket
	App.API("/:appid/wechat/wxcode/send", wxcode.NewSendMessgHandler) //接口，发送微信验证码
	App.API("/wechat/app/create", app.NewCreateHandler)               //接口,用于添加微信公众号基础参数

	App.CRON("/wechat/token/refresh", token.NewRefreshTokenHandler, "@every 1m")    //自动流程，定时刷新access token
	App.CRON("/wechat/ticket/refresh", ticket.NewRefreshTicketHandler, "@every 1m") //自动流程，定时刷新jsapi ticket
}
