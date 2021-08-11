package main

import "fmt"

// 中间下标：middle = (left + right) / 2
// 1. 如果 arr[middle] = findVal
// 2. 如果 arr[middle] > findVal 查找范围 left-=middle-1
// 3. 如果 arr[middle] < findVal 查找范围 middle+1--right
// 4. 如果left > right 没找到
func binarySearch(arr *[6]int, left int, right int, findVal int) {
	if left > right {
		fmt.Println("没找到")
		return
	}
	middle := (left + right) / 2
	if (*arr)[middle] == findVal {
		fmt.Printf("找到了 下标为%d", middle)
		return
	} else if (*arr)[middle] > findVal {
		binarySearch(arr, left, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		binarySearch(arr, middle+1, right, findVal)
	}
}

func main() {
	// 只能对有序数组
	arr := [6]int{1, 2, 3, 6, 8, 9}
	binarySearch(&arr, 0, len(arr)-1, 9)
}
