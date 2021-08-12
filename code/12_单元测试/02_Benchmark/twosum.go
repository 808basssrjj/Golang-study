package twosum

// import "fmt"

func TwoSumfor(nums []int, target int) []int {
	for k1, v1 := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if v1+nums[k2] == target {
				return []int{k1, k2}
			}
		}
	}
	return nil
}

func TwoSummap(nums []int, target int) []int {
	m := make(map[int]int,len(nums))
	for k, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{index, k}
		}
		m[v] = k
	}
	return nil
}

// var nums =[]int{2, 7, 11, 15}
// var target = 26

// func main() {
// 	r1 := twoSumfor(nums, target)
// 	r2 := twoSummap(nums, target)
// 	fmt.Println(r1)
// 	fmt.Println(r2)
// }
