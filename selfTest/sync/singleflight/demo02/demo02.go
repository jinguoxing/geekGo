package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var count int32

func getArticle(id int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)

	return fmt.Sprintf("article: %d", id), nil
}

func singleflightGetArticle(ctx context.Context,sg *singleflight.Group, id int) (string, error) {
	result := sg.DoChan(fmt.Sprintf("%d", id), func() (interface{}, error) {
		select {}
		return getArticle(id)
	})

	select {
		case r:=<-result:
			return r.Val.(string),r.Err
		case <-ctx.Done():
			return "",ctx.Err()
	}
}

func main() {

	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sg  = &singleflight.Group{}
	)

	ctx,cancel :=  context.WithTimeout(context.Background(),2*time.Second)

	defer cancel()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			 res, _ := singleflightGetArticle(ctx,sg, 1)
			//res, _ := getArticle(1)
			if res != "article: 1" {
				panic("err")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))

}
