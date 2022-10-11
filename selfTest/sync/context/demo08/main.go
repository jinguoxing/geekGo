package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// 30个任务，5个worker
// 单个任务最长执行时间 5秒
// 完成所有任务或超时60秒退出

// 解析：建立 5个协程 ，每个协程完成6个任务
// 单个任务执行时间超过5秒，则记录一个error
//
var (
	tasks   = 30
	workers = 5
	wg      sync.WaitGroup
)

func main(){

	fmt.Println("任务开始",time.Now().Format("2006-01-02 15:04:05"))
	sTasks :=  tasks/workers

	ctx, cancelFunc := context.WithTimeout(context.Background(),60*time.Second)
	defer cancelFunc()
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go workLogic(i, sTasks)
	}

	go func() {
		defer func() {
			fmt.Println("go 协程退出:",time.Now().Format("2006-01-02 15:04:05"))
		}()

		for {
			select {
				case <-ctx.Done():
					fmt.Println("Done:",time.Now().Format("2006-01-02 15:04:05"))
				   return
			default:
				time.Sleep(70 * time.Second)
				fmt.Println("default:",time.Now().Format("2006-01-02 15:04:05"))

			}
		}
	}()

	// 任务编排  所有任务都完成后，整个退出
	wg.Wait()
	fmt.Println("任务结束：",time.Now().Format("2006-01-02 15:04:05"))
}

func workLogic(i, sTasks int) {
	defer wg.Done()
	// 每个worker 里面处理的任务数
	for j := 1; j <= sTasks; j++ {
		handle(i, j, sTasks)
	}
}

//模拟单个任务的处理逻辑
func handle(i, j ,m int) {

	// 单个任务超时5秒
	timeOut := time.After(5 * time.Second)

	data, err := logic(i, j, m)
	res := result{data: data, err: err}

	select {
	case <-timeOut:
		fmt.Printf("任务号:%d,错误原因：%+v\n", res.data,res.err)
		return
	default:
		fmt.Printf("任务号:%d ,错误原因:%+v\n", res.data, res.err)
		return
	}
}


type result struct {
	data int
	err error
}

// 模拟单个任务的处理逻辑
func logic(i, j, m int) (int, error) {
	// 模拟业务的处理时间，5的倍数的超过 6秒 ，其他的默认1秒
	t := 1
	var err error

	res := (i-1)*m + j

	if res%5 == 0 {
		t = 6
		err = errors.New(fmt.Sprintf("任务号%d 模拟任务超时", res))
	}
	time.Sleep(time.Duration(t) * time.Second)

	return res, err
}