package main

import (
	"fmt"
	"reflect"
)

func main() {
	floatVar := 3.14
	v := reflect.ValueOf(floatVar)
	newFloat := v.Interface().(float64)
	fmt.Println(newFloat + 1.2)

	sliceVar := make([]int, 5)
	v = reflect.ValueOf(sliceVar)
	v = reflect.Append(v, reflect.ValueOf(4))
	newSlice := v.Interface().([]int)
	newSlice = append(newSlice, 5)
	fmt.Println(newSlice)
}
