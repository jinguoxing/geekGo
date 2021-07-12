package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func main()  {

	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)

	ctx, canceFunc := context.WithCancel(context.Background())

	//ctx,canceFunc := context.WithTimeout(context.Background(),1*time.Second)
	//ctx,canceFunc := context.WithDeadline(context.Background(),time.Now().Add(2*time.Second))

	fmt.Println(time.Now())
	for i := 1; i <= total; i++ {

		go addNum(&num, i, func() {

			if atomic.LoadInt32(&num) == int32(total) {
				canceFunc()
			}
		})
	}
	<-ctx.Done()
	fmt.Println("END.",time.Now())
}

func addNum(nump *int32,id int,deferFunc func()){

		defer deferFunc()

		for i:=0;;i++{

			currNum := atomic.LoadInt32(nump)
			newNum := currNum+1
			time.Sleep(time.Millisecond*200)
			if atomic.CompareAndSwapInt32(nump,currNum,newNum){

				fmt.Printf("The number: %d [%d-%d-%d]\n", newNum, currNum,id, i)
				break
			}
		}


}
