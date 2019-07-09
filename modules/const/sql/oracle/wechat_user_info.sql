
	create table wechat_user_info(
		user_id  number(10)      not null ,
		appid  varchar2(32)    not null ,
		openid  varchar2(32)    not null ,
		unionid  varchar2(32)     ,
		user_name  varchar2(32)    not null ,
		nick_name  varchar2(32)     ,
		head_url  varchar2(128)    ,
		phone  varchar2(16)     ,
		email  varchar2(64)     ,
		status  number(1)       not null ,
		create_time  date           default sysdate not null ,
		update_time  date           default sysdate not null 
				
  );

	comment on table wechat_user_info is '用户基本信息';
	comment on column wechat_user_info.user_id is '用户编号';	
	comment on column wechat_user_info.appid is 'appid';	
	comment on column wechat_user_info.openid is 'openid';	
	comment on column wechat_user_info.unionid is 'unionid';	
	comment on column wechat_user_info.user_name is '用户名';	
	comment on column wechat_user_info.nick_name is '昵称';	
	comment on column wechat_user_info.head_url is '头像地址';	
	comment on column wechat_user_info.phone is '联系电话';	
	comment on column wechat_user_info.email is '邮箱地址';	
	comment on column wechat_user_info.status is '状态（0：启用 1.禁用 2.锁定）';	
	comment on column wechat_user_info.create_time is '创建时间';	
	comment on column wechat_user_info.update_time is '更新时间';	
	
 
	alter table wechat_user_info
	add constraint pk_user_info primary key(user_id);

	alter table wechat_user_info
  add constraint wechat_user_info_appid unique(appid,user_name);

	
	create sequence seq_user_info_id
	minvalue 100
	maxvalue 99999999999
	start with 100
	cache 20;
	
