package ticket

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/wechat/ticket"
)

//RefreshTicketHandler token刷新操作
type RefreshTicketHandler struct {
}

//NewRefreshTicketHandler 创建服务
func NewRefreshTicketHandler() (u *RefreshTicketHandler) {
	return &RefreshTicketHandler{}
}

//Handle 刷新jsapi ticket
//1. 检查数据库jsapi ticket是否存在和过期
//2. 不存在或过期，从微信官方获取后更新到本地缓存和数据库
func (u *RefreshTicketHandler) Handle(ctx hydra.IContext) (r interface{}) {

	ticketObj := ticket.NewTicket("")
	_, b, err := ticketObj.Refresh(false)
	if err != nil {
		return err
	}
	if !b {
		ctx.Response().Write(204)
		return
	}
	ctx.Log().Infof("appid:%s jsapi ticket 刷新成功", "appid")
	return nil
}
