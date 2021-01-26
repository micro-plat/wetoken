// +build !oracle

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/vars/db"
)

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func install() {
	hydra.OnReady(func() error {
		//配置共有配置
		pubConf()

		if hydra.G.IsDebug() {
			//测试环境配置
			devConf()
			return nil
		}

		//生产环境的配置
		prodConf()
		return nil
	})
}

//公共配置
func pubConf() {
	hydra.Conf.Vars().RLog("/log/save@logging")
}

//测试环境配置
func devConf() {
	hydra.Conf.API("9999")
	hydra.Conf.CRON()
	hydra.Conf.Vars().DB().MySQL("db", "wechat", "12345678", "192.168.0.36:3306", "wechat_v2", db.WithConnect(20, 10, 600))
}

//生产环境配置
func prodConf() {
	hydra.Conf.API("9999")
	hydra.Conf.CRON()
	hydra.Conf.Vars().DB().MySQLByConnStr("db", "###mysql_db_string", db.WithConnect(20, 10, 600))
}
