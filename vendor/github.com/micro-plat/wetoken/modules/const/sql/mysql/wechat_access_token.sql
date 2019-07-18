
 drop table wechat_access_token;

	create table wechat_access_token(
		appid varchar(64)  not null PRIMARY KEY   comment 'appid' ,
		access_token varchar(64)  not null    comment 'access token' ,
		expire bigint  not null    comment '过期时长' ,
		expire_date datetime  not null    comment '过期日期' ,
		update_time datetime default current_timestamp not null    comment '更新日期' ,
		message varchar(64)  not null    comment '刷新消息' 
				
  )COMMENT='后端通信 token';

 




