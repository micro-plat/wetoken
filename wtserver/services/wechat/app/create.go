package app

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/wechat/app"
)

//CreateHandler 创建微信公众号信息
type CreateHandler struct {
	app app.IWechatApp
}

//NewCreateHandler 创建服务
func NewCreateHandler() (u *CreateHandler) {
	return &CreateHandler{}
}

//Handle 创建app base info
//向数据库添加微信公众号信息
func (u *CreateHandler) Handle(ctx hydra.IContext) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	var input app.AppBaseInfo
	if err := ctx.Request().Bind(&input); err != nil {
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
