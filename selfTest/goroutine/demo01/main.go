package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)




func hardWork(job interface{}) error {
	//time.Sleep(time.Minute)
	panic("oops")
	return nil
}

func requestWork(ctx context.Context,job interface{})error {

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	defer cancel()

	done := make(chan error,1)
	panicChan := make(chan interface{},1)

	go func() {

		defer func() {
			if p := recover(); p != nil {
				panicChan <-p
				fmt.Println("oops,panic")
			}
		}()

		done <- hardWork(job)
	}()

	select {

	case err:= <-done:
		return err
	case p:=<-panicChan:
		panic(p)
	case <- ctx.Done():
		return ctx.Err()
	}

}


func main()  {

	const total = 1000
	var wg sync.WaitGroup

	wg.Add(total)
	now := time.Now()

	for i:=0;i<total;i++{
		go func() {

			defer func() {

				if p:=recover();p!=nil{
					fmt.Println("oops panic")
				}
			}()

			defer wg.Done()
			requestWork(context.Background(),"any")
		}()
	}

	wg.Wait()

	fmt.Println("花费时间:",time.Since(now))

	time.Sleep(time.Minute*2)
	fmt.Println("Goroutine numbers:",runtime.NumGoroutine())

}
