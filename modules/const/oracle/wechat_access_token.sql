
	create table wechat_access_token(
		appid varchar2(64)  not null ,
		access_token varchar2(64)  not null ,
		expire number(20)  not null ,
		expire_date date  not null ,
		update_time date default sysdate not null ,
		message varchar2(64)  not null 
		);
	

	comment on table wechat_access_token is '后端通信 token';
	comment on column wechat_access_token.appid is 'appid';	
	comment on column wechat_access_token.access_token is 'access token';	
	comment on column wechat_access_token.expire is '过期时长';	
	comment on column wechat_access_token.expire_date is '过期日期';	
	comment on column wechat_access_token.update_time is '更新日期';	
	comment on column wechat_access_token.message is '刷新消息';	
	

 
	alter table wechat_access_token
	add constraint pk_access_token primary key(appid);
	