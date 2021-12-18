package main

import "fmt"

func f1() {
	a := 0
	if a == 1 {
		goto LABEL1 //跳到指定标签
	} else {
		fmt.Println("other")
	}
LABEL1:
	fmt.Printf("next...")

}

func f2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if 1 == 2 && j == 2 {
				goto LABEL2 //跳出嵌套循环
			}
		}
	}
LABEL2:
	fmt.Println("label2")
}
func main() {
	f1()
	f2()
}
