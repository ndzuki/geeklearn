package main

import "fmt"

func main() {
	ff := calc("+")
	r := ff(3, 4)
	fmt.Println("result:", r)
	ff = calc("-")
	r = ff(4, 2)
	fmt.Println("result:", r)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}
