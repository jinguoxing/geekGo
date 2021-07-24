package main

import (
	"context"
	"fmt"
)

type options struct {
	id string
	name string
	ctx context.Context
}
type App struct {
	opts     options
	ctx  	 context.Context
	cancel 	 func()
}
type Option func(o *options)

func ID(id string)Option {

	return func(o *options) {
		o.id = id
	}
}
func Name(name string)Option{

	return func(o *options) {
		o.name = name
	}
}


func New(opts ...Option)*App{

	options := options{
		ctx : context.Background(),
	}

	for _ ,f := range opts {
		f(&options)
	}
	ctx,cancel := context.WithCancel(options.ctx)
	return &App{
		opts: options,
		ctx: ctx,
		cancel: cancel,
	}

}

func main(){
    m := 	New(
    	ID("id111"),
    	Name("name1111"))
    fmt.Println(m)
}

