package main

import (

	"fmt"
	"geekGo/selfTest/ckafka"
)

func main(){


	 ckafka.InitProducer(ckafka.PConf{
		Hosts:             []string{"101.133.146.180:9093","101.133.158.246:9093"},
		EnableSASL : true,
		SASLPlainUsername: "alikafka_pre-cn-2r42f9yl6004",
		SASLPlainPassword: "Y032MGMQLIkZwlK9uqNbrwE4wJV34BYk",
	})


	err := ckafka.Send("kingtest_001","hello golang")

	fmt.Println(err)
}
