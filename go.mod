module github.com/micro-plat/wetoken

go 1.16

replace github.com/micro-plat/hydra => ../../../github.com/micro-plat/hydra

replace github.com/micro-plat/lib4go => ../../../github.com/micro-plat/lib4go

replace github.com/lib4dev/wechat => ../../../github.com/lib4dev/wechat

require (
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef
	github.com/go-sql-driver/mysql v1.4.1
	github.com/lib4dev/wechat v0.0.0-00010101000000-000000000000
	github.com/mattn/go-oci8 v0.1.1
	github.com/micro-plat/hydra v0.0.0-00010101000000-000000000000
	github.com/micro-plat/lib4go v1.0.10
	github.com/micro-plat/wechat v0.0.0-20190124071457-da29af712665 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
)
