package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
)


func HandleRoot(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("success\n"))
}

func main() {
	var listenaddr = "localhost:8082"
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http.HandleFunc("/", HandleRoot)

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		select {
		case <- ctx.Done():
			return ctx.Err()
		default:
			return http.ListenAndServe(listenaddr, nil )
		}
	})

	s := <-c
	fmt.Printf("Got signal:%v\n", s)
	cancel()

	if err := g.Wait(); err != nil {
		log.Fatalf("%v\n", err)
	}
}
