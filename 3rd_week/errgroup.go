// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

// 以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中： https://jinshuju.net/f/LMjE3K
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))

	})

	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}

	})
	server := http.Server{
		Handler: mux,
		Addr:    ":8090",
	}

	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("g1 error, will exit.", err.Error())
		}
		return err
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("g2 errgroup exit...")
		case <-serverOut:
			log.Println("g2, request `/shutdown`, server will out...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		defer cancel()

		err := server.Shutdown(timeoutCtx)
		log.Println("shutdown server...")
		return err
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("g3, ctx execute cancel...")
			log.Println("g3 error", ctx.Err().Error())
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("g3 get os signal: %v", sig)
		}
	})

	fmt.Printf("end, errgroup exiting, %+v\n", g.Wait())
}
