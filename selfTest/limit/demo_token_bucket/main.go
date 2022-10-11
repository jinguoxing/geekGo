package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func main(){

	e := gin.Default()
	// 新建一个限速器，允许突发 10 个并发，限速 3rps，超过 500ms 就不再等待
	e.Use(NewLimiter(3, 10, 5*time.Millisecond))

	e.GET("/token_bucket", func(context *gin.Context) {

		time.Sleep(600*time.Millisecond)
		context.String(http.StatusOK,"pong")
	})

	e.Run(":9527")
}


func NewLimiter(r rate.Limit,b int,t time.Duration) gin.HandlerFunc {

	limiters := sync.Map{}

	return func(c *gin.Context) {

		key := c.ClientIP()
		l ,_ := limiters.LoadOrStore(key,rate.NewLimiter(r,b))

		ctx ,cancel := context.WithTimeout(c,t)
		defer cancel()

		if err:= l.(*rate.Limiter).Wait(ctx);err!=nil{

			c.AbortWithStatusJSON(http.StatusTooManyRequests,gin.H{"err":err})
		}
		c.Next()
	}

}
