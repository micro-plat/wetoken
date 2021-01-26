package message

import (
	"fmt"

	"github.com/micro-plat/wetoken/modules/wechat/app"
	"github.com/micro-plat/wetoken/modules/wechat/token"
	"github.com/micro-plat/wetoken/modules/wechat/users"
)

type IMessage interface {
	SendWXCode(userName string, sysName string, code string) error
}

type Message struct {
	app   app.IWechatApp
	token token.IToken
	user  users.IDBUser
	appid string
}

func NewMessage(appid string) *Message {
	return &Message{
		app:   app.NewWechatApp(),
		user:  users.NewDBUser(),
		token: token.NewToken(appid),
		appid: appid,
	}
}

//SendWXCode 发送微信验证码
func (v *Message) SendWXCode(userName string, sysName string, code string) error {
	app, err := v.app.Get(v.appid)
	if err != nil {
		return err
	}
	if app.WcodeTemplateID == "" {
		return fmt.Errorf("未配置微信验证码模板编号")
	}
	user, err := v.user.Get(v.appid, userName)
	if err != nil {
		return err
	}
	if user.OpenID == "" {
		return fmt.Errorf("用户：%s未绑定公众号", userName)
	}
	ctx, err := v.token.GetContext()
	if err != nil {
		return err
	}
	return wxSend(v.appid, user.OpenID, sysName, app.WcodeTemplateID, code, ctx)
}
