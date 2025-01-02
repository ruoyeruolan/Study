package main

import "fmt"

type Animal interface { // 本质是一个指针
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  // 获取动物的类型
}

// 具体的类
// 定义完接口后，需要实现接口的方法，每日一个借口都需要实现
type Cat struct {
	color string
}

func (ths *Cat) Sleep() {
	println("Cat is sleep")
}

func (ths *Cat) GetColor() string {
	return ths.color
}

func (ths *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (ths *Dog) Sleep() {
	println("Dog is sleep")
}

func (ths *Dog) GetColor() string {
	return ths.color
}

func (ths *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal Animal) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("kind=", animal.GetType())
}

// interface{}  空接口，没有任何方法，所以所有的类型都实现了空接口，所以可以把任何一个变量赋值给空接口
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//interface{}如何识别具体的底层数据类型 ==》类型断言
	//如果arg是string类型，就返回string类型的数据
	value, ok := arg.(string) //断言会返回两个结果
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value=", value)
	}
}

type book struct {
	auth string
}

func main() {
	/*
		var animal Animal  //接口的数据类型，父类指针
		animal = &Cat{"green"}
		animal.Sleep()
		animal = &Dog{"yellow"}
		animal.Sleep()
	*/
	cat := Cat{"green"}
	dog := Dog{"yellow"}

	showAnimal(&cat)
	showAnimal(&dog)

	book := book{"Python"}
	myFunc(book)
	myFunc(100)
}
