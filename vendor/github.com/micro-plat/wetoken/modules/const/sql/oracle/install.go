package oracle

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/hydra"
)

//Install 安装mysql数据库
func Install(app *hydra.MicroApp, c component.IContainer) error {
	if !app.Conf.Confirm("创建数据库表结构,添加基础数据?") {
		return nil
	}
	path, err := getSQLPath()
	if err != nil {
		return err
	}
	sqls, err := app.Conf.GetSQL(path)
	if err != nil {
		return err
	}
	db, err := c.GetDB()
	if err != nil {
		return err
	}
	for _, sql := range sqls {
		if sql != "" {
			if _, q, _, err := db.Execute(sql, map[string]interface{}{}); err != nil {
				if !strings.Contains(err.Error(), "1050") && !strings.Contains(err.Error(), "1061") && !strings.Contains(err.Error(), "1091") {
					app.Conf.Log.Errorf("执行SQL失败： %v %s\n", err, q)
				}
			}
		}
	}
	return nil
}

//getSQLPath 获取getSQLPath
func getSQLPath() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", fmt.Errorf("未配置环境变量GOPATH")
	}
	path := strings.Split(gopath, ";")
	if len(path) == 0 {
		return "", fmt.Errorf("环境变量GOPATH配置的路径为空")
	}
	return filepath.Join(path[0], "src/github.com/micro-plat/wetoken/modules/const/sql/oracle"), nil
}
