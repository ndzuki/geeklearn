// go build -gcflags="-m" escapeAnalysis-5.go
// # command-line-arguments
// ./escapeAnalysis-5.go:17:6: can inline main
// ./escapeAnalysis-5.go:12:11: make([]int, 1000, 1000) does not escape
// 这里，没有产生内存逃逸，如果将slice的长度增长10倍,就会产生逃逸。
// ./escapeAnalysis-5.go:14:6: can inline main
// ./escapeAnalysis-5.go:9:11: make([]int, 10000, 10000) escapes to heap
package bigslice

func BigSlice() {
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}
func main() {
	BigSlice()
}
