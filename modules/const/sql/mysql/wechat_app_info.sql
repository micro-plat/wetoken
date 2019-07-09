
 drop table wechat_app_info;

	create table wechat_app_info(
		appid varchar(64)  not null PRIMARY KEY   comment 'appid' ,
		secret varchar(32)  not null    comment 'secret' ,
		token varchar(32)  not null    comment 'token' ,
		aes_key varchar(32)  not null    comment 'aes key' ,
		update_time datetime default current_timestamp not null    comment '最后更新时间' ,
		mchid varchar(64)  not null    comment '支付服务商编号' ,
		pay_key varchar(64)  not null    comment '支付服务商 key' ,
		sub_appid varchar(64)  not null    comment '子商户 app id' ,
		sub_mchid varchar(64)  not null    comment '子商户编户号' ,
		wechat_host varchar(64)      comment '微信授权域名' ,
		wcode_template_id varchar(64)      comment '微信验证码模板消息编号' 
				
  )COMMENT='微信公众号配置';

 




