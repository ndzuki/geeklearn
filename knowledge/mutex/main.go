package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		wt.Add(1)
		go add()
		wt.Add(1)
		go sub()
	}
	wt.Wait()
	fmt.Printf("m: %v\n", m)
}

var m int = 100
var lock sync.Mutex

var wt sync.WaitGroup

func add() {
	defer wt.Done()
	lock.Lock()
	m += 1
	fmt.Println("number: ", m)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	defer wt.Done()
	lock.Lock()
	m -= 1
	fmt.Println("number: ", m)
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}
