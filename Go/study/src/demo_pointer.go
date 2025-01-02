package main

import "fmt"

func changeVal(a int) {
	a = 10
}

func changeVal2(a *int) { // *表示是指针数据类型
	*a = 10
}

func swap(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func main() {
	var a = 1
	var b = 20
	changeVal(a)
	changeVal2(&a) // &表示取值地址
	fmt.Println("a=", a)

	fmt.Println("a=", a, "b=", b)
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)
}
