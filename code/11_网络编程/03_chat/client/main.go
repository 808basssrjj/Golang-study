package main

import "fmt"

var (
	userId  int
	userPwd string
)

func main() {
	// 接受用户选择
	var key int
	// 是否循环
	var loop bool

	// 菜单
	for {
		fmt.Println("----------欢迎登录多人聊天系统----------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择(1-3)")

		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("输入错误,请重新选择!")
		}

		if !loop {
			break
		}
	}

	if key == 1 {
		// 1.用户登录
		fmt.Println("请输入用户id")
		_, _ = fmt.Scanln(&userId)
		fmt.Println("请输入用户密码")
		_, _ = fmt.Scanln(&userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("login failed error:", err)
		} else {
			fmt.Println("登录成功!")
		}
	} else if key == 2 {
		// 2.用户注册
	}
}
