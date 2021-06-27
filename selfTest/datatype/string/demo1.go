package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){

	s:= "你好,中国!"
	s1:= "Hello china!"

	fmt.Println(len(s),utf8.RuneCountInString(s))
	fmt.Println(len(s1),utf8.RuneCountInString(s1))

	fmt.Println()


}
