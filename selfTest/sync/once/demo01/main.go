package main

import (
    "fmt"
    "sync"
)

func main(){


    var once sync.Once

    f1 := func() {

        fmt.Println("in f1")
    }
    once.Do(f1)

    f2 := func() {
        fmt.Println("in f2")
    }

    once.Do(f2)

}
