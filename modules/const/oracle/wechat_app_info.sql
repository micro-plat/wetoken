
	create table wechat_app_info(
		appid varchar2(64)  not null ,
		secret varchar2(32)  not null ,
		token varchar2(32)  not null ,
		aes_key varchar2(32)  not null ,
		update_time date default sysdate not null ,
		mchid varchar2(64)  not null ,
		pay_key varchar2(64)  not null ,
		sub_appid varchar2(64)  not null ,
		sub_mchid varchar2(64)  not null ,
		wechat_host varchar2(64)   ,
		wcode_template_id varchar2(64)   
		);
	

	comment on table wechat_app_info is '微信公众号配置';
	comment on column wechat_app_info.appid is 'appid';	
	comment on column wechat_app_info.secret is 'secret';	
	comment on column wechat_app_info.token is 'token';	
	comment on column wechat_app_info.aes_key is 'aes key';	
	comment on column wechat_app_info.update_time is '最后更新时间';	
	comment on column wechat_app_info.mchid is '支付服务商编号';	
	comment on column wechat_app_info.pay_key is '支付服务商 key';	
	comment on column wechat_app_info.sub_appid is '子商户 app id';	
	comment on column wechat_app_info.sub_mchid is '子商户编户号';	
	comment on column wechat_app_info.wechat_host is '微信授权域名';	
	comment on column wechat_app_info.wcode_template_id is '微信验证码模板消息编号';	
	

 
	alter table wechat_app_info
	add constraint pk_app_info primary key(appid);
	