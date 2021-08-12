package twosum

import (
	"math/rand"
	"testing"
)

const N = 1000


// go test -bench=.

func BenchmarkTwoSumfor(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 11, 15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSumfor(nums, 26)
	}
}


func BenchmarkTwoSummap(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 11, 15)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSummap(nums, 26)
	}
}
