package token

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/const/sql"
)

//DBToken 数据库access token
type DBToken struct {
	appid string
}

//NewDBToken 构建数据库access token访问对象
func NewDBToken(appid string) *DBToken {
	return &DBToken{
		appid: appid,
	}
}

//Update 修改access token查询消息
func (t *DBToken) Update(msg string) error {
	db := hydra.C.DB().GetRegularDB()
	_, err := db.Execute(sql.UpdateAccessTokenMessage, map[string]interface{}{
		"appid":   t.appid,
		"message": msg,
	})
	if err != nil {
		return err
	}
	return nil
}

//Save 保存access token到db,不存在时创建
func (t *DBToken) Save(tk *AccessToken) error {
	db := hydra.C.DB().GetRegularDB()
	//更新access token
	r, err := db.Execute(sql.UpdateAccessToken, map[string]interface{}{
		"appid":   t.appid,
		"token":   tk.Token,
		"expire":  tk.ExpiresIn,
		"message": "success",
	})
	if err != nil {
		return err
	}
	if r == 1 {
		return nil
	}
	//创建access token记录
	_, err = db.Execute(sql.InsertAccessToken, map[string]interface{}{
		"appid":   t.appid,
		"token":   tk.Token,
		"expire":  tk.ExpiresIn,
		"message": "success",
	})
	if err != nil {
		return err
	}
	return nil
}

//Query 从数据库获取access Token
func (t *DBToken) Query() (token *AccessToken, err error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(sql.QueryAccessToken, map[string]interface{}{"appid": t.appid})
	if err != nil {
		return nil, fmt.Errorf("从数据库中获取token失败:%s %v", t.appid, err)
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("从数据库中获取token失败:%s app信息未配置", t.appid)
	}

	//处理过期时间
	expire, err := data.Get(0).GetDatetime("expire_date")
	if err != nil {
		return nil, err
	}
	token = &AccessToken{
		Token:       data.Get(0).GetString("access_token"),
		ExpiresIn:   data.Get(0).GetInt64("expire"),
		ExpiresDate: jtime(expire),
	}
	return token, nil
}

//Get 从数据库获取access Token
func (t *DBToken) Get() (token *AccessToken, err error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(sql.QueryAccessToken, map[string]interface{}{"appid": t.appid})
	if err != nil {
		return nil, fmt.Errorf("从数据库中获取token失败:%s %v", t.appid, err)
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("从数据库中获取token失败:%s app信息未配置", t.appid)
	}

	//处理过期时间
	expire, err := data.Get(0).GetDatetime("expire_date")
	if err != nil {
		return nil, err
	}
	//检查是否过期
	if time.Since(expire) >= 0 {
		return nil, errTokenExpired
	}
	token = &AccessToken{
		Token:       data.Get(0).GetString("access_token"),
		ExpiresIn:   data.Get(0).GetInt64("expire"),
		ExpiresDate: jtime(expire),
	}
	//更新expire
	token.Reset()
	return token, nil
}

//GetSecret 获取secret
func (t *DBToken) GetSecret() (secret string, err error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(sql.QueryWechatApp, map[string]interface{}{"appid": t.appid})
	if err != nil {
		return "", fmt.Errorf("从数据库中获取secret失败:%s %v", t.appid, err)
	}
	if len(data) != 1 {
		return "", fmt.Errorf("从数据库中获取secret失败:[%s] app配置出错(%d)", t.appid, len(data))
	}
	return data.Get(0).GetString("secret"), nil
}
