//每个Go程序都必须有一个包名 package main
//每个go程序都是.go结尾，没有.h/.o只有.go
package main

import "fmt"

//主函数，所有的函数必须使用func开头
//一个函数的返回值不会放在func前，二十放在参数后面
//函数左花括号必须与函数名同行，不能写到下一行
func main() {
	//go语言不需要‘;’号结尾
	//fmt.Println("hello word")

	// 演示for-range遍历数组
	var heroes [3]string = [3]string{"宋江", "吴用", "卢俊义"}
	for index, value := range heroes {
		fmt.Printf("数组heroes[%d] = %v \n", index, value)
	}

}
