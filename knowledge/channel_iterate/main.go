package main

import (
	"fmt"
)

func main() {
	var c = make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//遍历channel，通过err状态判断是否可读取值
	for {
		v, ok := <-c
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
