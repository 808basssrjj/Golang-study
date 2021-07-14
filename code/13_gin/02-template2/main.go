package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 自定义函数
	shuai := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	// t, err := template.New("f").Parse("./hello.tmpl")
	t, err := template.New("hello.tmpl").Funcs(template.FuncMap{"shuai": shuai}).ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("tmplate failed err :%v\n", err)
		return
	}

	name := "特利迦"
	t.Execute(w, name)
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/tmplDemo", tmplDemo)
	err := http.ListenAndServe(":9092", nil)
	if err != nil {
		fmt.Println("server start failed err:", err)
		return
	}
}
