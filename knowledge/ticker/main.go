package main

import (
	"fmt"
	"time"
)

func main() {
	chanInt := make(chan int)
	//创建ticker，第秒执行一次
	ticker := time.NewTicker(time.Second)

	go func() {
		//遍历ticker.C，使chanInt每秒获取一个随机整数（1,2,3)
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			}
		}
	}()
	sum := 0
	for v := range chanInt {
		fmt.Println("recived: ", v)
		sum += v
		if sum >= 10 {
			fmt.Println("sum: ", sum)
			break
		}
	}
}
