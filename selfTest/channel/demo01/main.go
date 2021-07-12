package main

import (
	"time"
)

func  main()  {

	ch := make(chan string,6)

	// 方法一 ：
	//go func() {
	//	for  {
	//	v,ok := <-ch
	//
	//	if !ok{
	//
	//		fmt.Println("end")
	//		return
	//	}
	//	fmt.Println(v)
	//	}
	//}()

	//// 方法二：
	//done := make(chan struct{})
	//go func() {
	//
	//	for {
	//		select {
	//		case ch<-"":
	//
	//
	//	    case <-done:
	//	    	close(ch)
	//			return
	//		}
	//	}
	//}()


	ch <- "协程数据1"
	ch <- "协程数据2"
	ch <- "协程数据3"

	close(ch)

	time.Sleep(time.Second)
}
