package main

import "fmt"

// 一个数组中大部分数据的元素为0(无意义数据)   用来压缩数据
// 1.记录数组一共有几行几列，有多少不同的值
// 2.吧具有不同值的元素的行列机值记录再一个小规模的数组中

type node struct {
	row int //行
	col int //列
	val int //值
}

func main() {

	// 1.创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// 2.存入稀疏数组
	var sparseArr []node
	node0 := node{11, 11, 0}
	sparseArr = append(sparseArr, node0)
	// 如果值部位0就存下来
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := node{i, j, v2}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	// 输出稀疏数组
	for i, v := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, v.row, v.col, v.val)
	}

	// 3/恢复成原来数组  打开存盘的文件恢复
	var chessMap2 [11][11]int
	for i, v := range sparseArr {
		if i == 0 {
			continue
		} else {
			chessMap2[v.row][v.col] = v.val
		}
	}
	fmt.Println("恢复后的数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
