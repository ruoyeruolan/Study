package main

// 使用go mod init 「name」初始化一个module，导入自己的包时，从module的名字作为的根目录开始导入
import (
	"fmt"
	_ "src/lib1"      // 匿名调用，只执行init函数，不调用其他函数，不使用它不会报错，不匿名调用的话，不使用函数会报错
	mylib2 "src/lib2" //使用别名，避免和lib2包冲突;前面使用"."的话，调用lib2包的函数时，可以省略lib2，直接调用函数名称
)

func main() {
	//lib1.DemoLib1()
	//lib2.DemoLib2()
	mylib2.DemoLib2()
	fmt.Println("Hello, World!")
}
