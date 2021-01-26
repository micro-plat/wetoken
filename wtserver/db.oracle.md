# 微信中控服务器数据字典

### 1. 微信公众号配置[wechat_app_info]

| 字段名            | 类型         | 默认值  | 为空 | 约束  | 描述                   |
| ----------------- | ------------ | :-----: | :--: | :---: | :--------------------- |
| appid             | varchar2(64) |         |  否  | PK, IS | appid                  |
| secret            | varchar2(32) |         |  否  |  IUS  | secret                 |
| token             | varchar2(32) |         |  否  |  IUS  | token                  |
| aes_key           | varchar2(32) |         |  否  |  IUS  | aes key                |
| update_time       | date         | sysdate |  否  |  IS   | 最后更新时间           |
| mchid             | varchar2(64) |         |  否  |  IUS  | 支付服务商编号         |
| pay_key           | varchar2(64) |         |  否  |  IUS  | 支付服务商 key         |
| sub_appid         | varchar2(64) |         |  否  |  IUS  | 子商户 app id          |
| sub_mchid         | varchar2(64) |         |  否  |  IUS  | 子商户编户号           |
| wechat_host       | varchar2(64) |         |  是  |  IUS  | 微信授权域名           |
| wcode_template_id | varchar2(64) |         |  是  |  IUS  | 微信验证码模板消息编号 |

### 2. 网页票据信息[wechat_jsapi_ticket]

| 字段名      | 类型         | 默认值  | 为空 | 约束  | 描述          |
| ----------- | ------------ | :-----: | :--: | :---: | :------------ |
| appid       | varchar2(64) |         |  否  | PK, IS | appid         |
| ticket      | varchar2(64) |         |  否  |  IS   | js api ticket |
| expire      | number(20)   |         |  否  |  IS   | 过期时长      |
| expire_date | date         |         |  否  |  IS   | 过期日期      |
| update_time | date         | sysdate |  否  |  IS   | 更新日期      |
| message     | varchar2(64) |         |  否  |  IS   | 刷新消息      |

### 3. 后端通信 token[wechat_access_token]

| 字段名       | 类型         | 默认值  | 为空 | 约束  | 描述         |
| ------------ | ------------ | :-----: | :--: | :---: | :----------- |
| appid        | varchar2(64) |         |  否  | PK, IS | appid        |
| access_token | varchar2(64) |         |  否  |  IS   | access token |
| expire       | number(20)   |         |  否  |  IS   | 过期时长     |
| expire_date  | date         |         |  否  |  IS   | 过期日期     |
| update_time  | date         | sysdate |  否  |  IS   | 更新日期     |
| message      | varchar2(64) |         |  否  |  IS   | 刷新消息     |

### 4. 用户基本信息[wechat_user_info]

| 字段名      | 类型          | 默认值  | 为空 |   约束    | 描述                          |
| ----------- | ------------- | :-----: | :--: | :-------: | :---------------------------- |
| user_id     | number(10)    |         |  否  | PK, IS, SEQ | 用户编号                      |
| appid       | varchar2(32)  |         |  否  |  IS, UNQ   | appid                         |
| openid      | varchar2(32)  |         |  否  |    IS     | openid                        |
| unionid     | varchar2(32)  |         |  是  |    IS, | unionid                       |
| user_name   | varchar2(32)  |         |  否  |  IS, UNQ   | 用户名                        |
| nick_name   | varchar2(32)  |         |  是  |    IS     | 昵称                          |
| head_url    | varchar2(128) |         |  是  |    IS     | 头像地址                      |
| phone       | varchar2(16)  |         |  是  |    IS     | 联系电话                      |
| email       | varchar2(64)  |         |  是  |    IS     | 邮箱地址                      |
| status      | number(1)     |         |  否  |    IS     | 状态（0：启用 1. 禁用 2. 锁定） |
| create_time | date          | sysdate |  否  |    IS     | 创建时间                      |
| update_time | date          | sysdate |  否  |    IS     | 更新时间                      |
