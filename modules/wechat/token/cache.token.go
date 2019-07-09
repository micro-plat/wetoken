package token

import (
	"fmt"
	"time"

	cache "github.com/patrickmn/go-cache"
)

//CacheToken 本地缓存的acess token操作对象
type CacheToken struct {
	cache     *cache.Cache
	appid     string
	accessKey string
}

//NewCacheToken 构建基于本地缓存的acess token操作对象
func NewCacheToken(appid string) *CacheToken {

	return &CacheToken{
		cache:     cache.New(5*time.Minute, 10*time.Minute),
		accessKey: fmt.Sprintf("access-token:%s", appid),
		appid:     appid,
	}
}

//Save 保存access token到缓存中
func (t *CacheToken) Save(tk *AccessToken) error {
	tk.Reset() //根据过期时间重置ExpiresIn
	t.cache.Set(t.accessKey, tk, time.Second*time.Duration(tk.ExpiresIn))
	return nil
}

//Get 从缓存中获取access token
func (t *CacheToken) Get() (cacheToken *AccessToken, err error) {
	v, b := t.cache.Get(t.accessKey)
	if !b {
		return nil, errTokenDontExist
	}
	//检查是否过期
	if time.Since(time.Time(v.(*AccessToken).ExpiresDate)) >= 0 {
		return nil, errTokenExpired
	}
	cacheToken = v.(*AccessToken)
	cacheToken.Reset()
	return
}
