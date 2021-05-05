package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"log"
	"time"
	"syscall"
	pkgerrors "github.com/pkg/errors"
	"fmt"
)

func main(){


	g ,ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {

		writer.Write([]byte("PONG"))
		
	})

	serverOut := make(chan struct{})

	mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {

		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr: ":8888",
	}

	//协程1
	g.Go(func() error{
		return  server.ListenAndServe()
	})


	// 协程2
	g.Go(func() error {

		select {
		case <-ctx.Done():
			log.Println("errGroup exit...")
		case <-serverOut:
			log.Println("Server will out ...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		defer cancel()

		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx)

	})


	// 协程3

	g.Go(func() error{

		quit := make(chan os.Signal,1)
		signal.Notify(quit,syscall.SIGTERM,syscall.SIGQUIT,syscall.SIGINT)

		select {

		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:

			return pkgerrors.Errorf("Get OS Signal :%v",sig)

		}

	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())

}
