package main

import "fmt"

func main(){

	sl := make([]int, 0, 2)
	//var appenFunc = func(s []int) {
	//	s = append(s, 10, 20, 30)
	//	fmt.Println(s)
	//}

	fmt.Printf("111====%p",&sl)
	fmt.Println(sl)

	//fmt.Println(sl[:3])
	sl = append(sl,1,2,3,4)
	//sl[0]  = 100
	fmt.Printf("111===%p",&sl)
	fmt.Println(sl[:])
	fmt.Printf("333===%p",&sl)
	//appenFunc(sl)
	fmt.Println(sl)
	//fmt.Println(sl[:10])


}
