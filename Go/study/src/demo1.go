package main //  声明文件所在的包

import "fmt" //导入程序中的包，只有使用的包才要引入， 否则报错

func main() {
	fmt.Println("Hello world")

	//变量的声明,赋值，使用
	var age int = 10

	fmt.Println("age=", age)

	gender := "male"

	fmt.Println("gender=", gender) // 必须使用双引号， 单双引号不等价
}
