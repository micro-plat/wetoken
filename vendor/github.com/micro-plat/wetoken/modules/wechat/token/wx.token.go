package token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//WechatToken 微信官网access token操作对象
type WechatToken struct {
	httpClient *http.Client
	appid      string
}

//NewWechatToken 创建通过微信官网获取access token操作对象
func NewWechatToken(appid string) *WechatToken {
	client := *http.DefaultClient
	client.Timeout = time.Second * 15
	return &WechatToken{
		httpClient: &client,
		appid:      appid,
	}
}

//Get 从微信服务器获取新的 access_token 并存入缓存, 同时返回该 access_token.
func (t WechatToken) Get(secret string) (token *AccessToken, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + t.appid +
		"&secret=" + secret
	httpResp, err := t.httpClient.Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		AccessToken
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}
	if result.ErrCode != ErrCodeOK {
		err = fmt.Errorf("err:%d %s", result.ErrCode, result.ErrMsg)
		return
	}

	// // 由于网络的延时, access_token 过期时间留有一个缓冲区
	// switch {
	// case result.ExpiresIn > 31556952: // 60*60*24*365.2425
	// 	err = errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
	// 	return
	// case result.ExpiresIn > 60*60:
	// 	result.ExpiresIn -= 60 * 10
	// case result.ExpiresIn > 60*30:
	// 	result.ExpiresIn -= 60 * 5
	// case result.ExpiresIn > 60*5:
	// 	result.ExpiresIn -= 60
	// case result.ExpiresIn > 60:
	// 	result.ExpiresIn -= 10
	// default:
	// 	err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
	// 	return
	// }
	result.AccessToken.ExpiresDate = jtime(time.Now().Add(time.Second * time.Duration(result.AccessToken.ExpiresIn)))
	tokenCopy := result.AccessToken
	token = &tokenCopy
	return
}
