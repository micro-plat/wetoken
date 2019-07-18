// +build oci

package sql

//----------------------jsapi ticket---------------------------------

const QueryJSAPITicket = `select t.appid,t.ticket,t.expire,to_char(t.expire_date,'yyyy/mm/dd hh24:mi:ss') expire_date from wechat_jsapi_ticket t where t.appid=@appid`
const UpdateJSAPITicket = `update wechat_jsapi_ticket t set t.ticket=@ticket,expire=@expire,expire_date=sysdate+@expire/24/60/60 where t.appid=@appid`
const InsertJSAPITicket = `insert into wechat_jsapi_ticket(appid,ticket,expire,expire_date,update_time,message)
values
(@appid,@ticket,@expire,sysdate+@expire/24/60/60,sysdate,@message)`

const UpdateJSAPITicketMessage = `update wechat_jsapi_ticket t set 
 t.update_time=sysdate,t.message=@message
	where t.appid=@appid`
