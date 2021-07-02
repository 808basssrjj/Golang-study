package split

import (
	"reflect"
	"testing"
)

// 一.测试函数
// 1.1测试
// func TestSplit(t *testing.T) { //测试函数名必须以Test开头，必须接收一个*testing.T类型参数
// 	// got := Split("abcbd", "b")      // 程序输出的结果
// 	// want := []string{"a", "c", "d"} // 期望的结果
// 	got := Split("我爱你", "爱")      // 程序输出的结果
// 	want := []string{"我", "你"} // 期望的结果
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
// 	}
// }

// 1.2测试组
// func TestSplit(t *testing.T) { //测试函数名必须以Test开头，必须接收一个*testing.T类型参数
// 	type test struct {
// 		input string
// 		sep   string
// 		want  []string
// 	}
// 	tests := []test{
// 		test{input: "我爱你", sep: "爱", want: []string{"我", "你"}},
// 		test{input: "abcbd", sep: "b", want: []string{"a", "c", "d"}},
// 		test{input: "沙河有沙又有河", sep: "沙", want: []string{"","河有", "又有河"}},
// 	}
// 	// 遍历切片，逐一执行测试用例
// 	for _, ts := range tests {
// 		got := Split(ts.input, ts.sep)
// 		if !reflect.DeepEqual(got, ts.want) {
// 			t.Errorf("excepted:%#v, got:%#v", ts.want, got) // 测试失败输出错误提示
// 		}
// 	}

// }

// 1.3子测试
func TestSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
	// 我们都知道可以通过-run=RegExp来指定运行的测试用例，还可以通过/来指定要运行的子测试用例，
	// 例如：go test -v -run=Split/simple只会运行simple对应的子测试用例。
}

// 1.4测试覆盖率   测试覆盖率是你的代码被测试套件覆盖的百分比。
// 可以使用go test -cover来查看测试覆盖率
// 提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件
// 例: go test -cover -coverprofile=c.out
// 使用 go tool cover -html=c.out 打开本地的浏览器窗口生成一个HTML报告

// 二.基准测试
func BenchmarkSplit(b *testing.B) {
	// b.n不是固定的数
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
//go test -bench=Split
// 还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据
