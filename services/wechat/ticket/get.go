package ticket

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wetoken/modules/wechat/ticket"
)

type GetHandler struct {
	container component.IContainer
	appid     string
	ticket    ticket.ITicket
}

func NewGetHandlerBy(appid string) func(container component.IContainer) (u *GetHandler) {
	return func(container component.IContainer) (u *GetHandler) {
		return &GetHandler{
			container: container,
			appid:     appid,
			ticket:    ticket.NewTicket(container, appid),
		}
	}
}

//NewGetHandler 创建服务
func NewGetHandler(container component.IContainer) (u *GetHandler) {
	return &GetHandler{
		container: container,
	}
}

//Handle 获取jsapi ticket
//1. 从缓存中获取，不存在或过期时从数据库中获取
//2. 从数据库中获取，不存在或过期时从微信官网获取
//3. 从微信官网获取成功后，更新本地缓存和数据库
func (u *GetHandler) Handle(ctx *context.Context) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		ticket.JSTicket
	}
	ticket, err := u.ticket.Get()
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	result.JSTicket = *ticket
	result.JSTicket.Reset()
	return result
}
