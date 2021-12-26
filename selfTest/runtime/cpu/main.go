package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

func main(){

	cpuInfos ,err := cpu.Info()

	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)

	}

	for index ,ci := range cpuInfos {
		fmt.Println(index,ci)
	}


}
