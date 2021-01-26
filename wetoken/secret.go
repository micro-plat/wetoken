package wetoken

import (
	"github.com/micro-plat/wetoken/modules/wechat/app"
)

//AppInfoResult .
type AppInfoResult struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	app.WechatAppInfo
}

// GetWxSecret 获取平台信息
func GetWxSecret(appID string) (results *AppInfoResult, err error) {
	appLib := app.NewWechatApp()
	appInfo, err := appLib.Get(appID)
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
