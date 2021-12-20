//go build -gcflags="-m" escapeAnalysis-1.go
//# command-line-arguments
// ./escapeAnalysis-1.go:16:6: can inline main
// ./escapeAnalysis-1.go:9:13: leaking param: name
// ./escapeAnalysis-1.go:10:10: new(Cat) escapes to heap
// 可以看到,./cat.go:10:10: new(Cat) escapes to heap， 有变量的内存逃逸。
package main

type Cat struct {
	Name string
	Age  int
}

//go:noinline
func NewCat(name string, age int) *Cat {
	c := new(Cat) // c will excape to heap
	c.Name = name
	c.Age = age
	return c
}

func main() {
	NewCat("Tom", 5)
}
