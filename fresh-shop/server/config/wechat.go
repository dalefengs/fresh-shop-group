package config

type Wechat struct {
	Appid          string `mapstructure:"appid" json:"appid" yaml:"appid"`
	Secret         string `mapstructure:"secret" json:"secret" yaml:"secret"`
	RandomPassword string `mapstructure:"randomPassword" json:"randomPassword" yaml:"randomPassword"`
}
