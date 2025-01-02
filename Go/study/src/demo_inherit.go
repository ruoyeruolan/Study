package main

import "fmt"

type Human struct {
	name   string
	age    int
	gender string
}

func (ths *Human) Eat() {
	println("Human Eat()...")
}

func (ths *Human) Walk() {
	println("Human Walk()...")
}

type SuperMan struct {
	Human // SuperMan继承Human

	level int
}

// 重新定义弗雷方法
func (ths *SuperMan) Eat() {
	fmt.Println("SuperMan Eat()...")
}

// 重新定义子类的新方法
func (ths *SuperMan) Fly() {
	fmt.Println("SuperMan Fly()...")
}

func printSuperMan(ths *SuperMan) {
	fmt.Println("Name = ", ths.name)
	fmt.Println("Age = ", ths.age)
	fmt.Println("Gender=", ths.gender)
}

func main() {
	human := Human{"ryrl", 18, "male"}
	human.Eat()
	human.Walk()

	//定义一个子类对象
	s := SuperMan{Human{"ryrl", 18, "male"}, 88}

	/*
		var s SuperMan
		s.name = "ryrl"
		s.age = 18
		s.gender ="male"
		s.level = 88
	*/

	s.Walk() //调用父类方法
	s.Eat()  //调用子类方法
	s.Fly()  //调用子类方法
}
