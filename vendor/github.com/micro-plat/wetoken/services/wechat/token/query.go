package token

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

type QueryHandler struct {
	container component.IContainer
	appid     string
	token     token.IToken
}

func NewQueryHandlerBy(appid string) func(container component.IContainer) (u *QueryHandler) {
	return func(container component.IContainer) (u *QueryHandler) {
		return &QueryHandler{
			container: container,
			appid:     appid,
			token:     token.NewToken(container, appid),
		}
	}
}

//NewQueryHandler 创建服务
func NewQueryHandler(container component.IContainer) (u *QueryHandler) {
	return &QueryHandler{
		container: container,
	}
}
func (u *QueryHandler) Handle(ctx *context.Context) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		token.AccessToken
	}
	token, err := u.token.Query()
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	result.AccessToken = *token
	result.AccessToken.Reset()
	return result
}
