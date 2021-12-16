package main

import (
	"context"
	"fmt"
)

func main(){

	ctx := context.TODO()

	ctx = context.WithValue(ctx,"key1","value1")
	ctx = context.WithValue(ctx,"key2","value2")
	ctx = context.WithValue(ctx,"key3","value3")
	ctx = context.WithValue(ctx,"key4","value4")

	fmt.Println(ctx.Value("key1"))

}
