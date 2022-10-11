package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func a() error {

	return errors.Wrap(fmt.Errorf(""),"a")
}


func b() error {
	return errors.Wrap(a(),"b")
}

func c() error  {

	return errors.Wrap(b(),"c")
}

func main(){

	//fmt.Printf("err :%+v",c())
	fmt.Println(c())
}