package request

type Jscode2SessionReq struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
	Jscode string `json:"js_code"`
}
