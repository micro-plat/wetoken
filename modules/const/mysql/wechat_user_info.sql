
	CREATE TABLE  wechat_user_info (
		user_id BIGINT(10)  not null AUTO_INCREMENT comment '用户编号' ,
		appid VARCHAR(32)  not null  comment 'appid' ,
		openid VARCHAR(32)  not null  comment 'openid' ,
		unionid VARCHAR(32)    comment 'unionid' ,
		user_name VARCHAR(32)  not null  comment '用户名' ,
		nick_name VARCHAR(32)    comment '昵称' ,
		head_url VARCHAR(128)    comment '头像地址' ,
		phone VARCHAR(16)    comment '联系电话' ,
		email VARCHAR(64)    comment '邮箱地址' ,
		status TINYINT(1)  not null  comment '状态（0：启用 1. 禁用 2. 锁定）' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		update_time DATETIME default current_timestamp not null  comment '更新时间' ,
		PRIMARY KEY (user_id),
		UNIQUE KEY unq_wechat_user_info_appid (appid),
		UNIQUE KEY unq_wechat_user_info_user_name (user_name)
		) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户基本信息';
