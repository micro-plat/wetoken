package token

import (
	"fmt"
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

type GetTokenHandler struct{}

//NewGetTokenHandler 创建服务
func NewGetTokenHandler() (u *GetTokenHandler) {
	return &GetTokenHandler{}
}
func (u *GetTokenHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		token.AccessToken
	}

	appid := ctx.Request().Path().Params().GetString("appid")
	if appid == "" {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("参数appid错误,%s", appid))
	}
	tokenObj := token.NewToken(appid)
	tokenInfo, err := tokenObj.Query()
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	result.AccessToken = *tokenInfo
	result.AccessToken.Reset()
	return result
}

//Handle 获取access token
//1. 从缓存中获取，不存在或过期时从数据库中获取
//2. 从数据库中获取，不存在或过期时从微信官网获取
//3. 从微信官网获取成功后，更新本地缓存和数据库
func (u *GetTokenHandler) Handle(ctx hydra.IContext) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		token.AccessToken
	}

	appid := ctx.Request().Path().Params().GetString("appid")
	if appid == "" {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("参数appid错误,%s", appid))
	}
	tokenObj := token.NewToken(appid)
	tokenInfo, err := tokenObj.Get()
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	result.AccessToken = *tokenInfo
	result.AccessToken.Reset()
	return result
}
