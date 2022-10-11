package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 300个任务，5个worker
// 单个任务最长执行时间 5秒
// 完成所有任务或超时60秒退出

// 解析：建立 5个协程 ，每个协程完成6个任务
// 单个任务执行时间超过5秒，则记录一个error
//
var (
	tasks   = 300
	workers = 5
	wg      sync.WaitGroup
)

func main(){
	fmt.Println("任务开始",time.Now().Format("2006-01-02 15:04:05"))
	taskCh := make(chan struct{}, tasks)
	doneCh := make(chan struct{})
	for i := 0; i < tasks; i++{
		taskCh <- struct{}{}
	}
	close(taskCh)
	for i := 0; i < workers; i++{
		wg.Add(1)
		go worker(taskCh)
	}
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	select {
	case <- time.After(60*time.Second):
		fmt.Println("60秒超时退出")
	case <- doneCh:
		fmt.Println("所有任务完成，所有worker退出")
	}
	fmt.Println("任务结束",time.Now().Format("2006-01-02 15:04:05"))
}

func worker(ch <-chan struct{}) {
	defer wg.Done()
	for  {
		select {
		case task, ok := <-ch:
			if !ok{
				fmt.Println("worker 无可做任务")
				goto end
			}
			handle(task)
		}
	}
end:
}

func handle(task struct{}){
	done := make(chan struct{})
	go something(task, done)
	select {
	case <-time.After(5* time.Second):
		fmt.Println("5秒超时退出")
		return
	case <- done:
		fmt.Println("正常退出")
		return
	}
}

func something(task struct{}, done chan <- struct{} ){
	time.Sleep(time.Duration(rand.Intn(30)) * time.Second)
	done <- struct{}{}
}