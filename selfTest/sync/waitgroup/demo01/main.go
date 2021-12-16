package main

import (
	"fmt"
	"sync"
	"time"
)
func main(){
	var wg sync.WaitGroup
	fmt.Println("开始让他们处理任务", time.Now())
	wg.Add(3)
	wg.Add(-1)

	go func() {
		defer func() {
			wg.Done()
		}()
		time.Sleep(time.Second)
		fmt.Println("这是一个协程，等待了1s", time.Now())

	}()

	go func() {
		defer func() {
			wg.Done()
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("这是一个协程，等待了2s", time.Now())
	}()


	wg.Wait()
	fmt.Println("终于等他们都处理完成",time.Now())
}
