// struct {}
// struct {} 是一个无元素的结构体类型，通常在没有信息存储时使用。
// 优点：不需要内存来存储 struct {} 类型的值。
// struct{}{}
// struct {}{} 是一个复合字面量，它构造了一个 struct {} 类型的值，该值也是空。
// 两个 structt {}{} 地址相等

// ————————————————
// 原文作者：ljq
// 转自链接：https://learnku.com/articles/59899
// 版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。

package main

import "fmt"

type idBval struct {
	Id int
}

func main() {
	idA := struct{}{}
	fmt.Printf("idA: %T and %v \n\n", idA, idA)

	idB := idBval{
		1,
	}
	idB.Id = 2
	fmt.Printf("idB: %T and %v \n\n", idB, idB)

	idC := struct {
		Id int
	}{
		1,
	}
	fmt.Printf("idC: %T and %v \n\n", idC, idC)

	mapD := make(map[string]struct{})
	mapD["mapD"] = struct{}{}
	_, ok := mapD["mapD"]
	fmt.Printf("mapD['mapD'] is %v \n\n", ok)

	sliceE := make([]interface{}, 2)
	sliceE[0] = 1
	sliceE[1] = struct{}{}
	fmt.Printf("idE: %T and %v \n\n", sliceE, sliceE)

}
