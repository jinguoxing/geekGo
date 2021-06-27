package main

import "fmt"

type Message struct {

	msg string
}

type Greeter struct {

	Message Message
}

type Event struct {

	Greeter Greeter
}


func NewMessage(m string) Message{

	return Message{
		msg: m,
	}
}

func NewGreeter(m Message) Greeter {

	return Greeter{
		Message: m,
	}
	
}


func NewEvent(g Greeter) Event {

	return Event{
		Greeter: g,
	}
}

func (g Greeter)Greet() Message {

	return g.Message
}

func(e Event)Start(){

	msg := e.Greeter.Greet()

	fmt.Println(msg)

}

// 未使用 wire
//func main(){
//
//
//	message := NewMessage("hello world")
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}

func main(){


	event := InitializeEvent("hello world")

	event.Start()

}


