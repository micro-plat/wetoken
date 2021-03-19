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

	appids, err := token.GetList()
	if err != nil {
		return err
	}
	if len(appids) <= 0 {
		ctx.Log().Debug("没有需要刷新的token")
		return
	}

	for _, appid := range appids {
		tokenObj := token.NewToken(appid)
		_, b, err := tokenObj.Refresh(false)
		if err != nil {
			ctx.Log().Errorf("appid:%s access token 刷新失败,err:%v", appid, err)
			continue
		}
		if !b {
			ctx.Log().Debugf("appid:%s access token 不用刷新", appid)
			continue
		}
		ctx.Log().Infof("appid:%s access token 刷新成功", appid)
	}
	return nil
}
