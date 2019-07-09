// +build oci

package sql

//----------------------access token---------------------------

const QueryAccessToken = `select t.appid,t.access_token,t.expire,to_char(t.expire_date,'yyyy/mm/dd hh24:mi:ss') expire_date from wechat_access_token t where t.appid=@appid`
const UpdateAccessToken = `update wechat_access_token t set 
t.access_token=@token,t.expire=@expire,t.expire_date=sysdate+@expire/24/60/60,t.update_time=sysdate,t.message=@message
 where t.appid=@appid`

const UpdateAccessTokenMessage = `update wechat_access_token t set 
 t.update_time=sysdate,t.message=@message
	where t.appid=@appid`

const InsertAccessToken = `insert into wechat_access_token(appid,access_token,expire,expire_date,update_time,message)
values
(@appid,@token,@expire,sysdate+@expire/24/60/60,sysdate,@message)`
