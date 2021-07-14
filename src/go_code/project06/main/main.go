package main

import "fmt"

func main() {

	var a float64
	var b float64
	fmt.Println("请输入两个float64型变量（空格隔开）：")
	fmt.Scanln(&a, &b)

	/*if a > 10.0 {
		if b < 20.0 {
			fmt.Println(a + b)
		}else{
			fmt.Println("不符合要求！")
		}
	}else{
		fmt.Println("不符合要求！")
	}*/
	if a > 10.0 && b < 20.0 {
		fmt.Println(a + b)
	} else {
		fmt.Println("不符合要求！")
	}
}
