package main

import (
	"fmt"
	"sync"
	"time"
)

func initApp(){
	go func() {
		fmt.Println(time.Now())
	}()
}


func main(){

	var once sync.Once

	once.Do(initApp)
	//go func() {
	//	initApp()
	//}()
	//
	////time.Sleep(2 * time.Second)
	//go func() {
	//	initApp()
	//}()
	once.Do(initApp)

	time.Sleep(10*time.Second)
}
