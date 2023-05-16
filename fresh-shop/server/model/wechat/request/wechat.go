package request

type Jscode2SessionReq struct {
	Appid  string `json:"appid" form:"appid"`
	Secret string `json:"secret" form:"secret"`
	Jscode string `json:"js_code" form:"js_code"`
}

type LoginReq struct {
	SessionKey    string `json:"sessionKey" form:"sessionKey"`
	EncryptedData string `json:"encryptedData" form:"encryptedData"`
	Iv            string `json:"iv" form:"iv"`
	OpenId        string `json:"openid" form:"openid"`
}

type WechatPayReq struct {
	OpenId   string `json:"openId" form:"openId"`
	Body     string `json:"body" form:"body"`
	ClientIP string `json:"clientIP" form:"clientIP"` // 客户端IP
}
