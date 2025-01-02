package main

import (
	"fmt"
)

// 数组

func main() {
	var myarray [10]int // 固定长度的数组
	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8} // 动态数组

	for i := 0; i < len(myarray); i++ { // 循环遍历数组
		fmt.Println(myarray[i])
	}

	for index, value := range myarray {
		fmt.Println("index=", index, "value=", value)
	}

	for index, val := range myArray2 {
		fmt.Println("index=", index, "val=", val)
	}

	fmt.Println("myArray types = %T\n", myarray)
	fmt.Println("myArray2 types = %T\n", myArray2)
	fmt.Println("myArray3 types = %T\n", myArray3)

	//声明一个切片并初始化
	slice1 := []int{1, 2, 3}

	//声明一个切片，但是并没有分配空间
	var slice2 []int
	slice2 = make([]int, 3) // 分配空间
	fmt.Printf("length = %d, slice = %v\n", len(slice1), slice2)

	//声明一个切片，同时指定了长度和容量
	slice3 := make([]int, 3, 5)
	fmt.Printf("length = %d, slice = %v\n", len(slice3), slice3)

	//切片的截取

	number := make([]int, 3, 5)                                                      // 声明一个长度为3，容量为5的切片
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number) // cap()函数用于计算切片的容量

	//向numbers切片追加一个元素1，由于超出了numbers的长度，因此这个时候numbers的长度会发生变化
	number = append(number, 1) // TODO: 如何添加多个元素
	number = append(number, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	//向numbers切片追加一个元素1，由于超出了numbers的容量，会开辟一个与之前容量相同的新的底层数组，然后将原来的值拷贝过去
	number = append(number, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(number), cap(number), number)

	// 切片的截取, 切片截取使用的是浅拷贝，即引用，底层数组相同，改变一个切片的值，另一个切片的值也会改变
	// copy可以将底层数组的值拷贝到另一个数组中，这样就不会影响到另一个切片的值
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(s[0:4]) // [1 2 3 4]与python的切片类似，左闭右开
}
