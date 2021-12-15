package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 3)
	go func() {
		<-timer1.C //调用定时器timer1，定义的是3秒时间，C 是channel的意思。倒数3秒后取出timer1 channel的值
		fmt.Println("Timer 1 expired")
	}()
	stop := timer1.Stop() //停止定时器
	if stop {
		fmt.Println("Timer 1 stopped")
	}

	fmt.Println("before")
	timer2 := time.NewTimer(time.Second * 4)
	timer2.Reset(time.Second * 1) //重置timer2的等待时间
	<-timer2.C
	fmt.Println("after")
	fmt.Println("Now:", time.Now())
}
