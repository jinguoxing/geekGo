package main

import "fmt"

func main(){

	defer func() {

		if err := recover();err != nil {

			fmt.Println(err)
		}

	}()

	 // map声明
     var newMap map[int]string

    // 对未初始化的map 赋值会报 panic
    //  newMap[1] = "king"

    // 通过此方式 可以正常赋值
    newMap = map[int]string{1:"xx1",2:"xx2"}

    fmt.Println(newMap)

    newMap1 := make(map[int]string)

    newMap1[1] = "king"

    getMap , exit := newMap1[1]

    if exit {

		fmt.Println(getMap)
	}

	delete(newMap,1)
	delete(newMap,2)

    fmt.Println(newMap)


}
