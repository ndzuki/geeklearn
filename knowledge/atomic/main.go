package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int32 = 100
	atomic.AddInt32(&i, 1) //加1
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -1) //减1
	fmt.Printf("i: %v\n", i)

	atomic.LoadInt32(&i) //加载
	fmt.Printf("i: %v\n", i)

	atomic.StoreInt32(&i, 200) //存储，写入
	fmt.Printf("i: %v\n", i)

	b := atomic.CompareAndSwapInt32(&i, 200, 300) //比较新旧值并交换，返回bool值
	fmt.Printf("i: %v\n", b)
	fmt.Printf("i: %v\n", i)
}
