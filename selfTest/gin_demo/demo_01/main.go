package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type testF struct {
	Test string `header:"test_1"`
}

type  testRe struct {
	Test map[int]testF
	Hello string `header:"hello"`
}

type testHeader struct {
	Rate int `header:"Rate"`
	Domain string `header:"domain"`
}

type testS struct {
	testRe
	testHeader
}


func main(){

	r := gin.Default()
	r.Use(Logger())

	yy := make(map[int]testF)
	yy[1]  = testF{ "test1"}

	r.GET("/test_header", func(c *gin.Context) {

		h := testS{}

		if err:= c.ShouldBindHeader(&h); err != nil {

			c.JSON(200,err)
		}

		h.Test = yy

		fmt.Printf("%#v\n", h)
		fmt.Println("")
	//	c.JSON(200,gin.H{"Rate":h.Rate,"Domain":h.Domain})
	//	key,_ := c.Get("st")
	//	c.SecureJSON(200,gin.H{"Rate":h.Rate,"Domain":h.Domain,"key":key})
		c.JSON(200,h)
	})

	r.Run()
}

func Logger() gin.HandlerFunc  {

	return func(c *gin.Context) {
		t := time.Now()

		c.Set("st",t)

		c.Next()

		latency := time.Since(t)

		log.Print(latency)
	}
}