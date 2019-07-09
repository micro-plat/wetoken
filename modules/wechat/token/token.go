package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/concurrent/cmap"
	"github.com/micro-plat/wechat/mp"
)

var errTokenExpired = errors.New("access token has expired")
var errTokenDontExist = errors.New("access token does not exist")

const (
	ErrCodeOK                 = 0
	ErrCodeInvalidCredential  = 40001 // access_token 过期错误码
	ErrCodeAccessTokenExpired = 42001 // access_token 过期错误码(maybe!!!)
)

type jtime time.Time

func (this jtime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("20060102150405"))
	return []byte(stamp), nil
}

type AccessToken struct {
	Token       string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresDate jtime  `json:"expires_date"`
}

func (a *AccessToken) Reset() {
	a.ExpiresIn = -1 * int64(time.Since(time.Time(a.ExpiresDate)).Seconds())
}

//IToken access token操作接口
type IToken interface {
	//刷新token
	Refresh(force bool) (token *AccessToken, r bool, err error)
	//获取token
	Query() (token *AccessToken, err error)
	//获取token
	Get() (token *AccessToken, err error)

	GetContext() (*mp.Context, error)
}

//Token access token统一访问对象
type Token struct {
	dbToken    *DBToken
	cacheToken *CacheToken
	wxToken    *WechatToken
	lock       sync.Mutex
	appid      string
}

//NewToken 构建access token统一访问对象
func NewToken(c component.IContainer, appid string) *Token {
	return &Token{
		dbToken:    NewDBToken(c, appid),
		cacheToken: NewCacheToken(appid),
		wxToken:    NewWechatToken(appid),
		appid:      appid,
	}
}

//Save 保存access token到数据库和缓存
func (t *Token) save(tk *AccessToken) error {
	//保存到数据库
	if err := t.dbToken.Save(tk); err != nil {
		return err
	}
	//保存到缓存
	return t.cacheToken.Save(tk)
}

//Refresh 刷新当前token
func (t *Token) Refresh(force bool) (token *AccessToken, r bool, err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if !force {
		token, err := t.dbToken.Get()
		if err == nil {
			return token, false, nil
		}
	}

	//获取secret用于通过微信官方获取token
	secret, err := t.dbToken.GetSecret()
	if err != nil {
		return
	}
	token, err = t.wxToken.Get(secret)
	if err != nil {
		t.dbToken.Update(err.Error())
		return
	}

	//保存token
	if err = t.save(token); err != nil {
		return
	}
	return token, true, nil
}

//Query 从缓存中获取access token，不存在时自动通过官网刷新到本地
func (t *Token) Query() (token *AccessToken, err error) {
	return t.dbToken.Query()
}

//Get 从缓存中获取access token，不存在时自动通过官网刷新到本地
func (t *Token) Get() (token *AccessToken, err error) {
	//从本地缓存获取access token
	token, err = t.cacheToken.Get()
	if err == nil {
		return token, nil
	}

	//从数据库获取access token,并保存到本地
	token, err = t.dbToken.Get()
	if err == nil {
		err = t.cacheToken.Save(token)
		return token, err
	}
	token, _, err = t.Refresh(true)
	return
}

func (a *Token) GetContext() (*mp.Context, error) {
	_, ctx, err := appTokens.SetIfAbsentCb(a.appid, func(input ...interface{}) (interface{}, error) {
		return mp.NewContext(&accessToken{token: a}), nil
	})
	return ctx.(*mp.Context), err
}

var appTokens cmap.ConcurrentMap

func init() {
	appTokens = cmap.New(2)
}

type accessToken struct {
	token *Token
}

func (a *accessToken) Token() (token string, err error) {
	tk, err := a.token.dbToken.Get()
	if err != nil {
		return "", err
	}
	return tk.Token, nil
}
func (a *accessToken) RefreshToken(currentToken string) (token string, err error) {
	return a.Token()
}
