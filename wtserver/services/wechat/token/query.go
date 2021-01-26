package token

import (
	"fmt"
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/wetoken/modules/wechat/token"
)

type QueryTokenHandler struct {
	appid string
	token token.IToken
}

//NewQueryTokenHandler 创建服务
func NewQueryTokenHandler() (u *QueryTokenHandler) {
	return &QueryTokenHandler{}
}
func (u *QueryTokenHandler) Handle(ctx hydra.IContext) (r interface{}) {
	var result struct {
		ErrCode int64  `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		token.AccessToken
	}
	appid := ctx.Request().Path().Params().GetString("appid")
	if appid == "" {
		return errs.NewError(http.StatusNotAcceptable, fmt.Errorf("参数appid错误,%s", appid))
	}
	tokenObj := token.NewToken(appid)
	tokenInfo, err := tokenObj.Query()
	if err != nil {
		result.ErrCode = 400
		result.ErrMsg = err.Error()
		return result
	}
	result.ErrCode = 0
	result.ErrMsg = "success"
	result.AccessToken = *tokenInfo
	result.AccessToken.Reset()
	return result
}
