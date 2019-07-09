// +build prod

package main

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func (app *wtserver) install() {
	app.IsDebug = false
	app.Conf.API.SetMainConf(`{"address":":9999"}`)
	app.Conf.API.SetSubConf("app", `{
		"wx":[{
			"appid": "wx5260e02d76f306ca"
		}]	
	}`)
	app.Conf.CRON.SetSubConf("app", `{
		"wx":[{
			"appid": "wx5260e02d76f306ca"
		}]	
	}`)
	app.Conf.CRON.SetSubConf("task", `{
		"tasks":[
		{"cron":"@every 1m","service":"/wx5260e02d76f306ca/wechat/token/refresh"},
		{"cron":"@every 1m","service":"/wx5260e02d76f306ca/wechat/ticket/refresh"}
		]		
		}`)

	app.Conf.Plat.SetVarConf("db", "db", `{			
				"provider":"mysql",
				"connString":"wechat:12345678@tcp(192.168.0.36)/wechat_v2",
				"maxOpen":10,
				"maxIdle":1,
				"lifeTime":10		
		}`)

}
