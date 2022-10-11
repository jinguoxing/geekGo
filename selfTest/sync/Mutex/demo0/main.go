package main

import (
	"fmt"
	"sync"
)

var (
	sum int
	wg sync.WaitGroup
)

func add(i int){
	defer wg.Done()
	sum = sum +i
}

func main()  {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go add(i)
	}
	wg.Wait()
	fmt.Println("sum 和为：",sum)
}