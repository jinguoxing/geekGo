package main

import "fmt"

func main(){


	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:6]

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	s2 = append(s2,100)
	s2 = append(s2,200)
	//s2 = append(s2,300)

	s1[2] = 1000

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(slice, len(slice), cap(slice))


	x := []int{1,2}
	x = append(x,4,5,6)
	fmt.Printf("len=%d, cap=%d",len(x),cap(x))

	newS := myAppend(slice)
	fmt.Println(newS, len(newS), cap(newS))
	fmt.Println(slice, len(slice), cap(slice))

}


func myAppend(s []int) []int {

	s = append(s,100)
	return  s

}