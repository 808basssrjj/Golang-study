package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello pig"
	err := ioutil.WriteFile("./ioutil.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
