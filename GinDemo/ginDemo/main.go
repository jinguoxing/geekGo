package main

import (
	//"fmt"
	//"geekGo/GinDemo/ginDemo/config"
	//"geekGo/GinDemo/ginDemo/router"
	//"github.com/gin-gonic/gin"
	//"github.com/tidwall/gjson"
	"fmt"
	"encoding/json"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	//engine := gin.New()
	//router.InitRouter(engine) // 设置路由
	//

	const jsonArr = `[1003001424713059,"1003001424713058"]`

	//skuArr := []int64{1003001424713059,1003001424713058}

	skuArr := make([]int64,0)

	//SkuMap := make(map[string])

	//data, err := json.Marshal(skuArr)
	//if err != nil {
	//	fmt.Printf("json.marshal failed, err:", err)
	//	return
	//}

	//fmt.Println(string(data))

	err := json.Unmarshal([]byte(jsonArr), &skuArr)

	json.Valid([]byte(jsonArr))
	if err != nil {

		fmt.Println(err)
	}else{
		fmt.Println(skuArr)
	}
	//fmt.Println(err)

		//value := gjson.Get(string(data), "")
		//fmt.Println(value.Array())



	//err := engine.Run(config.PORT)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("已经启动")
}
