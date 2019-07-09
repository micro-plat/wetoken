package token

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

type GetHandler struct {
	container component.IContainer
	appid     string
	token     token.IToken
}

func NewGetHandlerBy(appid string) func(container component.IContainer) (u *GetHandler) {
	return func(container component.IContainer) (u *GetHandler) {
		return &GetHandler{
			container: container,
			appid:     appid,
			token:     token.NewToken(container, appid),
		}
	}
}

//NewGetHandler 创建服务
func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{
		container: container,
	}
}
func (u *GetHandler) Query(ctx *context.Context) (r interface{}) {
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

//Handle 获取access token
//1. 从缓存中获取，不存在或过期时从数据库中获取
//2. 从数据库中获取，不存在或过期时从微信官网获取
//3. 从微信官网获取成功后，更新本地缓存和数据库
func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		token.AccessToken
	}
	token, err := u.token.Get()
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
