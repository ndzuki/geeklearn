package main

import (
	"fmt"
	"strings"
)

// return functionlity
func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	var f = add()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println("==Class 1==")
	f1 := add()
	fmt.Println(f1(40))
	fmt.Println(f1(50))
	fmt.Println("==Class 2==")
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
	fmt.Println("==Class 3==")
	f2, f3 := calc(10)
	fmt.Println(f2(1), f3(2))
	fmt.Println(f2(3), f3(4))
	fmt.Println(f2(5), f3(6))
}
