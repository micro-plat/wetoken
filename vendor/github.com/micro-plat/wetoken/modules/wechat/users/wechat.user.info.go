package users

//WechatUserInfo 用户基本信息
type WechatUserInfo struct {

	//UserID 用户编号
	UserID int `json:"user_ID" form:"user_ID" m2s:"user_ID" valid:"required"`

	//AppID Appid
	AppID string `json:"appID" form:"appID" m2s:"appID" valid:"required"`

	//OpenID Openid
	OpenID string `json:"openID" form:"openID" m2s:"openID" valid:"required"`

	//UnionID Unionid
	UnionID string `json:"unionID" form:"unionID" m2s:"unionID" `

	//UserName 用户名
	UserName string `json:"user_name" form:"user_name" m2s:"user_name" valid:"required"`

	//NickName 昵称
	NickName string `json:"nick_name" form:"nick_name" m2s:"nick_name" `

	//HeadURL 头像地址
	HeadURL string `json:"head_url" form:"head_url" m2s:"head_url" `

	//Phone 联系电话
	Phone string `json:"phone" form:"phone" m2s:"phone" `

	//Email 邮箱地址
	Email string `json:"email" form:"email" m2s:"email" `

	//Status 状态（0：启用 1.禁用 2.锁定）
	Status int `json:"status" form:"status" m2s:"status" valid:"required"`
}
