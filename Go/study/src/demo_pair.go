package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var a string
	//pair<statictype: string, value: "ryrl">
	a = "ryrl"

	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)

	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)

	w.Write([]byte("Hello World!"))
}
