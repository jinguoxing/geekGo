package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {

	var slice0 []int

	slice1:= *new([]int)
	slice2 := make([]int,5,10)


	fmt.Printf("数据结构长度：%d, %p,%d,%d\r\n",unsafe.Sizeof(slice0),&slice0,len(slice0),cap(slice0))
	fmt.Printf("数据结构长度：%d, %p\r\n", (unsafe.Sizeof(slice1)),&slice1)
	fmt.Printf("数据结构长度：%d, %p\r\n",unsafe.Sizeof(slice2),&slice2)

	var ti time.Duration

	fmt.Println(ti)




}