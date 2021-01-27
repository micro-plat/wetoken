
	DROP TABLE IF EXISTS  wechat_access_token;
	CREATE TABLE  wechat_access_token (
		appid VARCHAR(64)  not null  comment 'appid' ,
		access_token VARCHAR(64)  not null  comment 'access token' ,
		expire BIGINT(20)  not null  comment '过期时长' ,
		expire_date DATETIME  not null  comment '过期日期' ,
		update_time DATETIME default current_timestamp not null  comment '更新日期' ,
		message VARCHAR(64)  not null  comment '刷新消息' ,
		PRIMARY KEY (appid)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='后端通信 token';
