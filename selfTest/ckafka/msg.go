package ckafka

import (
	"encoding/json"
	"fmt"
	"lib.go.ickey.cn/serialize"
	"sync"
)

type Msg struct {
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`

	serializeType string

	mu sync.Mutex
}

func NewMsg(msg interface{}) (*Msg, error) {
	m := NewMsg1()
	err := m.SetMsg(msg)
	if nil != err {
		return nil, err
	}

	return m, nil
}

func NewMsg1() *Msg {
	return &Msg{
		Headers:       make(map[string]string),
		serializeType: SerializePHP,
	}
}

// 设置消息序列化类型
func (m *Msg) SetSerializeType(serializeType string) *Msg {
	if SerializeJson == serializeType {
		m.serializeType = SerializeJson
	}

	return m
}

func (m *Msg) SetMsg(msg interface{}) error {
	serializeFunc := serialize.Marshal
	if SerializeJson == m.serializeType {
		serializeFunc = json.Marshal
	}

	msgStr, err := serializeFunc(msg)
	if nil != err {
		return fmt.Errorf("msg 序列化错误: %v", err)
	}

	m.Body = string(msgStr)

	return nil
}

func (m *Msg) SetHeader(key, val string) *Msg {
	m.mu.Lock()
	m.Headers[key] = val
	m.mu.Unlock()
	return m
}

func (m *Msg) SetHeaders(headers map[string]string) *Msg {
	m.mu.Lock()
	for k, v := range headers {
		m.Headers[k] = v
	}
	m.mu.Unlock()
	return m
}
