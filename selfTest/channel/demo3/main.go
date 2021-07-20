package main

import (
    "context"
    "fmt"
    "time"
)

//d第三种方法：context

func main(){

    ch := make(chan struct{})
    ctx,cancel := context.WithCancel(context.Background())


    go func(ctx context.Context) {

        for {
            select {

            case <- ctx.Done():
                    ch <- struct{}{}
                    return
            default:
                fmt.Println("default")
            }

            time.Sleep(500*time.Millisecond)
        }
    }(ctx)

    go func() {
        fmt.Println(time.Now())
        time.Sleep(3*time.Second)
        cancel()
        fmt.Println(time.Now())
    }()

    <- ch
    fmt.Println("end")




}
