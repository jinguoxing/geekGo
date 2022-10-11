package main

import (
	"fmt"
	"runtime"
)

func query() int {

	ch := make(chan int,1)

	for i := 0; i < 1000; i++ {
		go func() {
			ch <- 0
		}()
	}

	select {

		case <-ch:

	default:


	}

	return 0

}




func main(){

	for i := 0; i < 4; i++ {

		query()
		fmt.Println("goroutines: %d\n", runtime.NumGoroutine())
	}

}
