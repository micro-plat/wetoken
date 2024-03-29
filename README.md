# wetoken
公众号token服务，统一管理token，jsticket，过期后自动从官网获取，保存到数据库、本地缓存，并对外提供获取、刷新接口。


√ 　支持独立部署

√ 　支持代码集成

√ 　代码、接口添加公众号

√ 　代码、接口获取token，jsticket

√ 　获取时自动检查过期

√ 　定时从官网获取

√ 　基于 hydra 构建




## 一、独立服务


多个系统共用一个公众号，可使用wtserver，集中维护token，jstikcet。根据向导即可安装数据表结构，启动后即可对外提供服务。


 安装数据表

 ```sh
go install -tags "dev" github.com/micro-plat/wetoken/wtserver # mysql
go install -tags "dev oracle" github.com/micro-plat/wetoken/wtserver # oracle
wtserver db install -r zk://192.168.106.18
 ```


编译安装

```sh
go install github.com/micro-plat/wetoken/wtserver # mysql

 或

go install -tags "oracle" github.com/micro-plat/wetoken/wtserver # oracle
```


启动服务
```sh
wtserver run -r zk://192.168.106.18

```


## 二、代码集成
只有一个系统使用公众号，为减少子系统维护，可直接通过代码集成到现有系统中。


## 1、创建任务表

#### 1. 引入 `wetoken`

```go
package main

import "github.com/micro-plat/wetoken/wetoken"
```

#### 2. 运行命令

`server [注册中心地址]` 即可将 `wetoken` 需要的表创建到`/平台/var/db/db` 配置对应的数据库中。

```sh
go install -tags "dev" # mysql
go install -tags "dev oracle" # oracle
server db install -r zk://192.168.106.18
```

## 2、编码

绑定服务

```go

App.OnStarting(func(appConf app.IAPPConf) error {
  err := wetoken.Bind(false)　//绑定公众号token,ticket对应的刷新，获取，消息推送接口等
  return err
})

或

App.OnStarting(func(appConf app.IAPPConf) error {
  err := wetoken.Cron(false)　//只绑定定时任务，每隔1分钟检查一次token,ticket
  return err
})
```



> 当绑定参数为`true`时会自动绑定公众号添加服务，可通过接口添加公众号

获取token

```go

token, err:=wetoken.GetWxToken(appid) //不存在或已过期，自动从官网获取

```

获取ticket

```go

ticket, err:=wetoken.GetWxTicket(appid) //不存在或已过期，自动从官网获取

```

获取secret

```go

secret, err:=wetoken.GetWxSecret(appid)

```
## 三、微信 sdk 中使用 token

```go
package main

import "github.com/micro-plat/wechat/mp"
import "github.com/micro-plat/wechat/mp/menu"


func main(){
  tk := mp.NewDefaultAccessTokenByURL(
  "wx9e02ddcc34513fd4"， "6acb2b999177524beba3d97d54df2de5"， 
  "http://localhost:9999/wx9e02ddcc34513fd4/wechat/token/get")
  ctx := mp.NewContext(tk)
  mu := &menu.Menu{
    Buttons: []menu.Button{
      menu.Button{Type: menu.ButtonTypeView， Name: "搜索1"， URL: "http://www.baidu.com"}，
      menu.Button{Type: menu.ButtonTypeView， Name: "搜索2"， URL: "http://www.baidu.com"}，
      menu.Button{Type: menu.ButtonTypeView， Name: "搜索3"， URL: "http://www.baidu.com"}，
    }，
  }
  if err := menu.Create(ctx， mu);err!=nil{
    fmt.Println(err)
  }
}
```



## 四、其它


##### 服务列表

| 服务名                     | 类型  | wetoken.Bind | wetoken.Cron | 说明               |
| :------------------------- | :---: | :----------: | :----------: | :----------------- |
| /wechat/app/create         |  api  |     可选     |     可选     | 添加公众号信息     |
| /[appid]/wechat/token/get  |  api  |      √       |      ×       | 获取最新token      |
| /[appid]/wechat/ticket/get |  api  |      √       |      ×       | 获取最新ticket     |
| /wechat/token/refresh      | cron  |      √       |      √       | 定时刷新token任务  |
| /wechat/ticket/refresh     | cron  |      √       |      √       | 定时刷新ticket任务 |




##### 添加公众号

```sh
~/work/bin$ curl "http://192.168.5.71:9999/wechat/app/create?appid=wx9e02ddcc34513fd4&secret=6acb2b999177524beba3d97d54df2de5&token=oTSvVuXdjx1FPi6bz"

{"errcode":0，"errmsg":"success"}
```

##### 获取 access token

```sh
~/work/bin$  curl http://192.168.5.71:9999/wx9e02ddcc34513fd4/wechat/token/get

{"errcode":0，"errmsg":"success"，"access_token":"9_iZx3iItbJssKKzJTe9CeoJqY7678POnYkHRnZ_AfMfuV38CYxPHmOOvc7U0liXqon5vuZoIoU50RBLAZejTUSEwlXy5hl09KiyoWze65IXswjBnf6wFoUppoSRk4Z9opaOOJTSAxfR0DMGuaHCKgAHAAVD"，"expires_in":6599，"expires_date":"20180511113918"}
```

##### 获取 jsapi ticket

```sh
~/work/bin$  curl http://192.168.5.71:9999/wx9e02ddcc34513fd4/wechat/ticket/get

{"errcode":0，"errmsg":"success"，"ticket":"HoagFKDcsGMVCIY2vOjf9ki38m7AHQZG34U1VtA70B09ycvZdjvbjGm6qCTAF6_9fziC7iIQzWS49ZypavCK2g"，"expires_in":6599，"expires_date":"20180511121655"}
```




