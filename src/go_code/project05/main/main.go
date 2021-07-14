package main

import "fmt"

func main() {
	//分支使用
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("您已成年!")
	} else {
		fmt.Println("您未成年!")
	}

	//分支使用
	var a int32
	var b int32
	fmt.Println("请输入两个数(请使用空格隔开)：")
	fmt.Scanln(&a, &b)
	if a+b >= 50 {
		fmt.Println("Hello world!")
	}

	//
}
