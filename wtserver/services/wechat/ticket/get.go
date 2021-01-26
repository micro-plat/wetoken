package ticket

import (
	"fmt"
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/wetoken/modules/wechat/ticket"
)

type GetTicketHandler struct {
}

//NewGetTicketHandler 创建服务
func NewGetTicketHandler() (u *GetTicketHandler) {
	return &GetTicketHandler{}
}

//Handle 获取jsapi ticket
//1. 从缓存中获取，不存在或过期时从数据库中获取
//2. 从数据库中获取，不存在或过期时从微信官网获取
//3. 从微信官网获取成功后，更新本地缓存和数据库
func (u *GetTicketHandler) Handle(ctx hydra.IContext) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		ticket.JSTicket
	}

	appid := ctx.Request().Path().Params().GetString("appid")
	if appid == "" {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("参数appid错误,%s", appid))
	}
	ticketObj := ticket.NewTicket(appid)
	ticketInfo, err := ticketObj.Get()
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	result.JSTicket = *ticketInfo
	result.JSTicket.Reset()
	return result
}
