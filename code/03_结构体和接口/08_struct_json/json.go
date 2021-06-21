package main

import (
	"encoding/json"
	"fmt"
)

type persion struct {
	// 字段首字母大写才能外部访问
	// 通过指定tag实现json序列化该字段时的key
	Name string `json:"name" db:"name"` // json,db解析时为name
	Age  int    `json:"age"`
}

func main() {
	// 1.JSON序列化：结构体-->JSON格式的字符串
	p1 := &persion{
		Name: "张三",
		Age:  18,
	}
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err: %v", err)
		return
	}
	fmt.Printf("json:%s\n", data)

	// 2.JSON反序列化：JSON格式的字符串-->结构体
	str := `{"name":"张三","age":18}`
	var p2 persion
	err = json.Unmarshal([]byte(str), &p2) //传指针是为了能在json.Unmarshal()中改变p2
	// p3 := &persion{}
	// json.Unmarshal([]byte(str), p3)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Println(p2)
}
