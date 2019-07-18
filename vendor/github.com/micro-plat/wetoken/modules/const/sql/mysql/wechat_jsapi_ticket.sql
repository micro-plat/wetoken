
 drop table wechat_jsapi_ticket;

	create table wechat_jsapi_ticket(
		appid varchar(64)  not null PRIMARY KEY   comment 'appid' ,
		ticket varchar(64)  not null    comment 'js api ticket' ,
		expire bigint  not null    comment '过期时长' ,
		expire_date datetime  not null    comment '过期日期' ,
		update_time datetime default current_timestamp not null    comment '更新日期' ,
		message varchar(64)  not null    comment '刷新消息' 
				
  )COMMENT='网页票据信息';

 




