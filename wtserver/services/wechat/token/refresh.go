package token

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

//RefreshTokenHandler token刷新操作
type RefreshTokenHandler struct {
}

//NewRefreshTokenHandler 创建服务
func NewRefreshTokenHandler() (u *RefreshTokenHandler) {
	return &RefreshTokenHandler{}
}

//Handle 刷新access token
//1. 检查数据库access token是否存在和过期
//2. 不存在或过期，从微信官方获取后更新到本地缓存和数据库
func (u *RefreshTokenHandler) Handle(ctx hydra.IContext) (r interface{}) {

	tokenObj := token.NewToken("")
	_, b, err := tokenObj.Refresh(false)
	if err != nil {
		return err
	}
	if !b {
		ctx.Response().Write(204)
		return
	}
	ctx.Log().Infof("appid:%s access token 刷新成功", "appid")
	return nil
}
