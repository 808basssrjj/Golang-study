package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 修改默认标识符
	t, err := template.New("hello.tmpl").Delims("{[", "]}").ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("paser failed err:", err)
		return
	}
	name := "法外狂徒张三"
	_ = t.Execute(w, name)
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9093", nil)
	if err != nil {
		fmt.Println("server start err", err)
		return
	}

}
