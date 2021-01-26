package message

import (
	"fmt"
	"time"

	"github.com/lib4dev/wechat/mp"

	"github.com/lib4dev/wechat/mp/message/template"
)

//wxSend 发送微信验证码
func wxSend(appid string, openid string, sysName string, tpid string, code string, ctx *mp.Context) error {
	if _, err := template.Send(ctx, &template.TemplateMessage2{
		ToUser:     openid,
		TemplateId: tpid,
		Data: map[string]interface{}{
			"first": map[string]string{
				"value": fmt.Sprintf("您正在登录[%s]", sysName),
			},
			"keyword1": map[string]string{
				"value": code,
				"color": "#43CD80",
			},
			"keyword2": map[string]string{
				"value": "5分钟",
			},
			"keyword3": map[string]string{
				"value": time.Now().Format("2006/01/02 15:04:05"),
			},
			"remark": map[string]string{
				"value": "若非本人操作请注意账户安全",
			},
		},
	}); err != nil {
		return fmt.Errorf("发送验证码失败:%v(%s)%s", err, openid, code)
	}
	return nil
}
