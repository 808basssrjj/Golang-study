package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type hero struct {
	name string
	age  int
}
type heroSlice []hero

func (h heroSlice) Len() int {
	return len(h)
}
func (h heroSlice) Less(i, j int) bool {
	return h[i].age < h[j].age
}
func (h heroSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	var hs heroSlice
	for i := 0; i < 10; i++ {
		h := hero{
			name: fmt.Sprintf("hero%d", rand.Intn(100)),
			age:  rand.Intn(100),
		}
		hs = append(hs, h)
	}

	for _, v := range hs {
		fmt.Printf("name:%s age:%d\n", v.name, v.age)
	}

	// 实现Len(), Less(), Swap()即可方法 即可使用sort.Sort
	sort.Sort(hs)
	fmt.Println("排序后-----------------")
	for _, v := range hs {
		fmt.Printf("name:%s age:%d\n", v.name, v.age)
	}
}
