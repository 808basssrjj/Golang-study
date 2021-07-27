package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)




// 2.子测试
func TestDoubleHandler(t *testing.T) {
	// go test -v

	testCases := []struct {
		name   string
		input  string
		result int
		status int
		err    string
	}{
		{name: "double of two", input: "2", result: 4, status: http.StatusOK, err: ""},
		{name: "double of nine", input: "9", result: 18, status: http.StatusOK, err: ""},
		{name: "double of nil", input: "", status: http.StatusBadRequest, err: "misssing value"},
	}
	for _, testCase := range testCases {
		// 子测试
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "localhost:4000/double?v="+testCase.input, nil)
			if err != nil {
				t.Fatalf("could not create a new request: %v, err: %v", req, err)
			}

			rec := httptest.NewRecorder() //创建一个用来测试的 ResponseWriter
			doubleHandler(rec, req)
			res := rec.Result() //获得响应结果
			if res.StatusCode != testCase.status {
				t.Errorf("received status code %s, expect %d", res.Status, testCase.status) // 测试失败返回错误提示
				return
			}

			respBytes, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("cannot read all from the response body err: %v", err)
			}
			defer res.Body.Close()

			trimResult := strings.TrimSpace(string(respBytes))

			if res.StatusCode != http.StatusOK {
				// check th error message
				if trimResult != testCase.err {
					t.Errorf("received error message  %s, expect %s", trimResult, testCase.err)
				}
				return
			}

			doubleResult, err := strconv.Atoi(trimResult)
			if err != nil {
				t.Errorf("cannot convert body to int err: %v", err)
				return
			}

			if doubleResult != testCase.result {
				t.Errorf("received result %d, expect %d", doubleResult, testCase.result)
			}
		})
	}
}

// 1.单元测试
//func TestDoubleHandler(t *testing.T) {
//	req, err := http.NewRequest(http.MethodGet, "localhost:4000/double?v=4", nil)
//	if err != nil {
//		t.Fatalf("could not create a new request: %v, err: %v", req, err)
//	}
//
//	rec := httptest.NewRecorder() //创建一个用来测试的 ResponseWriter
//	doubleHandler(rec, req)
//	res := rec.Result() //获得响应结果
//	if res.StatusCode != http.StatusOK {
//		t.Errorf("recevied status code %s, expect %d", res.Status, http.StatusOK) // 测试失败返回错误提示
//		return
//	}
//
//	respBytes, err := io.ReadAll(res.Body)
//	if err != nil {
//		t.Fatalf("cannot read all from the response body err: %v", err)
//	}
//
//	result, err := strconv.Atoi(strings.TrimSpace(string(respBytes)))
//	if err != nil {
//		t.Errorf("cannot convert body to int err: %v", err)
//		return
//	}
//
//	if result != 8 {
//		t.Errorf("recevied %d, expect 8", result)
//	}
//}
