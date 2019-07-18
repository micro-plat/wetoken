
 drop table wechat_user_info;

	create table wechat_user_info(
		user_id int  not null PRIMARY KEY AUTO_INCREMENT  comment '用户编号' ,
		appid varchar(32)  not null    comment 'appid' ,
		openid varchar(32)  not null    comment 'openid' ,
		unionid varchar(32)      comment 'unionid' ,
		user_name varchar(32)  not null    comment '用户名' ,
		nick_name varchar(32)      comment '昵称' ,
		head_url varchar(128)      comment '头像地址' ,
		phone varchar(16)      comment '联系电话' ,
		email varchar(64)      comment '邮箱地址' ,
		status tinyint(1)  not null    comment '状态（0：启用 1.禁用 2.锁定）' ,
		create_time datetime default current_timestamp not null    comment '创建时间' ,
		update_time datetime default current_timestamp not null    comment '更新时间' 
				
  )COMMENT='用户基本信息';

 




	drop index wechat_user_info_appid ON wechat_user_info;
 create unique index wechat_user_info_appid ON wechat_user_info(appid,user_name);
 
