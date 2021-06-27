package main

import (
	"fmt"
	"unsafe"
)

func main(){

	ageMap := make(map[string]int ,10)

	fmt.Printf("长度：%d ",unsafe.Sizeof(ageMap))

	ageMap["key1"] = 1
	ageMap["key2"] = 2
	ageMap["key3"] = 3

	fmt.Printf("长度：%d ",unsafe.Sizeof(ageMap))


	var a,b string

	b = "hello world"



	fmt.Printf("长度：%d \r\n ",unsafe.Sizeof(a))
	fmt.Printf("字节长度：%d，长度： %d \r\n",unsafe.Sizeof(b), len(b))

}