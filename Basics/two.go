package main

import "fmt"

func main() {
	//先定义再赋值
	//定义变量  var
	//定义常量  const

	//先定义再复制
	var username string
	username = "Mr Ding"
	fmt.Println(username)

	//直接定义赋值
	var age = 30
	fmt.Printf("name is :%s,年龄 %d\n", username, age)

	//定义时直接赋值，使用推导(最常用的)
	sex := "男"
	fmt.Printf("name is :%s,年龄 %d, 性别 :%s\n", username, age, sex)

	//形参
	tesst(age, 1992)

	//平行赋值
	num, num1 := 10, 20
	fmt.Printf("赋值前=====> :%d %d\n", num, num1)

	num, num1 = num1, num
	fmt.Printf("赋值后=====> :%d %d", num, num1)

}
func tesst(age int, year int) {
	fmt.Printf("年龄 :%d", age)
	fmt.Printf("出生年 :%d\n", year)
}
