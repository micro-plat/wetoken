package ticket

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//WechatTicket 微信官网jsapi ticket操作对象
type WechatTicket struct {
	httpClient *http.Client
	appid      string
}

//NewWechatTicket 创建通过微信官网获取jsapi ticket操作对象
func NewWechatTicket(appid string) *WechatTicket {
	client := *http.DefaultClient
	client.Timeout = time.Second * 5
	return &WechatTicket{
		httpClient: &client,
		appid:      appid,
	}
}

//Get 从微信服务器获取新的 jsapi ticket 并存入缓存, 同时返回该 jsapi ticket.
func (t WechatTicket) Get(accessToken string) (Ticket *JSTicket, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token=" + accessToken
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
		JSTicket
	}
	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}
	if result.ErrCode != ErrCodeOK {
		err = fmt.Errorf("err:%d %s", result.ErrCode, result.ErrMsg)
		return
	}

	// 由于网络的延时, jsapi ticket 过期时间留有一个缓冲区
	switch {
	case result.ExpiresIn > 31556952: // 60*60*24*365.2425
		err = errors.New("expires_in too large: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	case result.ExpiresIn > 60*60:
		result.ExpiresIn -= 60 * 10
	case result.ExpiresIn > 60*30:
		result.ExpiresIn -= 60 * 5
	case result.ExpiresIn > 60*5:
		result.ExpiresIn -= 60
	case result.ExpiresIn > 60:
		result.ExpiresIn -= 10
	default:
		err = errors.New("expires_in too small: " + strconv.FormatInt(result.ExpiresIn, 10))
		return
	}
	result.JSTicket.ExpiresDate = jtime(time.Now().Add(time.Second * time.Duration(result.JSTicket.ExpiresIn)))
	TicketCopy := result.JSTicket
	Ticket = &TicketCopy
	return
}
