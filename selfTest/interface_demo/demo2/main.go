package main

import "fmt"

type coder interface {
	code()
	debug()
}


type Gopher struct {
	language string
}

func(p Gopher) code(){

	fmt.Printf("I am coding %s language\n",p.language)
}

func(p *Gopher) debug(){

	fmt.Printf("I am debug %s language\n",p.language)
}

func main()  {

	var c coder = &Gopher{"Go"}
	//var c coder = Gopher{"Go"} // Gopher 类型并没有实现 debug 方法
	c.code()
	c.debug()


}