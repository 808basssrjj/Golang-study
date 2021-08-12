package family

import "fmt"

type Fm struct {
	key     string
	loop    bool
	balance float64
	// 每次收支金额
	money float64
	// 每次说明
	note string
	// 收支详情
	detail string
}

func NewFamily() *Fm {
	return &Fm{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		detail:  "收支\t账户金额\t收支金额\t说明\t",
	}
}

func (f *Fm) showDetails() {
	fmt.Println("---------当前收支明细记录----------")
	if len(f.note) != 0 {
		fmt.Println(f.detail)
	} else {
		fmt.Println("暂无明细")
	}
}
func (f *Fm) inCome() {
	fmt.Print("本次收入金额:")
	_, _ = fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Print("本次收入说明:")
	_, _ = fmt.Scanln(&f.note)
	f.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v\t", f.balance, f.money, f.note)
}
func (f *Fm) outCome() {
	fmt.Print("本次支出金额:")
	_, _ = fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("余额不足")
		//break
	}
	f.balance -= f.money

	fmt.Print("本次支出说明:")
	_, _ = fmt.Scanln(&f.note)
	f.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", f.balance, f.money, f.note)
}
func (f *Fm) exit() {
	choice := ""
	fmt.Print("确定要退出吗?<Y/N>:")

	for {
		_, _ = fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("请输入Y/N")
	}

	if choice == "y" {
		f.loop = false
	}
}
func (f Fm) MainMenu() {
	for {
		fmt.Println("\n----------家庭收支记账软件----------")
		fmt.Println("          1.收支明细")
		fmt.Println("          2.登记收入")
		fmt.Println("          3.登记支出")
		fmt.Println("          4.退出软件")
		fmt.Print("请选择<1-4>:")
		_, _ = fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			f.showDetails()
		case "2":
			f.inCome()
		case "3":
			f.outCome()
		case "4":
			f.exit()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !f.loop {
			fmt.Println("退出成功!")
			break
		}
	}
}
