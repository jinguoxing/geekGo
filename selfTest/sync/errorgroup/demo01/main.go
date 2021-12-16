package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main(){

	g := new(errgroup.Group)

   var urls  = []string{

   	"https://www.baidu.com",
   	"https://www.google.com",
   }

  for _ ,url := range urls {

  	url := url

  	g.Go(func() error {

  		resp ,err := http.Get(url)

  		if err == nil {
  			resp.Body.Close()
		}
		return err
  	})
  }

	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}


}