package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main(){

	 mutex := &sync.Mutex{}
	c := sync.NewCond(mutex)

	var ready int

	for i:=0;i<=10;i++ {

		go func(i int) {

			time.Sleep(time.Duration(rand.Int63())*time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已经准备就绪",i)
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready!=10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	log.Println("所有运动员都准备就绪.....")


}
