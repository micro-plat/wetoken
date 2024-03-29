package ticket

import (
	"fmt"
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/wetoken/modules/const/sql"
)

//DBTicket 数据库jsapi ticket
type DBTicket struct {
	appid string
}

//NewDBTicket 构建数据库jsapi ticket访问对象
func NewDBTicket(appid string) *DBTicket {
	return &DBTicket{
		appid: appid,
	}
}

//Update 修改jsapi ticket查询消息
func (t *DBTicket) Update(msg string) error {
	db := hydra.C.DB().GetRegularDB()
	_, err := db.Execute(sql.UpdateJSAPITicketMessage, map[string]interface{}{
		"appid":   t.appid,
		"message": msg,
	})
	if err != nil {
		return err
	}
	return nil
}

//Save 保存jsapi ticket到db,不存在时创建
func (t *DBTicket) Save(tk *JSTicket) error {
	db := hydra.C.DB().GetRegularDB()
	//更新jsapi ticket
	r, err := db.Execute(sql.UpdateJSAPITicket, map[string]interface{}{
		"appid":   t.appid,
		"ticket":  tk.Ticket,
		"expire":  tk.ExpiresIn,
		"message": "success",
	})
	if err != nil {
		return err
	}
	if r == 1 {
		return nil
	}
	//创建jsapi ticket记录
	_, err = db.Execute(sql.InsertJSAPITicket, map[string]interface{}{
		"appid":   t.appid,
		"ticket":  tk.Ticket,
		"expire":  tk.ExpiresIn,
		"message": "success",
	})
	if err != nil {
		return err
	}
	return nil
}

//Get 从数据库获取jsapi ticket
func (t *DBTicket) Get() (token *JSTicket, err error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(sql.QueryJSAPITicket, map[string]interface{}{"appid": t.appid})
	if err != nil {
		return nil, fmt.Errorf("从数据库中获取ticket失败:%s %v", t.appid, err)
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("从数据库中获取ticket失败:%s app信息未配置", t.appid)
	}

	//处理过期时间
	expire, err := data.Get(0).GetDatetime("expire_date")
	if err != nil {
		return nil, err
	}
	//检查是否过期
	if time.Since(expire) >= 0 {
		return nil, errTicketExpired
	}
	token = &JSTicket{
		Ticket:      data.Get(0).GetString("ticket"),
		ExpiresIn:   data.Get(0).GetInt64("expire"),
		ExpiresDate: jtime(expire),
	}
	//更新expire
	token.Reset()
	return token, nil
}

//GetList 从数据库获取jsapi ticket
func GetList() (list []string, err error) {
	db := hydra.C.DB().GetRegularDB()
	datas, err := db.Query(sql.QueryJSAPITicketList, nil)
	if err != nil {
		return nil, fmt.Errorf("从数据库中获取ticket appid list失败: %v", err)
	}

	list = make([]string, 0)
	if datas.IsEmpty() {
		return
	}
	for _, item := range datas {
		appid := item.GetString("appid")
		if appid != "" {
			list = append(list, item.GetString("appid"))
		}
	}
	return
}
