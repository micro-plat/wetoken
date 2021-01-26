package ticket

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/wechat/ticket"
)

//RefreshHandler token刷新操作
type RefreshHandler struct {
	appid  string
	ticket ticket.ITicket
}

//NewRefreshHandlerBy 根据appid创建定时刷新对象
func NewRefreshHandlerBy(appid string) func() (u *RefreshHandler) {
	return func() (u *RefreshHandler) {
		return &RefreshHandler{
			appid:  appid,
			ticket: ticket.NewTicket(appid),
		}
	}
}

//NewRefreshHandler 创建服务
func NewRefreshHandler() (u *RefreshHandler) {
	return &RefreshHandler{}
}

//Handle 刷新jsapi ticket
//1. 检查数据库jsapi ticket是否存在和过期
//2. 不存在或过期，从微信官方获取后更新到本地缓存和数据库
func (u *RefreshHandler) Handle(ctx hydra.IContext) (r interface{}) {
	_, b, err := u.ticket.Refresh(false)
	if err != nil {
		return err
	}
	if !b {
		ctx.Response().Write(204)
		return
	}
	ctx.Log().Infof("appid:%s jsapi ticket 刷新成功", u.appid)
	return nil
}
