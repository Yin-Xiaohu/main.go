package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//随机生成1-100整数
	var count int = 0
	rand.Seed(time.Now().Unix())
	for {
		n := rand.Intn(100)
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共使用了", count)

	//指定标签的形式来使用break
liable1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break liable1
			}
			fmt.Println("j=", j)
		}
	}

	//100以内数求和，求出第一次大于20的当前数
	var sum int = 0
	for i := 1; i < 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("i=", i)
			break
		}
	}

	//实现登录验证，有三次机会
	var name string
	var pwd string
	var i int = 1
	for ; i <= 3; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pwd)

		if name == "张无忌" && pwd == "888" {
			fmt.Println("登录成功！")
			break
		} else if i == 3 {
			fmt.Printf("您已没有机会，请退出！")
		} else {
			fmt.Printf("登录错误，请重新输入！您还有%v次机会 \n", 3-i)
		}
	}
}
