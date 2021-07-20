package main

import (
	"fmt"
	"sync/atomic"
)

func main(){


	defer func() {

		if err:= recover(); err!=nil {

			fmt.Println(err)
		}
	}()

	var i32 int32

	fmt.Println("=====old i32 value=====")
	fmt.Println(i32)

	//第一个参数值必须是一个指针类型的值,因为该函数需要获得被操作值在内存中的存放位置,以便施加特殊的CPU指令
	//结束时会返回原子操作后的新值
	newI32 := atomic.AddInt32(&i32,3)
	fmt.Println("=====new i32 value=====")
	fmt.Println(i32)
	fmt.Println(newI32)
	fmt.Println(atomic.LoadInt32(&i32))


	var i64 int64

	fmt.Println("=====old i64 value=====")
	fmt.Println(i64)
	newI64 := atomic.AddInt64(&i64,-3)
	fmt.Println("=====new i64 value=====")
	fmt.Println(i64)
	fmt.Println(newI64)


	fmt.Println("======Store value=======")
	atomic.StoreInt32(&i32,10)
	fmt.Println(i32)
	fmt.Println(newI32)


	fmt.Println("======Swap value=======")
	old32 := atomic.SwapInt32(&i32,20)
	fmt.Println(old32)
	fmt.Println(i32)


}

