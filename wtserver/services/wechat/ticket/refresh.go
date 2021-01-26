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
	appids, err := ticket.GetList()
	if err != nil {
		return err
	}
	if len(appids) <= 0 {
		ctx.Log().Infof("没有需要刷新的ticket")
		return
	}

	for _, appid := range appids {
		ticketObj := ticket.NewTicket(appid)
		_, b, err := ticketObj.Refresh(false)
		if err != nil {
			ctx.Log().Errorf("appid:%s jsapi ticket 刷新失败,err:%v", appid, err)
			continue
		}
		if !b {
			ctx.Log().Infof("appid:%s jsapi ticket 不用刷新", appid)
			continue
		}
		ctx.Log().Infof("appid:%s jsapi ticket 刷新成功", appid)
	}
	return nil
}
