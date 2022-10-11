package main

import (
	"fmt"
	"sync"
)

func main(){

	var m sync.Map
	// Store的存储
	m.Store("php",20)
	m.Store("mysql",10)
	// Load的读取
	age,_ := m.Load("php")
	fmt.Println(age.(int))

	// Range的方法
	m.Range(func(key, value interface{}) bool {

		name := key.(string)
		age := value.(int)
		fmt.Println(name,age)
		return true
	})

	// Delete 的方法
	m.Delete("php")
	m.Delete("golang")

	age ,ok := m.Load("php")
	fmt.Println("Delete 操作：",age,ok)

	// LoadOrStore的方法
	m.LoadOrStore("golang",100)
	age,_ = m.Load("golang")

	fmt.Println(age)

}
