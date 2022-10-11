package ckafka

// 生产者配置
type PConf struct {
	Hosts []string `json:"hosts"`

	EnableSASL        bool   `json:"enable_sasl"`
	SASLPlainUsername string `json:"sasl_plain_username"`
	SASLPlainPassword string `json:"sasl_plain_password"`
}
