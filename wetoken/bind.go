package wetoken

import (
	"fmt"

	"github.com/micro-plat/hydra"
	xapp "github.com/micro-plat/wetoken/wtserver/services/wechat/app"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/ticket"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/token"
)

var once1 sync.Once
var once2 sync.Once

//Bind 绑定公众号token,ticket对应的刷新，获取，消息推送接口等
func Bind(bindAddApp bool, appids ...string) error {
	once1.Do(func() {
		for _, appid := range appids {
			hydra.S.Micro(fmt.Sprintf("/%s/wechat/token/get", appid), token.NewGetHandlerBy(appid))   //接口，获取access token
			hydra.S.Micro(fmt.Sprintf("/%s/wechat/ticket/get", appid), ticket.NewGetHandlerBy(appid)) //接口，获取jsapi ticket

			hydra.S.CRON(fmt.Sprintf("/%s/wechat/token/refresh", appid), token.NewRefreshHandlerBy(appid))   //自动流程，定时刷新access token
			hydra.S.CRON(fmt.Sprintf("/%s/wechat/ticket/refresh", appid), ticket.NewRefreshHandlerBy(appid)) //自动流程，定时刷新jsapi ticket
			hydra.CRON.Add("@every 1m", fmt.Sprintf("/%s/wechat/token/refresh", appid))
			hydra.CRON.Add("@every 1m", fmt.Sprintf("/%s/wechat/ticket/refresh", appid))
		}
		if bindAddApp {
			hydra.S.Micro("/wechat/app/create", xapp.NewCreateHandlerBy()) //接口,用于添加微信公众号基础参数
		}
	})
	return nil
}

//Cron 只绑定定时任务，每隔1分钟检查一次token,ticket
func Cron(bindAddApp bool, appids ...string) error {
	once2.Do(func() {
		for _, appid := range appids {
			hydra.S.CRON(fmt.Sprintf("/%s/wechat/token/refresh", appid), token.NewRefreshHandlerBy(appid))   //自动流程，定时刷新access token
			hydra.S.CRON(fmt.Sprintf("/%s/wechat/ticket/refresh", appid), ticket.NewRefreshHandlerBy(appid)) //自动流程，定时刷新jsapi ticket
			hydra.CRON.Add("@every 1m", fmt.Sprintf("/%s/wechat/token/refresh", appid))
			hydra.CRON.Add("@every 1m", fmt.Sprintf("/%s/wechat/ticket/refresh", appid))
		}
		if bindAddApp {
			hydra.S.Micro("/wechat/app/create", xapp.NewCreateHandlerBy()) //接口,用于添加微信公众号基础参数
		}
	})
	return nil
}
