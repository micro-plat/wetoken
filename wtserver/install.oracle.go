// +build oracle

package main

import (
	_ "github.com/mattn/go-oci8"
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
	// hydra.Conf.Vars().HTTP("http")
	hydra.Conf.Vars().RLog("/log/save@logging")
}

//测试环境配置
func devConf() {
	appConf := &AppWXConf{WX: []WXConfig{WXConfig{
		AppID: "wx5260e02d76f306ca",
	}}}

	hydra.Conf.API("9999").Sub("app", appConf)
	hydra.Conf.CRON().Task(task.NewTask("@every 1m", "/wx5260e02d76f306ca/wechat/token/refresh"), task.NewTask("@every 1m", "/wx5260e02d76f306ca/wechat/ticket/refresh")).Sub("app", appConf)
	hydra.Conf.Vars().DB().Oracle("db", "wechat_v2", "123456", "orcl136", db.WithConnect(20, 10, 600))
}

//生产环境配置
func prodConf() {
	appConf := &AppWXConf{WX: []WXConfig{WXConfig{
		AppID: "#appid",
	}}}
	hydra.Conf.API("9999").Sub("app", appConf)
	hydra.Conf.CRON().Task(task.NewTask("@every 1m", "/###appid/wechat/token/refresh"), task.NewTask("@every 1m", "/###appid/wechat/ticket/refresh")).Sub("app", appConf)
	hydra.Conf.Vars().DB().OracleByConnStr("db", "###oracle_db_string", db.WithConnect(20, 10, 600))
}
