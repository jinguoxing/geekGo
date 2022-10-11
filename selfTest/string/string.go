package main

import (
	"fmt"
	"time"
)

// go tool compile -N -l -S ./string.go
func main() {

	str := "kingnet"

	by := []byte(str)

	fmt.Println(by)
	time.Sleep(1000)

}
