// +build !oracle

package sql

import (
	"github.com/micro-plat/hydra"
)

func Install() {
	hydra.Installer.DB.AddSQL(`CREATE TABLE  uos_operate_log (
			log_id bigint  not null AUTO_INCREMENT comment 'id' ,
			user_id bigint    comment '用户id' ,
			user_name varchar(64)    comment '用户名' ,
			log_module varchar(64)  not null  comment '操作模块(角色权限、系统管理)' ,
			log_content varchar(1000)  not null  comment '操作内容' ,
			req_url varchar(64)  not null  comment '请求url' ,
			req_body varchar(1024)    comment '请求body' ,
			user_ip varchar(32)  not null  comment '用户ip' ,
			create_time datetime default current_timestamp not null  comment '创建时间' ,
			object_no1 varchar(32)    comment '自定义1' ,
			object_no2 varchar(32)    comment '自定义2' ,
			PRIMARY KEY (log_id),
			KEY key_uos_operate_log_user_id (user_id),
			KEY key_uos_operate_log_user_name (user_name),
			KEY key_uos_operate_log_log_module (log_module)
			) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='用户操作日志表';`)
}
