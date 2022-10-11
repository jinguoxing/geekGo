package main

import (
	"fmt"
	"sync"
)

var (
	sum   int
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func add(i int) {

	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	sum = sum + i
}

func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go add(i)
	}
	wg.Wait()
	fmt.Println("sum 和为：", sum)
}
