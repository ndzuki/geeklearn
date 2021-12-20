// go run -gcflags="-m -l" escapeAnalysis-3.go
// # command-line-arguments
// ./escapeAnalysis-3.go:10:2: moved to heap: b
// ./escapeAnalysis-3.go:13:2: moved to heap: d
// ./escapeAnalysis-3.go:16:2: moved to heap: f
// ./escapeAnalysis-3.go:9:11: make([]*int, 1) does not escape
// ./escapeAnalysis-3.go:12:11: make(map[string]*int) does not escape
// 在这里，变量b,d,f的内存都被移动到堆上，因为，Golang中，slice,map,channel引用指针的变量，一定会逃逸。 Golang中，slice，map，channel对指针的引用会比之保留变量的slice，map，channel性能低，这里是根本原因。
package main

func main() {
	a := make([]*int, 1)
	b := 12
	a[0] = &b
	c := make(map[string]*int)
	d := 14
	c["aaa"] = &d
	e := make(chan *int, 1)
	f := 15
	e <- &f
}
