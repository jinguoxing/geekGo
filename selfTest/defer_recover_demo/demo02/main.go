package main

import "fmt"

func main() {

	//for i := 0; i <= 5; i++ {
	//	defer fmt.Println(i)
	//}
	//
	fmt.Println(increase(1))
	fmt.Println(f1())

}


func increase(d int)(ret int){


	defer func(ret int) {
		ret++
	}(ret)
	return d
}

func f()(r int){

	t :=5
	defer func() {
		t = t+5
	}()

	return t
}

func f1()(r int)  {

	//defer func(r int) {
	//	r = r+5
	//}(r)

	defer func() {
		r = r+5
	}()

	return  1
}




