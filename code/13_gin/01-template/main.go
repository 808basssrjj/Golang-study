package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 解析模板文件
// func (t *Template) Parse(src string) (*Template, error)
// func ParseFiles(filenames ...string) (*Template, error)
// func ParseGlob(pattern string) (*Template, error)
// 模板渲染
// func (t *Template) Execute(wr io.Writer, data interface{}) error
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1.解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed err:", err)
		return
	}
	// 2.利用给定数据渲染模板，并将结果写入w
	// tmpl.Execute(w, "迪迦奥特曼") //1.传字符串
	u1 := UserInfo{
		Name:   "迪迦",
		Gender: "男",
		Age:    200000,
	}
	// err = tmpl.Execute(w, user)	//2.传结构体或map
	m1 := map[string]interface{}{
		"name":   "泰罗",
		"gender": "男",
		"age":    30000,
	}
	err = tmpl.Execute(w, map[string]interface{}{ //3.传多个
		"u1": u1,
		"m1": m1,
	})
	if err != nil {
		fmt.Println("render template failed", err)
		return
	}
}

func main() {
	// 启动服务
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("http start server failed err :", err)
		return
	}
}
