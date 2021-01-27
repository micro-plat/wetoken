
	DROP TABLE wechat_jsapi_ticket;
	create table wechat_jsapi_ticket(
		appid varchar2(64)  not null ,
		ticket varchar2(64)  not null ,
		expire number(20)  not null ,
		expire_date date  not null ,
		update_time date default sysdate not null ,
		message varchar2(64)  not null 
		);
	

	comment on table wechat_jsapi_ticket is '网页票据信息';
	comment on column wechat_jsapi_ticket.appid is 'appid';	
	comment on column wechat_jsapi_ticket.ticket is 'js api ticket';	
	comment on column wechat_jsapi_ticket.expire is '过期时长';	
	comment on column wechat_jsapi_ticket.expire_date is '过期日期';	
	comment on column wechat_jsapi_ticket.update_time is '更新日期';	
	comment on column wechat_jsapi_ticket.message is '刷新消息';	
	

 
	alter table wechat_jsapi_ticket
	add constraint pk_jsapi_ticket primary key(appid);
	