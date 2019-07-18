package app

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wetoken/modules/wechat/app"
)

//CreateHandler 创建微信公众号信息
type CreateHandler struct {
	container component.IContainer
	app       app.IWechatApp
}

func NewCreateHandlerBy() func(container component.IContainer) (u *CreateHandler) {
	return func(container component.IContainer) (u *CreateHandler) {
		return &CreateHandler{
			container: container,
			app:       app.NewWechatApp(container),
		}
	}
}

//NewCreateHandler 创建服务
func NewCreateHandler(container component.IContainer) (u *CreateHandler) {
	return &CreateHandler{
		container: container,
	}
}

//Handle 创建app base info
//向数据库添加微信公众号信息
func (u *CreateHandler) Handle(ctx *context.Context) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var input app.AppBaseInfo
	if err := ctx.Request.Bind(&input); err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	if err := u.app.Save(input); err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	return result
}
