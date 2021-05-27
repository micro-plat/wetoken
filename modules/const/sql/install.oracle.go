// +build oracle

package sql

import (
	"sync"

	_ "github.com/mattn/go-oci8"
	"github.com/micro-plat/hydra"
)

var onceLock sync.Once

func Install() {
	onceLock.Do(func() {
		installOnce()
	})
}

//seq_user_info_id
func installOnce() {
	hydra.Installer.DB.AddSQL(`DROP TABLE wechat_app_info;
	create table wechat_app_info(
		appid varchar2(64)  not null ,
		secret varchar2(32)  not null ,
		token varchar2(32)  not null ,
		aes_key varchar2(32)  not null ,
		update_time date default sysdate not null ,
		mchid varchar2(64)  not null ,
		pay_key varchar2(64)  not null ,
		sub_appid varchar2(64)  not null ,
		sub_mchid varchar2(64)  not null ,
		wechat_host varchar2(64)   ,
		wcode_template_id varchar2(64)   
		);
	

	comment on table wechat_app_info is '微信公众号配置';
	comment on column wechat_app_info.appid is 'appid';	
	comment on column wechat_app_info.secret is 'secret';	
	comment on column wechat_app_info.token is 'token';	
	comment on column wechat_app_info.aes_key is 'aes key';	
	comment on column wechat_app_info.update_time is '最后更新时间';	
	comment on column wechat_app_info.mchid is '支付服务商编号';	
	comment on column wechat_app_info.pay_key is '支付服务商 key';	
	comment on column wechat_app_info.sub_appid is '子商户 app id';	
	comment on column wechat_app_info.sub_mchid is '子商户编户号';	
	comment on column wechat_app_info.wechat_host is '微信授权域名';	
	comment on column wechat_app_info.wcode_template_id is '微信验证码模板消息编号';	
	

 
	alter table wechat_app_info
	add constraint pk_app_info primary key(appid);
	
	DROP TABLE wechat_jsapi_ticket;
	create table wechat_jsapi_ticket(
		appid varchar2(64)  not null ,
		ticket varchar2(64)  not null ,
		expire number(20)  not null ,
		expire_date date  not null ,
		update_time date default sysdate not null ,
		message varchar2(64)  not null 
		);
	

	comment on table wechat_jsapi_ticket is '网页票据信息';
	comment on column wechat_jsapi_ticket.appid is 'appid';	
	comment on column wechat_jsapi_ticket.ticket is 'js api ticket';	
	comment on column wechat_jsapi_ticket.expire is '过期时长';	
	comment on column wechat_jsapi_ticket.expire_date is '过期日期';	
	comment on column wechat_jsapi_ticket.update_time is '更新日期';	
	comment on column wechat_jsapi_ticket.message is '刷新消息';	
	

 
	alter table wechat_jsapi_ticket
	add constraint pk_jsapi_ticket primary key(appid);
	
	DROP TABLE wechat_access_token;
	create table wechat_access_token(
		appid varchar2(64)  not null ,
		access_token varchar2(64)  not null ,
		expire number(20)  not null ,
		expire_date date  not null ,
		update_time date default sysdate not null ,
		message varchar2(64)  not null 
		);
	

	comment on table wechat_access_token is '后端通信 token';
	comment on column wechat_access_token.appid is 'appid';	
	comment on column wechat_access_token.access_token is 'access token';	
	comment on column wechat_access_token.expire is '过期时长';	
	comment on column wechat_access_token.expire_date is '过期日期';	
	comment on column wechat_access_token.update_time is '更新日期';	
	comment on column wechat_access_token.message is '刷新消息';	
	

 
	alter table wechat_access_token
	add constraint pk_access_token primary key(appid);
	
	DROP TABLE wechat_user_info;
	create table wechat_user_info(
		user_id number(10)  not null ,
		appid varchar2(32)  not null ,
		openid varchar2(32)  not null ,
		unionid varchar2(32)   ,
		user_name varchar2(32)  not null ,
		nick_name varchar2(32)   ,
		head_url varchar2(128)   ,
		phone varchar2(16)   ,
		email varchar2(64)   ,
		status number(1)  not null ,
		create_time date default sysdate not null ,
		update_time date default sysdate not null 
		);
	

	comment on table wechat_user_info is '用户基本信息';
	comment on column wechat_user_info.user_id is '用户编号';	
	comment on column wechat_user_info.appid is 'appid';	
	comment on column wechat_user_info.openid is 'openid';	
	comment on column wechat_user_info.unionid is 'unionid';	
	comment on column wechat_user_info.user_name is '用户名';	
	comment on column wechat_user_info.nick_name is '昵称';	
	comment on column wechat_user_info.head_url is '头像地址';	
	comment on column wechat_user_info.phone is '联系电话';	
	comment on column wechat_user_info.email is '邮箱地址';	
	comment on column wechat_user_info.status is '状态（0：启用 1. 禁用 2. 锁定）';	
	comment on column wechat_user_info.create_time is '创建时间';	
	comment on column wechat_user_info.update_time is '更新时间';	
	

 
	alter table wechat_user_info
	add constraint pk_user_info primary key(user_id);
	alter table wechat_user_info
	add constraint unq_wechat_user_info_appid unique(appid);
	alter table wechat_user_info
	add constraint unq_wechat_user_info_user_name unique(user_name);
	
	create sequence seq_user_info
	increment by 1
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;`)
}
