package config

type Feie struct {
	User           string `mapstructure:"user" json:"user" yaml:"user"`                                     // 账号
	UKey           string `mapstructure:"ukey" json:"ukey" yaml:"ukey"`                                     // 用户Key
	SN             string `mapstructure:"sn" json:"sn" yaml:"sn"`                                           // SN
	Url            string `mapstructure:"url" json:"url" yaml:"url"`                                        // 接口地址
	BusinessQrCode string `mapstructure:"business-qr-code" json:"business-qr-code" yaml:"business-qr-code"` // 商家二维码
	BusinessPhone  string `mapstructure:"business-phone" json:"business-phone" yaml:"business-phone"`       // 商家电话
}
