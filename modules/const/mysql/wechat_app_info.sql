
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
