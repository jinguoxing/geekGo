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

type Php struct {

}

func(p Php) SayHello(){
	fmt.Println("Hi,I am PHP!")
}


func main()  {

	golang := Go{}
	php := Php{}

	//golang.SayHello()
	//php.SayHello()
	sayHello(golang)
	sayHello(php)


}
