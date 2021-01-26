// +build oracle

package sql

import (
	"github.com/micro-plat/hydra"
)

func Install() {
	hydra.Installer.DB.AddSQL(`create table uos_operate_log(
			log_id NUMBER(19)  not null ,
			user_id NUMBER(19)   ,
			user_name VARCHAR2(64)   ,
			log_module VARCHAR2(64)  not null ,
			log_content VARCHAR2(1000)  not null ,
			req_url VARCHAR2(64)  not null ,
			req_body VARCHAR2(1024)   ,
			user_ip VARCHAR2(32)  not null ,
			create_time TIMESTAMP default current_timestamp not null ,
			object_no1 VARCHAR2(32)   ,
			object_no2 VARCHAR2(32)   
			);
			
		comment on table uos_operate_log is '用户操作日志表';
		comment on column uos_operate_log.log_id is 'id';	
		comment on column uos_operate_log.user_id is '用户id';	
		comment on column uos_operate_log.user_name is '用户名';	
		comment on column uos_operate_log.log_module is '操作模块(角色权限、系统管理)';	
		comment on column uos_operate_log.log_content is '操作内容';	
		comment on column uos_operate_log.req_url is '请求url';	
		comment on column uos_operate_log.req_body is '请求body';	
		comment on column uos_operate_log.user_ip is '用户ip';	
		comment on column uos_operate_log.create_time is '创建时间';	
		comment on column uos_operate_log.object_no1 is '自定义1';	
		comment on column uos_operate_log.object_no2 is '自定义2';	
	
		alter table uos_operate_log
		add constraint pk_operate_log primary key(log_id);
		create index key_uos_operate_log_user_id on uos_operate_log(user_id);
		create index key_uos_operate_log_user_name on uos_operate_log(user_name);
		create index key_uos_operate_log_log_module on uos_operate_log(log_module);
		
		create sequence seq_operate_log_id
		increment by 1
		minvalue 1
		maxvalue 99999999999
		start with 1
		cache 20;`)
}
