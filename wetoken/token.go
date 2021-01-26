package wetoken

import (
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

//TokenResult .
type TokenResult struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	token.AccessToken
}

//GetWxToken 获取token
func GetWxToken(appID string) (results *TokenResult, err error) {
	tokenLib := token.NewToken(appID)
	token, err := tokenLib.Get()
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
