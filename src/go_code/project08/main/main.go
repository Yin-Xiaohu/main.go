package main

import "fmt"

func main() {
	//小写转大写
	var char byte
	fmt.Println("请输入一个字母（a-e）：")
	fmt.Scanf("%c", &char)

	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")
	}
	//输入月份知道是什么季节
	var month byte
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("春季！")
		fallthrough
	case 6, 7, 8:
		fmt.Println("夏季！")
	case 9, 10, 11:
		fmt.Println("秋季！")
	case 12, 1, 2:
		fmt.Println("冬季！")
	default:
		fmt.Println("输入错误！")
	}
}
