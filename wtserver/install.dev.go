// +build !prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (app *wtserver) install() {
	app.IsDebug = true
	app.Conf.API.SetMain(conf.NewAPIServerConf(":9999"))
	appConf := &AppWXConf{WX: []WXConfig{WXConfig{
		AppID: "wx5260e02d76f306ca",
	}}}
	app.Conf.API.SetApp(appConf)
	app.Conf.CRON.SetApp(appConf)
	app.Conf.CRON.SetTasks(conf.NewTasks().Append("@every 1m", "/wx5260e02d76f306ca/wechat/token/refresh").Append("@every 1m", "/wx5260e02d76f306ca/wechat/ticket/refresh"))
	app.Conf.Plat.SetDB(conf.NewMysqlConf("wechat", "12345678", "192.168.0.36", "wechat_v2"))

}
