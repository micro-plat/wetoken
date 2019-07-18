package sql

const QueryUserByName = `select t.* from wechat_user_info t where t.appid=@appid and t.user_name=@user_name and rownum<=1`
