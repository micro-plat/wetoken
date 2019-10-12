package wetoken

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/wetoken/modules/wechat/app"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

type TokenResult struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	token.AccessToken
}

type AppInfoResult struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	app.WechatAppInfo
}

//GetWxToken 获取token
func GetWxToken(container component.IContainer, appID string) (results *TokenResult, err error) {

	tokens := token.NewToken(container, appID)

	token, err := tokens.Get()
	result := &TokenResult{
		ErrCode: 0,
		ErrMsg:  "success",
	}
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result, nil
	}

	result.AccessToken = *token
	result.AccessToken.Reset()
	return result, nil
}

// GetWxSecret 获取平台信息
func GetWxSecret(container component.IContainer, appID string) (results *AppInfoResult, err error) {

	appHandle := app.NewWechatApp(container)

	appInfo, err := appHandle.Get(appID)
	result := &AppInfoResult{
		ErrCode: 0,
		ErrMsg:  "success",
	}
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result, nil
	}

	result.WechatAppInfo = *appInfo
	return result, nil
}
