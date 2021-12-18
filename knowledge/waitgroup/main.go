package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //当启动一个goroutine时，+1
		go hello(i)
	}
	wg.Wait() //等待所有登记的goroutine结束
	fmt.Println("Processed end.")
}

func hello(i int) {
	defer wg.Done() //goroutine结束时登记-1
	fmt.Println("Hello Goroutine!", i)
}
