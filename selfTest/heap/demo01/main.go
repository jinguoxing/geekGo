package main

import "fmt"

func dummy(b int)int {

	 var c int
	 c = b
	 return c
}

func void(){

}

func main(){

	var a int

	void()

	fmt.Println(a,dummy(0))

}
