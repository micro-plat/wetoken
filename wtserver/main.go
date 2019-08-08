package main

import "github.com/micro-plat/hydra/hydra"

type wtserver struct {
	*hydra.MicroApp
}

func main() {
	app := &wtserver{
		hydra.NewApp(
			hydra.WithPlatName("wetoken"),
			hydra.WithSystemName("wtserver"),
			hydra.WithServerTypes("api-cron")),
	}

	app.init()
	app.install()
	app.installDB()
	app.Start()
}
