package lib1

import "fmt"

// lib1提供的API
func DemoLib1() {
	fmt.Println("DemoLib1() runing...")
}

func init() {
	fmt.Println("lib1.init() runing...")
}
