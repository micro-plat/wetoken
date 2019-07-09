package ticket

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

//CacheTicket 本地缓存的jsapi ticket操作对象
type CacheTicket struct {
	cache     *cache.Cache
	appid     string
	accessKey string
}

//NewCacheTicket 构建基于本地缓存的jsapi ticket操作对象
func NewCacheTicket(appid string) *CacheTicket {

	return &CacheTicket{
		cache:     cache.New(5*time.Minute, 10*time.Minute),
		accessKey: fmt.Sprintf("jsapi-ticket:%s", appid),
		appid:     appid,
	}
}

//Save 保存jsapi ticket到缓存中
func (t *CacheTicket) Save(tk *JSTicket) error {
	tk.Reset() //根据过期时间重置ExpiresIn
	t.cache.Set(t.accessKey, tk, time.Second*time.Duration(tk.ExpiresIn))
	return nil
}

//Get 从缓存中获取jsapi ticket
func (t *CacheTicket) Get() (CacheTicket *JSTicket, err error) {
	v, b := t.cache.Get(t.accessKey)
	if !b {
		return nil, errTicketDontExist
	}
	//检查是否过期
	if time.Since(time.Time(v.(*JSTicket).ExpiresDate)) >= 0 {
		return nil, errTicketExpired
	}
	CacheTicket = v.(*JSTicket)
	CacheTicket.Reset()
	return
}
