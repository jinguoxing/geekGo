package ckafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tidwall/gjson"
	"lib.go.ickey.cn/common"
	"lib.go.ickey.cn/log"
	"strings"
)

// 错误: 读取队列阻塞超时
var ErrPollTimeout = errors.New("poll timeout")

type consumer struct {
	consumer *kafka.Consumer
}

type ReceiveMsg struct {
	Val string `json:"msg"`
	Err error  `json:"err"`
}

// 提供对外操作
type Consumer interface {
	Receive(ctx context.Context, topic string, timeout int) <-chan ReceiveMsg
	Close()
}

func InitConsumer(c CConf) (Consumer, error) {
	conn, err := NewConsumer(c)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewConsumer(c CConf) (*consumer, error) {
	if "" == c.GroupId {
		return nil, fmt.Errorf("new kafka consumer error: groupid empty")
	}

	conf := kafka.ConfigMap{
		"bootstrap.servers": strings.Join(c.Hosts, ","),
		"group.id":          c.GroupId,
	}

	// 和php端保持一致
	conf["enable.auto.commit"] = true
	conf["topic.auto.offset.reset"] = "largest"

	if true == c.EnableSASL {
		conf["sasl.mechanisms"] = "PLAIN"
		conf["security.protocol"] = "sasl_plaintext"
		conf["sasl.username"] = c.SASLPlainUsername
		conf["sasl.password"] = c.SASLPlainPassword
	}

	kc, err := kafka.NewConsumer(&conf)
	if nil != err {
		return nil, fmt.Errorf("new kafka consumer error: %v", err)
	}

	return &consumer{
		consumer: kc,
	}, nil
}

// 从topic中接收消息
func (c *consumer) Receive(ctx context.Context, topic string, timeout int) <-chan ReceiveMsg {
	ch := c.Receives(ctx, []string{topic}, timeout)

	return ch
}

// 从多个topic中接收消息
func (c *consumer) Receives(ctx context.Context, topics []string, timeout int) <-chan ReceiveMsg {
	ch := make(chan ReceiveMsg, 1)

	go func() {
		defer func() {
			if err := recover(); nil != err {
				log.Errorf(
					"kafka receive异常. topics: %v, err: %v, stack: %v",
					topics,
					err,
					common.GetTraceStack(10),
				)
			}

			close(ch)
		}()

		if 0 == len(topics) {
			ch <- ReceiveMsg{
				Err: fmt.Errorf("receive error: topic empty"),
			}
			return
		}

		err := c.consumer.SubscribeTopics(topics, nil)
		if nil != err {
			ch <- ReceiveMsg{
				Err: fmt.Errorf("receive error: %v", err),
			}
			return
		}

		run := true
		for run == true {
			select {
			case <-ctx.Done():
				run = false
			default:
				receiveMsg := ReceiveMsg{}

				ev := c.consumer.Poll(timeout * 1000)
				if nil == ev {
					receiveMsg.Err = ErrPollTimeout
					ch <- receiveMsg
					continue
				}

				switch e := ev.(type) {
				case *kafka.Message:
					receiveMsg.Val = gjson.ParseBytes(e.Value).Get("body").String()
					ch <- receiveMsg
				case kafka.Error:
					typ := "unknown"
					if e.IsFatal() {
						typ = "Fatal"
					} else if e.IsRetriable() {
						typ = "retriable"
					} else if e.TxnRequiresAbort() {
						typ = "txn requires abort"
					}

					receiveMsg.Err = fmt.Errorf("receive err: %v, code: %v, type: %v", e.Error(), e.Code(), typ)
					ch <- receiveMsg
				default:
					log.Debugf("接收到了未知信息. msg: %+v", e)
				}
			}
		}
	}()

	return ch
}

// 关闭kafka消费者
func (c *consumer) Close() {
	if nil == c.consumer {
		return
	}

	err := c.consumer.Close()
	if nil != err {
		log.Errorf("kafka消费者关闭失败. err: %v", err)
	}
}
