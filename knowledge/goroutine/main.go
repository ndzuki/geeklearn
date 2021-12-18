package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hellow world.")
	})
	//go func()是以后台运行的，当main函数退出时，go func()也会退出
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	select {} //空select将永远阻塞，使goroutine不会退出
}
