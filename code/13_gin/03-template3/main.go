package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 模板继承
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Println("paser failed err:", err)
		return
	}
	name := "法外狂徒张三"
	t.ExecuteTemplate(w, "index.tmpl", name)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 模板继承
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		fmt.Println("paser failed err:", err)
		return
	}
	name := "法外狂徒张三"
	t.ExecuteTemplate(w, "home.tmpl", name)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":9093", nil)
	if err != nil {
		fmt.Println("server start err", err)
		return
	}

}
