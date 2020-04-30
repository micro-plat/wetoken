package wetoken

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/wetoken/modules/wechat/ticket"
)

//TicketResult .
type TicketResult struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	ticket.JSTicket
}

//GetWxTicket 获取Ticket
func GetWxTicket(container component.IContainer, appID string) (results *TicketResult, err error) {
	ticketLib := ticket.NewTicket(container, appID)
	ticket, err := ticketLib.Get()
	result := &TicketResult{
		ErrCode: 0,
		ErrMsg:  "success",
	}
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result, nil
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	result.JSTicket = *ticket
	result.JSTicket.Reset()
	return result, nil
}
