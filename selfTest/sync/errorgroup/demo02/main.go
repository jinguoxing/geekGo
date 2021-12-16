package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func main(){

	g, ctx := errgroup.WithContext(context.Background())

	var urls  = []string{

		"https://www.baidu.com",
		"https://www.google.com",
	}
	fmt.Println(time.Now())
	for _ ,url := range urls {

		url := url

		g.Go(func() error {

			resp ,err := http.Get(url)

			if err == nil {
				resp.Body.Close()
			}

			timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			canceled := false

			select {
			case <- ctx.Done():
				fmt.Println("errgroup cancel ----",ctx.Err(),time.Now())
				canceled = true
			case <- timeout.Done():
				fmt.Println("errgroup timeout cancel ----",timeout.Err(),time.Now())
				canceled = true
			}
			defer cancel()

			if canceled {
				fmt.Println("ctx cancel ----",ctx.Err(),time.Now())
				ctx.Done()
				return ctx.Err()
			}
			return err
		})
	}


	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}else{
		fmt.Println("failedfetched all URLs.",err)
	}



}

