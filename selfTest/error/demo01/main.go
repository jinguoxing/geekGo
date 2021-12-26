package main

import (
	"errors"
	"fmt"
)

func main(){

	err1 := errors.New("this is a error")

	err2 := fmt.Errorf("this is a error")

	if err1 == err2 {

		fmt.Println("the same error")
	}else {
		fmt.Println("un same error")
	}




}
