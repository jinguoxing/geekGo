package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var counter int

func main() {

	for i := 1; i <= 2; i++ {

		wg.Add(1)


		go routine(i)
	}

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func routine(id int){

	fmt.Println(id)

	//for i := 0; i < 1; i++ {
	//	mu.Lock()
		counter++
	//	mu.Unlock()
	//}
	wg.Done()
}
