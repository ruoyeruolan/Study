package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1

	for {
		select { // 多路channel监控状态
		case c <- x: //如果c可写，则执行case
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")

			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 从通道c中接收数据，并打印到终端
		}

		quit <- 0
	}()

	fibonacci(c, quit)
}
