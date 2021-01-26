package ticket

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/micro-plat/wetoken/modules/wechat/token"
)

var errTicketExpired = errors.New("jsapi ticket has expired")
var errTicketDontExist = errors.New("jsapi ticket does not exist")

const (
	ErrCodeOK = 0
)

type jtime time.Time

func (this jtime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("20060102150405"))
	return []byte(stamp), nil
}

type JSTicket struct {
	Ticket      string `json:"ticket"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresDate jtime  `json:"expires_date"`
}

func (a *JSTicket) Reset() {
	a.ExpiresIn = -1 * int64(time.Since(time.Time(a.ExpiresDate)).Seconds())
}

//ITicket jsapi ticket操作接口
type ITicket interface {
	//刷新Ticket
	Refresh(force bool) (Ticket *JSTicket, r bool, err error)

	//获取Ticket
	Get() (Ticket *JSTicket, err error)
}

//Ticket jsapi ticket统一访问对象
type Ticket struct {
	dbTicket    *DBTicket
	cacheTicket *CacheTicket
	wxTicket    *WechatTicket
	token       *token.Token
	lock        sync.Mutex
}

//NewTicket 构建jsapi ticket统一访问对象
func NewTicket(appid string) *Ticket {
	return &Ticket{
		dbTicket:    NewDBTicket(appid),
		cacheTicket: NewCacheTicket(appid),
		wxTicket:    NewWechatTicket(appid),
		token:       token.NewToken(appid),
	}
}

//Save 保存jsapi ticket到数据库和缓存
func (t *Ticket) save(tk *JSTicket) error {
	//保存到数据库
	if err := t.dbTicket.Save(tk); err != nil {
		return err
	}
	//保存到缓存
	return t.cacheTicket.Save(tk)
}

//Refresh 刷新当前ticket
func (t *Ticket) Refresh(force bool) (Ticket *JSTicket, r bool, err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if !force {
		Ticket, err := t.dbTicket.Get()
		if err == nil {
			return Ticket, false, nil
		}
	}

	//获取token用于通过微信官方获取ticket
	token, err := t.token.Get()
	if err != nil {
		return
	}
	Ticket, err = t.wxTicket.Get(token.Token)
	if err != nil {
		t.dbTicket.Update(err.Error())
		return
	}

	//保存Ticket
	if err = t.save(Ticket); err != nil {
		return
	}
	return Ticket, true, nil
}

//Get 从缓存中获取jsapi ticket，不存在时自动通过官网刷新到本地
func (t *Ticket) Get() (Ticket *JSTicket, err error) {
	//从本地缓存获取jsapi ticket
	Ticket, err = t.cacheTicket.Get()
	if err == nil {
		return Ticket, nil
	}

	//从数据库获取jsapi ticket,并保存到本地
	Ticket, err = t.dbTicket.Get()
	if err == nil {
		err = t.cacheTicket.Save(Ticket)
		return Ticket, err
	}
	Ticket, _, err = t.Refresh(true)
	return
}
