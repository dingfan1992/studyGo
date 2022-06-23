package main

import (
	"fmt"
	"studyGo/Basics/user"
)

func main() {
	num := 3
	bank := user.Bank{BankPwd: "666666", BankCode: "88888888", Balance: 100}
	flag := true
	for num > 0 {
		// 获取键盘输入
		var bankCode string
		fmt.Print("请输入银行卡号：")
		_, _ = fmt.Scanln(&bankCode)

		var bankPwd string
		fmt.Print("请输入密码：")
		_, _ = fmt.Scanln(&bankPwd)
		isok, message := bank.Verification(bankCode, bankPwd)
		if !isok {
			num -= 1
			if num == 0 {
				fmt.Printf("您的账号密码错误3次，已被锁定！！")
				break
			}
			fmt.Printf("%s您还有%d次机会\n", message, num)
			continue
		} else {
			var operation int64
			for flag {
				fmt.Print("请选择您的操作:\n1,余额查询\n2,存款\n3.取款\n4,退出\n")
				_, _ = fmt.Scanln(&operation)
				switch operation {
				case 1:
					message := bank.Query(bankCode, bankPwd)
					fmt.Printf("%s", message)
					break
				case 2:
					var money float64
					fmt.Print("请输入存款金额：")
					_, _ = fmt.Scanln(&money)
					message := bank.Deposit(bankCode, bankPwd, money)
					fmt.Printf("%s", message)
					break
				case 3:
					var money float64
					fmt.Print("请输入取款金额：")
					_, _ = fmt.Scanln(&money)
					message := bank.Withdraw(bankCode, bankPwd, money)
					fmt.Printf("%s", message)
					break
				case 4:
					fmt.Println("感谢您的使用！")
					return
				}
			}
		}

	}

}
