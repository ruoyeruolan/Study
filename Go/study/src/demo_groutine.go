package main

import (
	"fmt"
	"runtime"
	"time"
)

// 子Groutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Groutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主Groutine
func main() {
	//创建一个go程，去执行newTask()函数
	//go newTask()
	//
	//i := 0
	//
	//for {
	//	i++
	//	fmt.Printf("main Groutine : i = %d\n", i)
	//	time.Sleep(1 * time.Second)
	//}

	//用go创建一个行参为空，返回为空的函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//终止所在的协程
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
