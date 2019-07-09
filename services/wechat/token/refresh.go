package token

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

//RefreshHandler token刷新操作
type RefreshHandler struct {
	container component.IContainer
	appid     string
	token     token.IToken
}

//NewRefreshHandlerBy 根据appid创建定时刷新对象
func NewRefreshHandlerBy(appid string) func(container component.IContainer) (u *RefreshHandler) {
	return func(container component.IContainer) (u *RefreshHandler) {
		return &RefreshHandler{
			container: container,
			appid:     appid,
			token:     token.NewToken(container, appid),
		}
	}
}

//NewRefreshHandler 创建服务
func NewRefreshHandler(container component.IContainer) (u *RefreshHandler) {
	return &RefreshHandler{
		container: container,
	}
}

//Handle 刷新access token
//1. 检查数据库access token是否存在和过期
//2. 不存在或过期，从微信官方获取后更新到本地缓存和数据库
func (u *RefreshHandler) Handle(ctx *context.Context) (r interface{}) {
	_, b, err := u.token.Refresh(false)
	if err != nil {
		return err
	}
	if !b {
		ctx.Response.SetStatus(204)
		return
	}
	ctx.Log.Infof("appid:%s access token 刷新成功", u.appid)
	return nil
}
