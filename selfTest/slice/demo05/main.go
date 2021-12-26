package main

import "fmt"

func main() {
	s := []int{1, 1, 1}
	f1(&s)
	fmt.Println(s)
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
		i++
	}
	*/

	for i := range s {
		s[i] += 1
	}
	s = append(s,100)
	fmt.Println(s)
}

func f1(s *[]int){

	*s = append(*s,100)
	return
}
