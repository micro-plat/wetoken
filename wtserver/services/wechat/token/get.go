package token

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

type GetHandler struct {
	appid string
	token token.IToken
}

func NewGetHandlerBy(appid string) func() (u *GetHandler) {
	return func() (u *GetHandler) {
		return &GetHandler{
			appid: appid,
			token: token.NewToken(appid),
		}
	}
}

//NewGetHandler 创建服务
func NewGetHandler() (u *GetHandler) {
	return &GetHandler{}
}
func (u *GetHandler) Query(ctx hydra.IContext) (r interface{}) {
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
func (u *GetHandler) Handle(ctx hydra.IContext) (r interface{}) {
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
