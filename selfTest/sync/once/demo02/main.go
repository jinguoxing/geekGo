package main

import (
	"fmt"
	once2 "geekGo/selfTest/sync/once"
	"time"
)

func initApp(){
	go func() {
		fmt.Println(time.Now())
		//time.Sleep(10*time.Second)
		panic("error ")
	}()
}


func main(){

	var once once2.Once

	once.Do(func() error {
		fmt.Println("111111",once.Done())
		initApp()
		return nil
	})

	fmt.Println("22222",once.Done())

	once.Do(func() error {
		initApp()
		return nil
	})




time.Sleep(10*time.Second)
}
