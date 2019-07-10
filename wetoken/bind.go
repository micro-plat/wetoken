package wetoken

import (
	"fmt"

	"github.com/micro-plat/hydra/conf"

	"github.com/micro-plat/hydra/hydra"
	xapp "github.com/micro-plat/wetoken/services/wechat/app"
	"github.com/micro-plat/wetoken/services/wechat/ticket"
	"github.com/micro-plat/wetoken/services/wechat/token"
)

//Bind 绑定公众号token,ticket对应的刷新，获取，消息推送接口等
func Bind(app *hydra.MicroApp, bindAddApp bool, appids ...string) error {
	crons := app.GetDynamicCron()
	for _, appid := range appids {
		app.Micro(fmt.Sprintf("/%s/wechat/token/get", appid), token.NewGetHandlerBy(appid))   //接口，获取access token
		app.Micro(fmt.Sprintf("/%s/wechat/ticket/get", appid), ticket.NewGetHandlerBy(appid)) //接口，获取jsapi ticket

		app.Flow(fmt.Sprintf("/%s/wechat/token/refresh", appid), token.NewRefreshHandlerBy(appid))   //自动流程，定时刷新access token
		app.Flow(fmt.Sprintf("/%s/wechat/ticket/refresh", appid), ticket.NewRefreshHandlerBy(appid)) //自动流程，定时刷新jsapi ticket

		crons <- &conf.Task{Cron: "@every 1m", Service: fmt.Sprintf("/%s/wechat/token/refresh", appid)}
		crons <- &conf.Task{Cron: "@every 1m", Service: fmt.Sprintf("/%s/wechat/ticket/refresh", appid)}

	}
	if bindAddApp {
		app.Micro("/wechat/app/create", xapp.NewCreateHandlerBy()) //接口,用于添加微信公众号基础参数
	}
	return nil
}

//Cron 只绑定定时任务，每隔1分钟检查一次token,ticket
func Cron(app *hydra.MicroApp, bindAddApp bool, appids ...string) error {
	crons := app.GetDynamicCron()
	for _, appid := range appids {
		app.Flow(fmt.Sprintf("/%s/wechat/token/refresh", appid), token.NewRefreshHandlerBy(appid))   //自动流程，定时刷新access token
		app.Flow(fmt.Sprintf("/%s/wechat/ticket/refresh", appid), ticket.NewRefreshHandlerBy(appid)) //自动流程，定时刷新jsapi ticket
		crons <- &conf.Task{Cron: "@every 1m", Service: fmt.Sprintf("/%s/wechat/token/refresh", appid)}
		crons <- &conf.Task{Cron: "@every 1m", Service: fmt.Sprintf("/%s/wechat/ticket/refresh", appid)}

	}
	if bindAddApp {
		app.Micro("/wechat/app/create", xapp.NewCreateHandlerBy()) //接口,用于添加微信公众号基础参数
	}
	return nil
}
