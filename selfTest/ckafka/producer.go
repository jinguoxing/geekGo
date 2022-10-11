package ckafka

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strings"
	"log"
)

// 序列化类型: json
const SerializeJson = "json"

// 序列化类型: php
// 序列化的结果和 php 的 serialize 函数保持一致
const SerializePHP = "php"
const (
	INT32_MAX = 2147483647 - 1000
)
// 生产者
var producer *kafka.Producer
var producer01 *kafka.Producer

// 初始化生产者
func InitProducer(c PConf) {
	conf := kafka.ConfigMap{
		"bootstrap.servers": strings.Join(c.Hosts, ","),
		"api.version.request": "true",
		"message.max.bytes": 1000000,
		"linger.ms": 500,
		"sticky.partitioning.linger.ms" : 1000,
		"retries": INT32_MAX,
		"retry.backoff.ms": 1000,

	}

	// 和php端保持一致
	conf["enable.idempotence"] = "true"
	conf["compression.type"] = "lz4"
	conf["topic.request.required.acks"] = -1

	if true == c.EnableSASL {
		conf["sasl.mechanisms"] = "PLAIN"
		conf["security.protocol"] = "sasl_plaintext"
		conf["sasl.username"] = c.SASLPlainUsername
		conf["sasl.password"] = c.SASLPlainPassword
	}
	conf["api.version.request"] = "true"

	conf.SetKey("security.protocol", "sasl_ssl")
	conf.SetKey("ssl.ca.location", "../conf/ca-cert.pem")


	var err error
	producer, err = kafka.NewProducer(&conf)
	if nil != err {
		panic(fmt.Sprintf("kafka producer init error: %v", err))
	}
}

// 初始化生产者
func InitProducer01(c PConf) {
	conf := kafka.ConfigMap{
		"bootstrap.servers": strings.Join(c.Hosts, ","),
	}

	// 和php端保持一致
	conf["enable.idempotence"] = true
	conf["compression.type"] = "lz4"
	conf["topic.request.required.acks"] = -1

	if true == c.EnableSASL {
		conf["sasl.mechanisms"] = "PLAIN"
		conf["security.protocol"] = "sasl_plaintext"
		conf["sasl.username"] = c.SASLPlainUsername
		conf["sasl.password"] = c.SASLPlainPassword
	}

	var err error
	producer01, err = kafka.NewProducer(&conf)
	if nil != err {
		panic(fmt.Sprintf("kafka producer init error: %v", err))
	}
}

// 发送消息到kafka队列
func Send01(topic string, msg interface{}, serializeType ...string) error {
	if "" == topic {
		return fmt.Errorf("send to kafka err: topic empty")
	}

	wrapMsg := NewMsg1()
	if len(serializeType) > 0 {
		wrapMsg.SetSerializeType(serializeType[0])
	}
	err := wrapMsg.SetMsg(msg)
	if nil != err {
		return fmt.Errorf("send to kafka err: %v", err)
	}

	val, err := json.Marshal(wrapMsg)
	if nil != err {
		return fmt.Errorf("send to kafka err: %v", err)
	}

	kMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          val,
		Headers:        []kafka.Header{},
	}

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)
	err = producer01.Produce(kMsg, deliveryChan)
	if nil != err {
		return fmt.Errorf("send to kafka err: %v", err)
	}

	e := <-deliveryChan
	switch m := e.(type) {
	case *kafka.Message:
		if nil != m.TopicPartition.Error {
			return fmt.Errorf("delivery failed: %v", m.TopicPartition.Error)
		}
		return nil
	default:
		return fmt.Errorf("unknown event: %+v", m)
	}
}


func SendTest(topic string,value string) error {
	var msg *kafka.Message = nil

	msg = &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}

	//deliveryChan := make(chan kafka.Event, 1)
	//defer close(deliveryChan)
	err := producer.Produce(msg, nil)
	defer producer.Close()
	if nil != err {
		return fmt.Errorf("send to kafka err: %v", err)
	}
	producer.Flush(1 * 1000)
	return nil

	//e := <-deliveryChan
	//switch m := e.(type) {
	//case *kafka.Message:
	//	if nil != m.TopicPartition.Error {
	//		return fmt.Errorf("delivery failed: %v", m.TopicPartition.Error)
	//	}
	//	return nil
	//default:
	//	return fmt.Errorf("unknown event: %+v", m)
	//}

}



func SendTest01(topic string,value string) error {



	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to write access log entry:%v", ev.TopicPartition.Error)
				} else {
					log.Printf("Send OK topic:%v partition:%v offset:%v content:%s\n", *ev.TopicPartition.Topic,  ev.TopicPartition.Partition, ev.TopicPartition.Offset, ev.Value)

				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)

		var msg *kafka.Message = nil

			msg = &kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(value),
			}

		producer.Produce(msg, nil)

	// Wait for message deliveries before shutting down
	producer.Flush(15 * 1000)

	return nil
}




// 发送消息到kafka队列
func Send(topic string, msg interface{}, serializeType ...string) error {
	if "" == topic {
		return fmt.Errorf("send to kafka err: topic empty")
	}

	wrapMsg := NewMsg1()
	if len(serializeType) > 0 {
		wrapMsg.SetSerializeType(serializeType[0])
	}
	err := wrapMsg.SetMsg(msg)
	if nil != err {
		return fmt.Errorf("send to kafka err: %v", err)
	}

	val, err := json.Marshal(wrapMsg)
	if nil != err {
		return fmt.Errorf("send to kafka err: %v", err)
	}

	kMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          val,
		Headers:        []kafka.Header{},
	}

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)
	err = producer.Produce(kMsg, deliveryChan)
	if nil != err {
		return fmt.Errorf("send to kafka err: %v", err)
	}

	e := <-deliveryChan
	switch m := e.(type) {
	case *kafka.Message:
		if nil != m.TopicPartition.Error {
			return fmt.Errorf("delivery failed: %v", m.TopicPartition.Error)
		}
		return nil
	default:
		return fmt.Errorf("unknown event: %+v", m)
	}
}

// 关闭kafka生产者
func CloseProducer() {
	if nil == producer {
		return
	}

	producer.Close()
}
