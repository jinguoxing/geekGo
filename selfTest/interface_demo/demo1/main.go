package main

import "fmt"

type  IGreeting interface {
	SayHello()
}
func sayHello(i IGreeting){
	i.SayHello()
}



type Go struct {
}
func (g Go) SayHello(){

	fmt.Println("Hi ,I am Go!")
}

func main()  {

	golang := Go{}

	sayHello(golang)
}
