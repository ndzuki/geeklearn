package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.NumCPU():", runtime.NumCPU())
	runtime.GOMAXPROCS(2) //设置最大使用线程数
	go print("Hello")
	go show("java")
	for i := 0; i < 2; i++ {
		//再次分配CPU时间
		runtime.Gosched() //注释查看前后结果
		fmt.Println("Go")
	}
}

func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

func print(s string) {
	for i := 0; i < 5; i++ {
		if i == 4 {
			runtime.Goexit() //退出协程
		} else {
			fmt.Printf("%v: %v\n", s, i)
		}
	}
}
