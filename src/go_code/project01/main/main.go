package main

import "fmt"

func main() {
	//Hello world
	fmt.Println("Hello World!")

	//打印一个心型图案
	fmt.Println("     **\t    **\t\n",
		"  *     *     *\n",
		"*      安迪     *\n",
		"   *         *\n",
		"      *   *\n",
		"        *")

	//Hello golang
	fmt.Println("Hello,Golang!")

	//转义字符
	fmt.Println("姓名\t性别\t籍贯\t住址\ntom\t男\t甘肃\t北京\t")

	//求地址
	var num int
	fmt.Printf("num的地址=%v", &num)

	//除法运算符
	n1 := 10 / 4
	fmt.Println(n1)

	var n2 float32 = 10 / 4
	fmt.Println(n2)

	var n3 float32 = 10.0 / 4
	fmt.Println(n3)

	n4 := 10.0 / 4
	fmt.Println(n4)

	//运算符
	var dayOne int = 97
	var week int = dayOne / 7
	var dayTwo int = dayOne % 7

	fmt.Println(week, dayTwo)

	//位运算符
	var a int = -1 >> 2
	fmt.Println("a=", a)
	fmt.Println(2 & 3)
	fmt.Println(2 ^ 3)

}
