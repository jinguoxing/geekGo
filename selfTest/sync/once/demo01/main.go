package main

import (
<<<<<<< HEAD
    "fmt"
    "sync"
)

func main(){


    var once sync.Once

    f1 := func() {

        fmt.Println("in f1")
    }
    once.Do(f1)

    f2 := func() {
        fmt.Println("in f2")
    }

    once.Do(f2)

=======
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
>>>>>>> 11118a04bff017e022d29e66f963b48bce143c8d
}
