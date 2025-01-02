package main

import "fmt"

func deferFunc() int {
	fmt.Println("deferFunc called ...")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc called ...")
	return 0
}

func returnAnDefer() int {
	defer deferFunc()   // deferFunc() 会在returnFunc()之后执行
	return returnFunc() // returnFunc() 会在deferFunc()之前执行
}
func main() {
	defer fmt.Println("deferred call in main 1") // 在main函数结束时，执行defer
	defer fmt.Println("deferred call in main 2") // 两个defer，压栈，先进后出，先写的后执行

	fmt.Println("main:: Hello, World! 1")
	fmt.Println("main:: Hello, World! 2")

	returnAnDefer() //先返回return的数据，再执行defer；写代码的时候尽量避免这种写法
}
