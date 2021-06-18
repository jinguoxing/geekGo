package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/errgroup"
	"math"
	"strconv"
	"time"
)

var (
	ctx = context.Background()
	rdb *redis.Client
	strBase = "a"
)


func init(){

	rdb = redis.NewClient(
		&redis.Options{
			Addr:         ":6379",
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
		})
}


func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "xxxx", // use default Addr
		Password: "xxx",               // no password set
		DB:       10,                // use default DB
		//PoolSize:     20,
	})


	var valueLen = 50
	var (
		maxNum float64
		loopNum float64
	)
	//最大的数据量
	maxNum = 500000
	//循环的次数
	loopNum = 1000

	redisKey, redisValue := strCom(valueLen)
	startMemory ,err := rdb.Info(ctx,"memory").Result()

	if err == nil {
		fmt.Printf("开始时间: %s ,key的前缀为 %s; Value的长度：%d， 内存统计开始: \r\n", time.Now(),  redisKey,len(redisValue))
		fmt.Println(startMemory)
	}

	g := new(errgroup.Group)

	e := int(math.Ceil(maxNum/loopNum))

	for s := 0; s < e; s++ {

		getS := s

		g.Go(func() error {

			begin := getS* int(loopNum) + 1

			end := (begin-1) + int(loopNum)

			if end >= int(maxNum) {
				end = int(maxNum)
			}

			for i := begin; i <= end; i++ {

				err := rdb.Set(ctx,redisKey+strconv.Itoa(i), redisValue,0).Err()

				if err != nil {
					fmt.Printf("SET redis err:%v\n", err)
				}
			}
			return nil

		})
	}


	if err := g.Wait(); err == nil {
		endMemory ,err := rdb.Info(ctx,"memory").Result()

		if err == nil {
			fmt.Printf("数据的个数 %f ；内存统计结束.结束时间：%s", maxNum, time.Now())
			fmt.Println(endMemory)
		}
	}




	// Output: PONG <nil>
}

func strCom(l int) (preKey string,keyVal string) {

	preKey = "testGoCamp:" + strconv.Itoa(l) +":"

	for i := 1; i <= l; i++ {
		keyVal += strBase
	}

	return preKey, keyVal
}


func main(){

	ExampleNewClient()
}
