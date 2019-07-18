package ticket

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wetoken/modules/wechat/ticket"
)

//RefreshHandler token刷新操作
type RefreshHandler struct {
	container component.IContainer
	appid     string
	ticket    ticket.ITicket
}

//NewRefreshHandlerBy 根据appid创建定时刷新对象
func NewRefreshHandlerBy(appid string) func(container component.IContainer) (u *RefreshHandler) {
	return func(container component.IContainer) (u *RefreshHandler) {
		return &RefreshHandler{
			container: container,
			appid:     appid,
			ticket:    ticket.NewTicket(container, appid),
		}
	}
}

//NewRefreshHandler 创建服务
func NewRefreshHandler(container component.IContainer) (u *RefreshHandler) {
	return &RefreshHandler{
		container: container,
	}
}

//Handle 刷新jsapi ticket
//1. 检查数据库jsapi ticket是否存在和过期
//2. 不存在或过期，从微信官方获取后更新到本地缓存和数据库
func (u *RefreshHandler) Handle(ctx *context.Context) (r interface{}) {
	_, b, err := u.ticket.Refresh(false)
	if err != nil {
		return err
	}
	if !b {
		ctx.Response.SetStatus(204)
		return
	}
	ctx.Log.Infof("appid:%s jsapi ticket 刷新成功", u.appid)
	return nil
}
