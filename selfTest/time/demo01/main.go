package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func main(){



	wg.Add(2)

	//timer1 := time.NewTimer(2 * time.Second)
	//
	//ticker1 := time.NewTicker(2*time.Second)
	//
	//go func(t *time.Ticker) {
	//
	//	wg.Done()
	//	for {
	//			<-t.C
	//			fmt.Println("get ticker1",time.Now().Format("2006-01-02 15:04:05"))
	//	}
	//}(ticker1)
	//
	//go func(t *time.Timer) {
	//	wg.Done()
	//	for {
	//		<-t.C
	//
	//		fmt.Println("get timer1",time.Now().Format("2006-01-02 15:04:05"))
	//
	//		t.Reset(2 * time.Second)
	//	}
	//
	//}(timer1)

	go getTicker()
	go getTimer()

	wg.Wait()

	fmt.Println("完成了 时间",time.Now().Format("2006-01-02 15:04:05"))

	select {

	}

}


func getTicker(){
	ticker1 := time.NewTicker(2*time.Second)
	defer wg.Done()
	defer ticker1.Stop()
	for {
		<-ticker1.C
		fmt.Println("get ticker1",time.Now().Format("2006-01-02 15:04:05"))
	}
}

func getTimer(){

	timer1 := time.NewTimer(2 * time.Second)
	defer wg.Done()
	defer timer1.Stop()
		for {
			<-timer1.C

			fmt.Println("get timer1",time.Now().Format("2006-01-02 15:04:05"))

			timer1.Reset(2 * time.Second)
		}

}