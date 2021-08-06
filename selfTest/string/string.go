package main

import "fmt"


// go tool compile -N -l -S ./string.go
func main(){

	str := "kingnet"

	by := []byte(str)

	fmt.Println(by)

}
