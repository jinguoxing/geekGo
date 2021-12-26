package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

const (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

func main(){

	cli ,err := clientv3.New(clientv3.Config{
		Endpoints : []string{"127.0.0.1:2379"},
		DialTimeout: dialTimeout,
	})

	if err != nil {
		fmt.Println(err)
	}

	defer cli.Close()


	ctx ,cancel := context.WithTimeout(context.Background(),requestTimeout)

	//_,err =  cli.Put(ctx,"key2","value2")
	res ,err :=  cli.Get(ctx,"key2")
	cancel()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)




}
