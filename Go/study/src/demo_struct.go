package main

import "fmt"

// Go语言的结构体

// 声明一种新的数据类型
type myInt int

// 定义一个结构体
type Book struct {
	title  string
	author string
} // Book具有两个属性: title和author

func printBook(book Book) {
	fmt.Println("book = ", book)
}

func ChangeBook(book Book) {
	book.title = "Python"
}

func ChangeBook2(book *Book) {
	book.title = "Python"
}

// 类的名字首字母大写，表示public，其他包也能调用
type Hero struct {
	name string
	age  int
}

func (ths *Hero) Show() {
	fmt.Println("hello, my name is ", ths.name)
}
func (ths *Hero) GetName() string {
	//fmt.Println(ths.name)
	return ths.name
}

func (ths *Hero) SetName(name string) {
	ths.name = name
}

func main() {
	var a myInt = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var book Book
	book.title = "Golang"
	book.author = "ryrl"

	fmt.Println("book = ", book)
	printBook(book)
	ChangeBook(book) // 并没有改变book的值，传入的是副本，不会修改参数原来的值
	printBook(book)

	ChangeBook2(&book) // 传入的是指针，会修改参数原来的值
	printBook(book)

	//类
	hero := Hero{
		name: "ryrl",
		age:  18,
	}

	hero.Show()
	hero.SetName("Wakala")
	hero.Show()
}
