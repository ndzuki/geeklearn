// go build -gcflags="-m -l" escapeAnalysis-2.go
// # command-line-arguments
// ./escapeAnalysis-2.go:11:2: moved to heap: s
// ./escapeAnalysis-2.go:16:13: ... argument does not escape
// ./escapeAnalysis-2.go:16:14: *x escapes to heap
// 在这里，变量b,d,f的内存都被移动到堆上，因为，Golang中，slice,map,channel引用指针的变量，一定会逃逸。 Golang中，slice，map，channel对指针的引用会比之保留变量的slice，map，channel性能低，这里是根本原因。
package main

import "fmt"

func test() *int {
	s := 3
	return &s
}
func main() {
	x := test()
	fmt.Println(*x)
}
