package main

import (
	"fmt"
	"reflect"
)

func main(){

	var x float64 = 3.14

//	v := reflect.ValueOf(x)
//	v.SetFloat(7.1)

	v := reflect.ValueOf(&x)

	v.Elem().SetFloat(7.3)

	fmt.Println(x)

}
