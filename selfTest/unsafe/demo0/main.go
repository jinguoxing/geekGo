package main

import (
	"fmt"
	"unsafe"
)

func main()  {

	i := 10
	ip := &i
	fmt.Println(i)
	fmt.Println(ip)

	var fp *float64 =  (*float64)(unsafe.Pointer(ip))

	*fp = *fp *3

	fmt.Println(i)

}



