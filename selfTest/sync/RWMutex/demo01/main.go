package main

import (
	"sync"
	"time"
)

type counter struct {

	mu sync.RWMutex
	count uint64
}

func (c *counter) Incr() {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) Count() uint64 {

	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main(){

	var counter counter
	for i:=0 ;i<10;i++ {

		go func() {
			for {
				counter.Count()
				time.Sleep(time.Millisecond)
			}
		}()
	}


	for {
		counter.Incr()
		time.Sleep(time.Second)
	}

}
