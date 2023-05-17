package config

type Wechat struct {
	Appid          string `mapstructure:"appid" json:"appid" yaml:"appid"`
	Secret         string `mapstructure:"secret" json:"secret" yaml:"secret"`
	RandomPassword string `mapstructure:"randomPassword" json:"randomPassword" yaml:"randomPassword"`
}

type WechatPay struct {
	MchId     string `mapstructure:"mchId" json:"mchId" yaml:"mchId"`             // 商户号
	ApiV2Key  string `mapstructure:"apiV2Key" json:"apiV2Key" yaml:"apiV2Key"`    // 商户号
	NotifyURL string `mapstructure:"notifyUrl" json:"notifyUrl" yaml:"notifyUrl"` // 微信支付通知地址
	Debug     bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	CertPath  string `mapstructure:"certPath" json:"certPath" yaml:"certPath"`
	KeyPath   string `mapstructure:"keyPath" json:"keyPath" yaml:"keyPath"`
}
