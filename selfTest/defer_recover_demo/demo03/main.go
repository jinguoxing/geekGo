package main

import (
	"errors"
	"fmt"
)


func f1() {
	var err error

	defer fmt.Println(err)

	err = errors.New("defer error1")
	return
}

func f2() {
	var err error

	defer func() {
		fmt.Println(err)
	}()

	err = errors.New("defer error2")
	return
}

func f3() {
	var err error

	defer func(err error) {
		fmt.Println(err)
	}(err)

	err = errors.New("defer error3")
	return
}

func main() {
	f1()
	f2()
	f3()

	fmt.Println(1<<0)
	fmt.Println(1<<2)



}


