// +build oci

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
	"github.com/micro-plat/wetoken/modules/const/sql/oracle"
	"github.com/urfave/cli"
)

func (app *wtserver) installDB() { //自定义安装程序

	app.Conf.API.Installer(func(c component.IContainer) error {
		app.Cli.Validate(hydra.ModeRegistry, func(t *cli.Context) error {
			if t.IsSet("db") && t.GlobalBool("db") {
				return oracle.Install(app.MicroApp, c)
			}
			return nil
		})
		return nil
	})
}
