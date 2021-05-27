// +build test
// +build mysql

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/vars/db"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	hydra.OnReady(func() {
		//配置共有配置
		hydra.Conf.Vars().RLog("/log/save@logging")
		hydra.Conf.API("8080", api.WithDNS("api.wetoken.18jiayou0.com"))
		hydra.Conf.CRON()
		hydra.Conf.Vars().DB().MySQL("db", "root", "rTo0CesHi2018Qx", "192.168.0.36:3306", "wechat_v2", db.WithConnect(10, 10, 600))
	})
}
