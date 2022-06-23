package user

import "fmt"

// Bank 定义一个结构体
type Bank struct {
	BankCode string
	BankPwd  string
	Balance  float64
}

// Deposit 存款
func (bank *Bank) Deposit(bankCode string, bankPwd string, money float64) string {
	var message string
	if bankCode != bank.BankCode || bankPwd != bank.BankPwd {
		message = "您输入的账号或者密码不正确!"
		return message
	}
	if money <= 0 {
		message = "存款金额不正确"
		return message
	}
	bank.Balance += money
	message = fmt.Sprintf("存款成功！\n您账户当前余额:%.2f", bank.Balance)
	return message
}

// Withdraw 取款
func (bank *Bank) Withdraw(bankCode string, bankPwd string, money float64) string {
	var message string
	if bankCode != bank.BankCode || bankPwd != bank.BankPwd {
		message = "您输入的账号或者密码不正确!"
		return message
	}
	if money <= 0 || money > bank.Balance {
		message = "取款金额不正确"
		return message
	}
	bank.Balance -= money
	message = fmt.Sprintf("取款成功！\n您账户当前余额:%.2f", bank.Balance)
	return message
}

// Query 查询
func (bank *Bank) Query(bankCode string, bankPwd string) string {
	var message string
	if bankCode != bank.BankCode || bankPwd != bank.BankPwd {
		message = "您输入的账号或者密码不正确!"
		return message
	}
	message = fmt.Sprintf("查询成功！\n您账户当前余额:%.2f", bank.Balance)
	return message
}

// Verification 验证账号密码是否正确,不正确3此过后锁卡
func (bank *Bank) Verification(bankCode string, bankPwd string) (bool, string) {
	var message string
	flag := false
	if bankCode != bank.BankCode || bankPwd != bank.BankPwd {
		message = "您输入的账号或者密码不正确!"
		return flag, message
	}
	flag = true
	return flag, message
}
