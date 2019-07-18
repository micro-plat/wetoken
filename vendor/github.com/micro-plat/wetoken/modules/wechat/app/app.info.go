package app

//WechatAppInfo 微信公众号配置
type WechatAppInfo struct {

	//AppID Appid
	AppID string `json:"appID" form:"appID" m2s:"appID" valid:"required"`

	//Secret Secret
	Secret string `json:"secret" form:"secret" m2s:"secret" valid:"required"`

	//Token Token
	Token string `json:"token" form:"token" m2s:"token" valid:"required"`

	//AesKey Aes key
	AesKey string `json:"aes_key" form:"aes_key" m2s:"aes_key" valid:"required"`

	//MchID 支付服务商编号
	MchID string `json:"mchID" form:"mchID" m2s:"mchID" valid:"required"`

	//PayKey 支付服务商 key
	PayKey string `json:"pay_key" form:"pay_key" m2s:"pay_key" valid:"required"`

	//SubAppID 子商户 app id
	SubAppID string `json:"sub_appID" form:"sub_appID" m2s:"sub_appID" valid:"required"`

	//SubMchID 子商户编户号
	SubMchID string `json:"sub_mchID" form:"sub_mchID" m2s:"sub_mchID" valid:"required"`

	//WechatHost 微信授权域名
	WechatHost string `json:"wechat_host" form:"wechat_host" m2s:"wechat_host" `

	//WcodeTemplateID 微信验证码模板消息编号
	WcodeTemplateID string `json:"wcode_template_ID" form:"wcode_template_ID" m2s:"wcode_template_ID" `
}
