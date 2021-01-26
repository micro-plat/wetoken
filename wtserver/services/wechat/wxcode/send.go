package wxcode

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/wetoken/modules/wechat/app"
	"github.com/micro-plat/wetoken/modules/wechat/message"
	"net/http"
)

type SendHandler struct {
	appid string
	app   app.IWechatApp
	msg   message.IMessage
}

func NewSendHandlerBy(appid string) func() (u *SendHandler) {
	return func() (u *SendHandler) {
		return &SendHandler{
			appid: appid,
			msg:   message.NewMessage(appid),
		}
	}
}

func (u *SendHandler) Handle(ctx hydra.IContext) (r interface{}) {
	if err := ctx.Request().Check("user_name", "system_name", "code"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}
	if err := u.msg.SendWXCode(
		ctx.Request().GetString("user_name"),
		ctx.Request().GetString("system_name"),
		ctx.Request().GetString("code")); err != nil {
		return err
	}
	return "success"

}
