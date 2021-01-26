package token

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

type QueryHandler struct {
	appid string
	token token.IToken
}

func NewQueryHandlerBy(appid string) func() (u *QueryHandler) {
	return func() (u *QueryHandler) {
		return &QueryHandler{
			appid: appid,
			token: token.NewToken(appid),
		}
	}
}

//NewQueryHandler 创建服务
func NewQueryHandler() (u *QueryHandler) {
	return &QueryHandler{}
}
func (u *QueryHandler) Handle(ctx hydra.IContext) (r interface{}) {
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
