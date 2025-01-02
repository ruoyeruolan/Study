package lib2

import "fmt"

// lib2提供的API
func DemoLib2() { //函数名首字母大写，表示该函数可以被外部包访问，小写的只能在当前包内使用
	fmt.Println("DemoLib2() runing...")
}

func init() {
	fmt.Println("lib2.init() runing...")
}
