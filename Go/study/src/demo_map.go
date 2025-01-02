package main

import "fmt"

func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key, "value=", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}
func main() {
	// map声明, key为string类型, value为int类型
	var map1 map[string]int

	if map1 == nil {
		fmt.Println("map1 is enpty")
	}

	// 使用map前需要先make给map分配空间
	map2 := make(map[string]string, 10)
	map2["one"] = "Python"
	map2["two"] = "Java"
	map2["three"] = "Go"
	map2["four"] = "C++"

	fmt.Println(map2) // 没有顺序，是一个hash表

	map3 := make(map[string]string)
	map3["one"] = "Python"
	map3["two"] = "Java"
	map3["three"] = "Go"
	map3["four"] = "C++"
	fmt.Println(map3)

	map4 := map[string]string{
		"one":   "Python",
		"two":   "Java",
		"three": "Go",
		"four":  "C++",
	}
	fmt.Println(map4)

	cityMap := make(map[string]string)

	// 添加元素
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "Washington"
	cityMap["France"] = "Paris"
	cityMap["Italy"] = "Rome"

	//遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, "value=", value)
	}
	printMap(cityMap) //将map作为参数传递给函数
	// 删除元素
	delete(cityMap, "China")
	// 修改
	cityMap["USA"] = "NewYork"
	ChangeValue(cityMap)

}
