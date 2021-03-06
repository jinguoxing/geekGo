package main

import (
	"context"
    "fmt"
    "time"
)

func main(){

    ctx ,cancel := context.WithCancel(context.Background())
    defer cancel()

    fmt.Println(time.Now())

    go func() {

        defer func() {
            fmt.Println("goroutine exit")
        }()

        for {
            select {

            case <-ctx.Done():
                fmt.Println(ctx.Err())
                return
            default:
                time.Sleep(1*time.Second)
                fmt.Println("default")
                fmt.Println(time.Now())
            }
        }
    }()


    fmt.Println(time.Now())
    time.Sleep(time.Second)

    time.Sleep(2 * time.Second)
    fmt.Println(time.Now())

    select {

    }

}