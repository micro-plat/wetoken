package app

import (
	"fmt"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/const/sql"
)

type AppBaseInfo struct {
	Appid    string `form:"appid" json:"appid" valid:"required"`
	Secret   string `form:"secret" json:"secret" valid:"required"`
	Token    string `form:"token" json:"token" valid:"required"`
	AesKey   string `form:"aesKey" json:"aesKey"`
	Mchid    string `form:"mchid" json:"mchid"`
	PayKey   string `form:"payKey" json:"payKey"`
	SubAppid string `form:"subAppid" json:"subAppid"`
	SubMchid string `form:"subMchid" json:"subMchid"`
}

//IWechatApp 微信app操作
type IWechatApp interface {
	//Get 根据appid获取当前微信app信息
	Get(appid string) (*WechatAppInfo, error)

	//Save 保存微信基本信息
	Save(b AppBaseInfo) error
}

//WechatApp 微信信息
type WechatApp struct {
}

//NewWechatApp 构建wechat app操作对象
func NewWechatApp() *WechatApp {
	return &WechatApp{}
}

//Get 根据appid获取当前微信app信息
func (a *WechatApp) Get(appid string) (*WechatAppInfo, error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(sql.QueryWechatApp, map[string]interface{}{"appid": appid})
	if err != nil {
		return nil, err
	}
	app := WechatAppInfo{}
	res := data.Get(0)
	if err = res.ToStruct(&app); err != nil {
		return nil, err
	}
	return &app, nil
}

//Save 保存微信基本信息
func (a *WechatApp) Save(b AppBaseInfo) error {
	db := hydra.C.DB().GetRegularDB()
	d, err := db.Query(sql.QueryWechatApp, map[string]interface{}{
		"appid": b.Appid,
	})
	if err != nil {
		return err
	}
	if d.Len() > 0 {
		err = fmt.Errorf("%s已经存在", b.Appid)
		return err
	}

	_, err = db.Execute(sql.InsertWebchatApp,
		map[string]interface{}{
			"appid":     b.Appid,
			"secret":    b.Secret,
			"token":     b.Token,
			"aes_key":   b.AesKey,
			"mchid":     b.Mchid,
			"pay_key":   b.PayKey,
			"sub_appid": b.SubAppid,
			"sub_mchid": b.SubMchid,
		})
	if err != nil {
		return err
	}
	return nil
}
