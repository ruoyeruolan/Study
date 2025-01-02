package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reflectnum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}

type Usr struct {
	Id   int
	Name string
	Age  int
}

func (ths *Usr) Call() {
	fmt.Println("usr is called ...")
	fmt.Printf("%v\n", ths)
}

func DoFileAndMethod(input interface{}) {
	// 获取type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType)
	// 获取value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)
	//通过type获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field.Name, field.Type, value)
	}

	//通过type获取里面的方法, 调用
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

// 结构体标签
type resume struct {
	Name   string `info:"name" doc:"my name"`
	Gender string `info:"gender"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("info")
		fmt.Println(tag)

		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println(tagdoc)
	}
}

// 结构体标签在json中的应用
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	var num = 1.2345
	reflectnum(num)

	usr := Usr{1, "ryrl", 18}
	DoFileAndMethod(usr)

	var re resume
	findTag(&re)

	//结构体标签在json中的应用
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	//将结构体转换为json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}
	fmt.Printf("%s\n", jsonStr)

	//将json转换为结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
	//jsonStr2 := `{"title":"喜剧之王","year":2000,"price":10,"actors":["周星驰","张柏芝"]}`
}
