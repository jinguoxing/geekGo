package ckafka

// 消费者配置
type CConf struct {
	Hosts []string `json:"hosts"`

	GroupId string

	EnableSASL        bool   `json:"enable_sasl"`
	SASLPlainUsername string `json:"sasl_plain_username"`
	SASLPlainPassword string `json:"sasl_plain_password"`
}
