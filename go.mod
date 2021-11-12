module github.com/micro-plat/wetoken

go 1.16

//replace (
//	github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra
//	github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go
//	github.com/lib4dev/wechat => ../../../github.com/lib4dev/wechat
//)

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lib4dev/wechat v1.0.0
	github.com/mattn/go-oci8 v0.1.1
	github.com/micro-plat/hydra v1.2.0
	github.com/micro-plat/lib4go v1.1.10
	github.com/patrickmn/go-cache v2.1.0+incompatible
)
