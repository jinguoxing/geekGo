package main

import (
    "context"
    "fmt"
)

func main(){

        ctx := context.TODO()

        ctx = context.WithValue(ctx,"key1","val1")
        ctx = context.WithValue(ctx,"key2","val2")
        ctx = context.WithValue(ctx,"key3","val3")
        ctx = context.WithValue(ctx,"key4","val4")

        fmt.Println(ctx.Value("key1"))

<<<<<<< HEAD

=======
>>>>>>> 621a11f5f2a25957d7b36554b3a7fde55be6f7ab
}