package main

import (
	"fmt"
	"time"
)

// channel通信，获取返回值
func main() {
	//定义一个channel, 无缓冲
	c := make(chan int)
	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("go go go!!!")

		c <- 666
	}()

	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("main goroutine ended")

	//
	go func() {
		defer fmt.Println("goroutine2结束")

		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("goroutine2 i = ", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
		close(c) // 关闭channel
	}()

	//for {
	//	// ok如果为true，表示channel没有关闭，如果为false，表示channel已经关闭
	//	if num, ok := <-c; ok == true {
	//		fmt.Println("num = ", num)
	//	} else {
	//		break
	//	}
	//}

	for data := range c {
		fmt.Println("data = ", data)
	} //for range方式，遍历channel，直到channel关闭, 等价与上面的for循环`

	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		num := <-c //从c中接收数据，并赋值给num, 不能有空格， num := <- c
		fmt.Println("goroutine2 num = ", num)
	}

	fmt.Println("main goroutine ended2")

	//关闭channel

	//select

}
