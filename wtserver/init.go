package main

import (
	"fmt"

	"github.com/micro-plat/hydra/hydra"
	"github.com/urfave/cli"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
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

//init 检查应用程序配置文件，并根据配置初始化服务
func (r *wtserver) init() {

	r.Cli.Append(hydra.ModeRegistry, cli.BoolFlag{
		Name:  "db,d",
		Usage: "创建数据库表结构",
	})

	r.Initializing(func(c component.IContainer) error {
		//获取微信配置
		var wxConf AppWXConf
		if err := c.GetAppConf(&wxConf); err != nil {
			return err
		}
		if b, err := govalidator.ValidateStruct(&wxConf); !b || len(wxConf.WX) == 0 {
			return fmt.Errorf("app 配置文件有误:%v", err)
		}

		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}

		//根据配置注册服务
		for _, wx := range wxConf.WX {
			r.Micro(fmt.Sprintf("/%s/wechat/token/get", wx.AppID), token.NewGetHandlerBy(wx.AppID)) //接口，获取access token

			r.Micro(fmt.Sprintf("/%s/wechat/token/query", wx.AppID), token.NewQueryHandlerBy(wx.AppID)) //接口，获取access token

			r.Micro(fmt.Sprintf("/%s/wechat/ticket/get", wx.AppID), ticket.NewGetHandlerBy(wx.AppID))   //接口，获取jsapi ticket
			r.Micro(fmt.Sprintf("/%s/wechat/wxcode/send", wx.AppID), wxcode.NewSendHandlerBy(wx.AppID)) //接口，发送微信验证码

			r.Flow(fmt.Sprintf("/%s/wechat/token/refresh", wx.AppID), token.NewRefreshHandlerBy(wx.AppID))   //自动流程，定时刷新access token
			r.Flow(fmt.Sprintf("/%s/wechat/ticket/refresh", wx.AppID), ticket.NewRefreshHandlerBy(wx.AppID)) //自动流程，定时刷新jsapi ticket

		}
		r.Micro("/wechat/app/create", app.NewCreateHandlerBy()) //接口,用于添加微信公众号基础参数
		return nil
	})
}
