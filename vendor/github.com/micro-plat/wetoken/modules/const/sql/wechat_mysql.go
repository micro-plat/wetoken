// +build !oci

package sql

//--------------------------wechat app------------------
const InsertWebchatApp = `insert into wechat_app_info
(appid,secret,token,aes_key,mchid,pay_key,sub_appid,sub_mchid,update_time)
values
(@appid,@secret,@token,@aes_key,@mchid,@pay_key,@sub_appid,@sub_mchid,now())`

const QueryWechatApp = `select t.* from wechat_app_info t where t.appid=@appid`
