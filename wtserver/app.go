package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra"
	xapp "github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/app"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/ticket"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/token"
	"github.com/micro-plat/wetoken/wtserver/services/wechat/wxcode"
)

//AppWXConf 应用程序配置
type AppWXConf struct {
	WX []WXConfig `json:"wx" valid:"required"`
}

//WXConfig 微信公众号配置
type WXConfig struct {
	AppID string `json:"appid" valid:"ascii,required"`
}

var App = hydra.NewApp(
	hydra.WithPlatName("wetoken", "微信token维护服务"),
	hydra.WithSystemName("wtserver"),
	hydra.WithServerTypes(http.API, cron.CRON),
	hydra.WithClusterName("prod"))

func init() {

	install()

	App.OnStarting(func(appConf xapp.IAPPConf) error {

		_, err := hydra.C.DB().GetDB()
		if err != nil {
			return fmt.Errorf("db数据库配置错误,err:%v", err)
		}

		var wxConf AppWXConf
		_, err = appConf.GetServerConf().GetSubObject("app", &wxConf)
		if err != nil {
			return fmt.Errorf("pgs_flow_config配置错误,err:%v", err)
		}

		if b, err := govalidator.ValidateStruct(&wxConf); !b || len(wxConf.WX) == 0 {
			return fmt.Errorf("app 配置文件有误:%v", err)
		}

		//根据配置注册服务
		for _, wx := range wxConf.WX {
			App.Micro(fmt.Sprintf("/%s/wechat/token/get", wx.AppID), token.NewGetHandlerBy(wx.AppID)) //接口，获取access token

			App.Micro(fmt.Sprintf("/%s/wechat/token/query", wx.AppID), token.NewQueryHandlerBy(wx.AppID)) //接口，获取access token

			App.Micro(fmt.Sprintf("/%s/wechat/ticket/get", wx.AppID), ticket.NewGetHandlerBy(wx.AppID))   //接口，获取jsapi ticket
			App.Micro(fmt.Sprintf("/%s/wechat/wxcode/send", wx.AppID), wxcode.NewSendHandlerBy(wx.AppID)) //接口，发送微信验证码

			App.Flow(fmt.Sprintf("/%s/wechat/token/refresh", wx.AppID), token.NewRefreshHandlerBy(wx.AppID))   //自动流程，定时刷新access token
			App.Flow(fmt.Sprintf("/%s/wechat/ticket/refresh", wx.AppID), ticket.NewRefreshHandlerBy(wx.AppID)) //自动流程，定时刷新jsapi ticket

		}
		App.Micro("/wechat/app/create", app.NewCreateHandlerBy()) //接口,用于添加微信公众号基础参数
		return nil
	})
}
