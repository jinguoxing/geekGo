package main

import (
    "fmt"
    "math/big"
    "sync"
)

var threeOnce struct{

    sync.Once
    v *big.Float
}

func three() *big.Float {



        threeOnce.Do(func() {
            threeOnce.v =   big.NewFloat(3.0)
        })
        return threeOnce.v
}

func main(){

    fmt.Println(three())
    fmt.Println(three())

}

