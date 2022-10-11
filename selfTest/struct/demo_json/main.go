package main

import (
	"encoding/json"
	"fmt"
)

type person struct {

	Name string `json:"name"`
	Age int `json:"age"`

}


func main(){


	p:= person{Name:"kingnet",Age: 20}

	jsonB,err := json.Marshal(p)

	if err ==nil {
		fmt.Println(string(jsonB))
	}

	//respJson :=


}
