// +build !oci

package sql

//----------------------jsapi ticket---------------------------------

const QueryJSAPITicket = `select t.appid,t.ticket,t.expire,DATE_FORMAT(t.expire_date,'%Y/%m/%d %H:%i:%S') expire_date from wechat_jsapi_ticket t where t.appid=@appid`
const UpdateJSAPITicket = `update wechat_jsapi_ticket t set t.ticket=@ticket,expire=@expire,expire_date=DATE_ADD(now(),interval #expire second) where t.appid=@appid`
const InsertJSAPITicket = `insert into wechat_jsapi_ticket(appid,ticket,expire,expire_date,update_time,message)
values
(@appid,@ticket,@expire,DATE_ADD(now(),interval #expire second),now(),@message)`

const UpdateJSAPITicketMessage = `update wechat_jsapi_ticket t set 
 t.update_time=now(),t.message=@message
	where t.appid=@appid`
