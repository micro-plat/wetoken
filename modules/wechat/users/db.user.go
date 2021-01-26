package users

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/const/sql"
)

type IDBUser interface {
	Get(appid string, userName string) (*WechatUserInfo, error)
}

type User struct {
}

func NewDBUser() *User {
	return &User{}
}

//Get 根据appid,用户名获取用户信息
func (a *User) Get(appid string, userName string) (*WechatUserInfo, error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(sql.QueryUserByName, map[string]interface{}{
		"appid":     appid,
		"user_name": userName,
	})
	if err != nil {
		return nil, err
	}
	u := WechatUserInfo{}
	res := data.Get(0)
	if err = res.ToStruct(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
