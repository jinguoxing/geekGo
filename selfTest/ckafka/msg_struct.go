package ckafka

type SkuMessageForKafka struct {
	PackageName string            `json:"package_name"`
	SkuList     []*SkuListMessage `json:"sku_list"`
	EsIndex     string            `json:"es_index"`
	Total       int               `json:"total"`
	MessageType MessageType       `json:"message_type"`
}

type SkuListMessage struct {
	Sku        string `json:"sku"`
	UpdateTime int64  `json:"update_time"`
}

type MessageType string

const (
	StatusMsg MessageType = "StatusMsg"
	RoundMsg  MessageType = "RoundMsg"
)
