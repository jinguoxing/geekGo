package main

import (
	"fmt"
	"time"

	//"github.com/confluentinc/confluent-kafka-go/kafka"
	"strings"
)

var replacer = strings.NewReplacer("Ö", "O", "ü", "u", "ö", "o", "Ü", "U", "Ä", "A", "ä", "a")


func FormatNameAsKey1(name string) string {
	name = strings.TrimRight(name, " \t\n")
	name = strings.ToUpper(name)
	fmt.Println(name)
	// 替换特殊字符
	name = replacer.Replace(name)

	return name
}

func testFunc(){

	defer deferFunc()
	panic("这是测试的函数")
}

func deferFunc()  {

	if panic_ := recover(); nil != panic_ {
		fmt.Errorf("%v", panic_)
	}

}


func main() {

	str :="ö"
	//go testFunc()
	fmt.Println(time.Now())
	after := time.After(10*time.Second)
	done := make(chan bool)
	done1 := make(chan struct{})
	go func() {
		defer func() {
			if err := recover(); nil != err {
				fmt.Errorf("程序异常, eccn缓存处理失败. err: %v", err)
			}
			done <- true
		}()


	}()

	select {
	//case <-done:
	//	fmt.Println("done 退出")
	case <-done1:
		//fmt.Println("done 退出")
	case <-after:
		//fmt.Println("after 退出")
	}
	fmt.Println(time.Now())
	fmt.Println(FormatNameAsKey1(str))



	//p, err := kafka.NewProducer(&kafka.ConfigMap{
	//	"bootstrap.servers": "10.8.11.191:9091"})
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer p.Close()
	//
	//// Delivery report handler for produced messages
	//go func() {
	//	for e := range p.Events() {
	//		switch ev := e.(type) {
	//		case *kafka.Message:
	//			if ev.TopicPartition.Error != nil {
	//				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	//			} else {
	//				fmt.Printf("Delivered message to %v\n,信息内容: %v\n", ev.TopicPartition,string(ev.Value))
	//			}
	//		}
	//	}
	//}()
	//
	////// Produce messages to topic (asynchronously)
	//topic := "myTopic"
	//for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	//	p.Produce(&kafka.Message{
	//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//		Value:          []byte(word),
	//	}, nil)
	//}
	//
	//// Wait for message deliveries before shutting down
	//p.Flush(30 * 1000)
}
