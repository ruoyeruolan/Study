package main

import "fmt"

// 函数
// 函数的定义
// func 函数名(参数列表) (返回值列表) {}
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// 多个返回值，但都是匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 777
}

// 多个返回值，有形参名
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// 给返回值赋值， r1,r2属于foo3的形参，如果不赋值，就是默认值0
	r1 = 666
	r2 = 777
	return
}

func foo4(a string, b int) (r1, r2 int) { //返回值的类型相同，可以简写
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	// 给返回值赋值
	r1 = 666
	r2 = 777
	return
}
func main() {
	// 函数的调用
	c := foo1("abc", 100)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("def", 200)
	fmt.Println("res1 = ", ret1)
	fmt.Println("res2 = ", ret2)

	ret3, ret4 := foo3("ghi", 300)
	fmt.Println("res3 = ", ret3, '\n', "res4 = ", ret4)

	ret5, ret6 := foo4("jkl", 400)
	fmt.Println("res5 = ", ret5, '\n', "res6 = ", ret6)
}
