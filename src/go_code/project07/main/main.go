package main

import (
	"fmt"
	"math"
)

func main() {
	//一元二次方程求解
	var a, b, c float64
	fmt.Println("请输入值：")
	fmt.Scanln(&a, &b, &c)

	if b*b-4*a*c > 0 {
		fmt.Println((-b+math.Sqrt(b*b-4*a*c))/(2*a), (-b-math.Sqrt(b*b-4*a*c))/(2*a))
	} else if b*b-4*a*c == 0 {
		fmt.Println((-b + math.Sqrt(b*b-4*a*c)) / (2 * a))
	} else {
		fmt.Println("无！")
	}
}
