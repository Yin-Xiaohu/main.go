package main

import "fmt"

func main() {
	//for循环控制
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello world!")
	}

	//字符串遍历方式1
	var str = "Hello world!"
	for i := 0; i < len(str); i++ {
		fmt.Printf(
			"%c \n",
			str[i],
		) //使用到下标
	}
	fmt.Println()

	//字符串遍历方式1（使用切片解决中文乱码）
	var str1 = "Hello world!北京"
	str2 := []rune(str1) //就是把 str 转成[]rune
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i]) //使用到下标
	}
	fmt.Println()
	//字符串遍历方式2
	str = "abc~ok"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

	//打印1~100之间所有是9的倍数的整数的个数及总和
	var i int
	var j = 0
	var k int
	for i = 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Printf("%d是9的倍数\n", i)
			j++
			k += i
		}
	}
	fmt.Println("个数", j)
	fmt.Println("总和", k)

	//练习题
	var a = 9
	var q = 0
	for ; q <= a; q++ {
		for ; a >= 0; a-- {
			fmt.Printf(
				"%v + %v = %v \n",
				q,
				a,
				q+a,
			)
			q++
		}
	}

}
