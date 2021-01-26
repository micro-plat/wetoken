package wetoken

import (
	"sync"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/const/sql"
	xapp "github.com/micro-plat/wetoken/wtserver/services/wechat/app"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/ticket"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/token"
)

func init() {
	sql.Install()
}

var once1 sync.Once

//Bind 绑定公众号token,ticket对应的刷新，获取，消息推送接口等
func Bind(bindAddApp bool) error {
	once1.Do(func() {
		hydra.S.Micro("/:appid/wechat/token/get", token.NewGetTokenHandler)    //接口，获取access token
		hydra.S.Micro("/:appid/wechat/ticket/get", ticket.NewGetTicketHandler) //接口，获取jsapi ticket

		hydra.S.CRON("/wechat/token/refresh", token.NewRefreshTokenHandler)    //自动流程，定时刷新access token
		hydra.S.CRON("/wechat/ticket/refresh", ticket.NewRefreshTicketHandler) //自动流程，定时刷新jsapi ticket
		hydra.CRON.Add("@every 1m", "/wechat/token/refresh")
		hydra.CRON.Add("@every 1m", "/wechat/ticket/refresh")
		if bindAddApp {
			hydra.S.Micro("/wechat/app/create", xapp.NewCreateHandler) //接口,用于添加微信公众号基础参数
		}
	})
	return nil
}

//Cron 只绑定定时任务，每隔1分钟检查一次token,ticket
func Cron(bindAddApp bool) error {
	once1.Do(func() {
		hydra.S.CRON("/wechat/token/refresh", token.NewRefreshTokenHandler)    //自动流程，定时刷新access token
		hydra.S.CRON("/wechat/ticket/refresh", ticket.NewRefreshTicketHandler) //自动流程，定时刷新jsapi ticket
		hydra.CRON.Add("@every 1m", "/wechat/token/refresh")
		hydra.CRON.Add("@every 1m", "/wechat/ticket/refresh")
		if bindAddApp {
			hydra.S.Micro("/wechat/app/create", xapp.NewCreateHandler) //接口,用于添加微信公众号基础参数
		}
	})
	return nil
}
