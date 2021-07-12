package main

import (
	"fmt"
	"time"
)

func main(){

	ch := make(chan  string,4)

	done := make(chan struct{})


	go func() {
		for  {
			select {
			case ch <- "你好":
			case <-done:
				close(ch)
				return
			}
		}
	}()


	go func() {

		time.Sleep(3* time.Second)
		done <- struct{}{}
	}()

	for i:= range ch {

		fmt.Println("接收到的值：",i)
	}

fmt.Println("end")

}
