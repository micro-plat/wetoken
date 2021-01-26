
	CREATE TABLE  wechat_jsapi_ticket (
		appid VARCHAR(64)  not null  comment 'appid' ,
		ticket VARCHAR(64)  not null  comment 'js api ticket' ,
		expire BIGINT(20)  not null  comment '过期时长' ,
		expire_date DATETIME  not null  comment '过期日期' ,
		update_time DATETIME default current_timestamp not null  comment '更新日期' ,
		message VARCHAR(64)  not null  comment '刷新消息' ,
		PRIMARY KEY (appid)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='网页票据信息';
