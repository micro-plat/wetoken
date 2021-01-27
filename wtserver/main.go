package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("wetoken", "微信token维护服务"),
	hydra.WithSystemName("wtserver"),
	hydra.WithServerTypes(http.API, cron.CRON),
	hydra.WithClusterName("prod"))

func main() {
	App.Start()
}
