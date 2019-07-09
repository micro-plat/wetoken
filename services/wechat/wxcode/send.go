package wxcode

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wetoken/modules/wechat/app"
	"github.com/micro-plat/wetoken/modules/wechat/message"
)

type SendHandler struct {
	container component.IContainer
	appid     string
	app       app.IWechatApp
	msg       message.IMessage
}

func NewSendHandlerBy(appid string) func(container component.IContainer) (u *SendHandler) {
	return func(container component.IContainer) (u *SendHandler) {
		return &SendHandler{
			container: container,
			appid:     appid,
			msg:       message.NewMessage(container, appid),
		}
	}
}

func (u *SendHandler) Handle(ctx *context.Context) (r interface{}) {
	if err := ctx.Request.Check("user_name", "system_name", "code"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	if err := u.msg.SendWXCode(
		ctx.Request.GetString("user_name"),
		ctx.Request.GetString("system_name"),
		ctx.Request.GetString("code")); err != nil {
		return err
	}
	return "success"

}
