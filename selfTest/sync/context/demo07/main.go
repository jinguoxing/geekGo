package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func f1(ctx context.Context) error {

	select {

	case <- ctx.Done():
		return fmt.Errorf("f1: %w",ctx.Err())

	case <- time.After(10*time.Millisecond):
		return fmt.Errorf("f1 err in 10ms")

	}

}

func f2(ctx context.Context)error {

	select {

	case <-ctx.Done():
		return fmt.Errorf("f2: %w", ctx.Err())
	case <-time.After(30 * time.Millisecond):
		return fmt.Errorf("f2 err in 10ms")
	}

}


func main(){

	g,ctx := errgroup.WithContext(context.Background())

	g.Go(func()(err error) {

		if err := f1(ctx); err != nil {
		return 	fmt.Errorf("f1执行后的错误 %w",err)
		}else{
			return nil
		}
	})

	g.Go(func() error {

		if err := f2(ctx); err != nil {
			return 	fmt.Errorf("f2执行后的错误 %w",err)
		}else{
			return nil
		}
	})


	if err := g.Wait(); err!=nil{
		fmt.Println(err)
	}


}





