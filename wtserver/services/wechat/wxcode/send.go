package wxcode

import (
	"fmt"
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/wetoken/modules/wechat/app"
	"github.com/micro-plat/wetoken/modules/wechat/message"
)

type SendMessgHandler struct {
	appid string
	app   app.IWechatApp
	msg   message.IMessage
}

func NewSendMessgHandler() (u *SendMessgHandler) {
	return &SendMessgHandler{}
}

func (u *SendMessgHandler) Handle(ctx hydra.IContext) (r interface{}) {
	if err := ctx.Request().Check("user_name", "system_name", "code"); err != nil {
		return errs.NewError(http.StatusNotAcceptable, err)
	}

	appid := ctx.Request().Path().Params().GetString("appid")
	if appid == "" {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("参数appid错误,%s", appid))
	}

	messageObj := message.NewMessage(appid)
	if err := messageObj.SendWXCode(
		ctx.Request().GetString("user_name"),
		ctx.Request().GetString("system_name"),
		ctx.Request().GetString("code")); err != nil {
		return err
	}
	return "success"

}
