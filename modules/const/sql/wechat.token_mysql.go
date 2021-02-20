// +build !oracle

package sql

//----------------------access token---------------------------

const QueryAccessToken = `select t.appid,t.access_token,t.expire,DATE_FORMAT(t.expire_date,'%Y/%m/%d %H:%i:%S') expire_date from wechat_access_token t where t.appid=@appid`
const QueryAccessTokenList = `select t.appid from wechat_access_token t`
const UpdateAccessToken = `update wechat_access_token t set 
t.access_token=@token,t.expire=@expire,t.expire_date=DATE_ADD(now(),interval #expire second),t.update_time=now(),t.message=@message
 where t.appid=@appid`

const UpdateAccessTokenMessage = `update wechat_access_token t set 
 t.update_time=now(),t.message=@message
	where t.appid=@appid`

const InsertAccessToken = `insert into wechat_access_token(appid,access_token,expire,expire_date,update_time,message)
values
(@appid,@token,@expire,DATE_ADD(now(),interval #expire second),now(),@message)`
