package main

import "shouzhang/family"

func main() {
	f := family.NewFamily()
	f.MainMenu()
}

//func main() {
//	var key string
//	loop := true
//	balance := 10000.0
//	// 每次收支金额
//	money := 0.0
//	// 每次说明
//	note := ""
//	// 收支详情
//	detail := "收支\t账户金额\t收支金额\t说明\t"
//	for {
//		fmt.Println("\n----------家庭收支记账软件----------")
//		fmt.Println("          1.收支明细")
//		fmt.Println("          2.登记收入")
//		fmt.Println("          3.登记支出")
//		fmt.Println("          4.退出软件")
//		fmt.Print("请选择<1-4>:")
//		_, _ = fmt.Scanln(&key)
//
//		switch key {
//		case "1":
//			fmt.Println("---------当前收支明细记录----------")
//			if len(note) != 0 {
//				fmt.Println(detail)
//			} else {
//				fmt.Println("暂无明细")
//			}
//		case "2":
//			fmt.Print("本次收入金额:")
//			_, _ = fmt.Scanln(&money)
//			balance += money
//			fmt.Print("本次收入说明:")
//			_, _ = fmt.Scanln(&note)
//			detail += fmt.Sprintf("\n收入\t%v\t%v\t%v\t", balance, money, note)
//		case "3":
//			fmt.Print("本次支出金额:")
//			_, _ = fmt.Scanln(&money)
//			if money > balance {
//				fmt.Println("余额不足")
//				break
//			}
//			balance -= money
//
//			fmt.Print("本次支出说明:")
//			_, _ = fmt.Scanln(&note)
//			detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
//		case "4":
//			choice := ""
//			fmt.Print("确定要退出吗?<Y/N>:")
//
//			for {
//				_, _ = fmt.Scanln(&choice)
//				if choice == "y" || choice == "n" {
//					break
//				}
//				fmt.Print("请输入Y/N")
//			}
//
//			if choice == "y" {
//				loop = false
//			}
//		default:
//			fmt.Println("请输入正确的选项")
//		}
//
//		if !loop {
//			fmt.Println("退出成功!")
//			break
//		}
//	}
//}
