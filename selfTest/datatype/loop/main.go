package main

import "fmt"

func main()  {

	var out []*int8
	var i int8

	for i=0 ; i < 3; i++ {

		//iCopy := i
		out = append(out, &i)
		fmt.Println(&i,&i,out)
	}

	fmt.Println("Values: ",*out[0],*out[1],*out[2])

}
