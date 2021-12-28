// go build -gcflags="-m" escapeAnalysis-4.go
// # command-line-arguments
// ./escapeAnalysis-4.go:3:6: can inline Fibonacci
// ./escapeAnalysis-4.go:5:9: can inline Fibonacci.func1
// ./escapeAnalysis-4.go:4:2: moved to heap: a
// ./escapeAnalysis-4.go:4:5: moved to heap: b
// ./escapeAnalysis-4.go:5:9: func literal escapes to heap
// 在Fibonacci()函数中，a,b是一个本地的变量，因为被闭包引用，所以被分配在了堆上。
package fib

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
