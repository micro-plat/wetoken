

	CREATE TABLE  wechat_app_info (
		appid VARCHAR(64)  not null  comment 'appid' ,
		secret VARCHAR(32)  not null  comment 'secret' ,
		token VARCHAR(32)  not null  comment 'token' ,
		aes_key VARCHAR(32)  not null  comment 'aes key' ,
		update_time DATETIME default current_timestamp not null  comment '最后更新时间' ,
		mchid VARCHAR(64)  not null  comment '支付服务商编号' ,
		pay_key VARCHAR(64)  not null  comment '支付服务商 key' ,
		sub_appid VARCHAR(64)  not null  comment '子商户 app id' ,
		sub_mchid VARCHAR(64)  not null  comment '子商户编户号' ,
		wechat_host VARCHAR(64)    comment '微信授权域名' ,
		wcode_template_id VARCHAR(64)    comment '微信验证码模板消息编号' ,
		PRIMARY KEY (appid)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='微信公众号配置';


	CREATE TABLE  wechat_jsapi_ticket (
		appid VARCHAR(64)  not null  comment 'appid' ,
		ticket VARCHAR(64)  not null  comment 'js api ticket' ,
		expire BIGINT(20)  not null  comment '过期时长' ,
		expire_date DATETIME  not null  comment '过期日期' ,
		update_time DATETIME default current_timestamp not null  comment '更新日期' ,
		message VARCHAR(64)  not null  comment '刷新消息' ,
		PRIMARY KEY (appid)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='网页票据信息';


	CREATE TABLE  wechat_access_token (
		appid VARCHAR(64)  not null  comment 'appid' ,
		access_token VARCHAR(64)  not null  comment 'access token' ,
		expire BIGINT(20)  not null  comment '过期时长' ,
		expire_date DATETIME  not null  comment '过期日期' ,
		update_time DATETIME default current_timestamp not null  comment '更新日期' ,
		message VARCHAR(64)  not null  comment '刷新消息' ,
		PRIMARY KEY (appid)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='后端通信 token';


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
