package main

import "fmt"

//Q：给定数组A[1...N]，满足: N >= 3, A[1] > A[2],  A[N] > A[N-1]，并且所有元素的值都不相同。
//请在O(logN)时间内，找到数组中的一个元素（下标为i）满足A[i-1] > A[i] < A[i+1]。


func main()  {
	fmt.Println(findX([]int{2,1,3,4}))
	fmt.Println(findX([]int{2,1,7,6,3,4,9,10}))

	fmt.Println(findX([]int{7,6,3,4,9,10,1,2}))
}

func findX(s []int) int {
	arrLen := len(s)
	midLen := arrLen/2
	if arrLen >= 3 && s[0] > s[1] && s[arrLen-1] > s[arrLen-2] {
		if s[midLen] < s[midLen+1] && s[midLen-1] > s[midLen] {
			return midLen
		} else {
			return findI(s[:midLen+1])
		}
	} else {
		return 0
	}
}

func findI(s []int) int {
	lt := 0
	arrLen := len(s)
	rt := arrLen - 1
	midLen := (rt-lt)/2 + lt

	if arrLen >= 3 && s[0] > s[1] && s[rt] > s[rt-1] {
		if s[midLen] < s[midLen+1] && s[midLen-1] > s[midLen] {
			return midLen
		}else  {
			return findI(s[:midLen+1])
		//}else {
		//	return findI(s[midLen:])
		}
	}else {
		return -1
	}
}