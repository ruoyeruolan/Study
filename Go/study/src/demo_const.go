package main

import "fmt"

// const定义枚举类型

const (
	// 可以在const()添加一个关键字iota,每行的iota都会累加1,第一行的iota默认是0, iota只能出现在const括号中
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	// 常量,只读属性,不能修改
	const length int = 10

	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("g = ", g)
	fmt.Println("h = ", h)
	fmt.Println("i = ", i)
	fmt.Println("k = ", k)
}

/*

 */
