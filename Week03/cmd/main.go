package main

import (
	"geekGo/Week03"
	"geekGo/Week03/transport/http"
	"log"
	"time"
)

func main(){


	hs := http.NewServer()

	app := Week03.New(
		Week03.Name("appTest"),
		Week03.Version("v1.0.0"),
		Week03.Server(hs),
	)
	time.AfterFunc(10*time.Second, func() {
		app.Stop()
	})
	if err := app.Run(); err != nil {
	  log.Fatal(err)
	}

}
