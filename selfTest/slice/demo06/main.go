package main

import "fmt"

func main() {
	sl := make([]int, 0, 10)
	var appenFunc = func(s []int) {
		s = append(s, 10, 20, 30)
		fmt.Println(s)
	}
	fmt.Println(sl)
	appenFunc(sl)
//	sl = append(sl,1000)
	//sl[3] = 30
	fmt.Println(sl)
	fmt.Println(sl[:])
	fmt.Println(sl[:10])
}
