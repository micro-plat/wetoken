// +build !oci

package sql

import "github.com/micro-plat/lib4go/db"

import _ "github.com/go-sql-driver/mysql"

func CreateDB(xdb db.IDB) error {
	return db.CreateDB(xdb, "src/github.com/micro-plat/wetoken/modules/const/sql/mysql")
}
