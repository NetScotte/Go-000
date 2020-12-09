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
	log.Printf("receive request from host: %v\n", request.RemoteAddr)
	response.Write([]byte("success\n"))
}

// 通过控制errgroup收集错误信息，实现进程的统一管理
func main() {
	var listenaddr = "localhost:8082"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http.HandleFunc("/", HandleRoot)
	httpServer := &http.Server{
		Addr:              listenaddr,
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	g, _ := errgroup.WithContext(ctx)
	// 监听signal，该进程在收到signal或者ctx关闭后，会退出
	g.Go(func() error {
		select {
		case s := <-signalChannel:
			cancel()
			return fmt.Errorf("Got signal: %v, cancel ctx\n", s)
		case <-ctx.Done():
			return fmt.Errorf("signal worker kill by ctx")

		}
	})

	// 启动http服务
	g.Go(func() error {
		err := httpServer.ListenAndServe()
		if err != nil {
			cancel()
		}
		return err
	})

	// 控制http server
	g.Go(func() error {
		select {
		case <-ctx.Done():
			return httpServer.Shutdown(ctx)
		}
	})


	if err := g.Wait(); err != nil {
		log.Printf("%v\n", err)
	}
}
