package users

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/wetoken/modules/const/sql"
)

type IDBUser interface {
	Get(appid string, userName string) (*WechatUserInfo, error)
}

type User struct {
	c component.IContainer
}

func NewDBUser(c component.IContainer) *User {
	return &User{
		c: c,
	}
}

//Get 根据appid,用户名获取用户信息
func (a *User) Get(appid string, userName string) (*WechatUserInfo, error) {
	db := a.c.GetRegularDB()
	data, _, _, err := db.Query(sql.QueryUserByName, map[string]interface{}{
		"appid":     appid,
		"user_name": userName,
	})
	if err != nil {
		return nil, err
	}
	u := WechatUserInfo{}
	if err = types.Map2Struct(data.Get(0), &u); err != nil {
		return nil, err
	}
	return &u, nil
}
