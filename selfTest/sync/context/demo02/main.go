package main

import (
	"context"
	"fmt"
	"time"
)

func main(){

	//ctx , cancelFunc := context.WithCancel(context.Background())
	ctx, cancelFunc := context.WithTimeout(context.Background(), 8*time.Second)


	go watch(ctx,"[监控1]")
	go watch(ctx,"[监控2]")
	go watch(ctx,"[监控3]")

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancelFunc()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}


func watch(ctx context.Context,name string) {

	for {
		select {
			case <- ctx.Done():
				fmt.Println(name,"监控退出，服务停止",ctx.Err())
				return
		default:
			fmt.Println(name, "监控...")
			time.Sleep(2 * time.Second)
		}
	}
}
