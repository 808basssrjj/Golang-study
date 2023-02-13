package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello Golang")

	s := []string{"你", "好", "呀"}
	str := strings.Join(s, "")
	fmt.Println(str)

	var b bytes.Buffer
	b.WriteString("1")
	b.WriteString("2")
	b.WriteString("3")
	str = b.String()
	fmt.Println(str)

	var b2 strings.Builder
	b2.WriteString("My")
	b2.WriteString("name")
	b2.WriteString("is")
	b2.WriteString("张三")
	str = b2.String()
	fmt.Println(str)
}
