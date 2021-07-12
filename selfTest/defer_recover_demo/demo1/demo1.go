package main

import (
	"fmt"
)

type number int

func (n number) printf()  {
	fmt.Println(n)
}

func (n *number) pprintf(){
	fmt.Println(*n)
}

func main() {
	defer fmt.Println("defer main")
	//var user = os.Getenv("USER_")

	var n number


	defer n.pprintf()

	defer n.printf()

	defer func() {
		n.printf()
	}()

	defer func() {
		n.pprintf()
	}()

	n = 1


	//go func() {
	//	defer func() {
	//		fmt.Println("defer caller")
	//		if err := recover(); err != nil {
	//			fmt.Println("recover success. err: ", err)
	//		}
	//	}()
	//
	//	func() {
	//		defer func() {
	//			fmt.Println("defer here")
	//		}()
	//
	//		if user == "" {
	//			panic("should set user env.")
	//		}
	//
	//		// 此处不会执行
	//		fmt.Println("after panic")
	//	}()
	//
	//
	//	fmt.Println("=========================")
	//}()

	//mytest()
	//defer defTest()
	//
	////fmt.Println(f())
	//fmt.Println(f1())
	//
	//time.Sleep(100)
	//fmt.Println("end of main function")
}


func mytest(){

	panic("这是一个panic的错误")

}

func defTest(){

	if err := recover(); err != nil {
		fmt.Println("recover success. err: ", err)
	}
}

func f()(r int){

	t := 5

	defer func() {
		t= t+5
	}()

	return t

}

func f1()(r int){

	defer func(r int) {

		r = r+5
	}(r)

	return 1

}
