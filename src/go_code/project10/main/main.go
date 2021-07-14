package main

import "fmt"

func main() {

	//使用while的方式输出十句hello world
	var i int = 1
	for {
		if i > 10 { //循环条件
			break //跳出循环
		}
		fmt.Println("Hello world!", i)
		i++ //循环变量的迭代
	}

	//使用do-while实现输出10句“Hello golang”
	var j int = 1
	for {
		fmt.Println("Hello golang!", j)
		j++
		if j > 10 {
			break
		}
	}

	//统计3个班的成绩，每班有五名同学，求平均分
	var f, sum float64
	for p := 1; p <= 3; p++ {
		var a, b, c, d, e float64
		fmt.Println("请输入成绩：")
		fmt.Scanln(&a, &b, &c, &d, &e)
		f = (a + b + c + d + e) / 5
		fmt.Printf("该班平均分：%v\n", f)
		sum += f
	}
	fmt.Printf("所有班级平均分：%v\n", sum/3)

	//同上
	var classNum int
	var stuNum int
	var totalSum = 0.0
	var passCount = 0

	fmt.Println("请输入班级个数：")
	fmt.Scanln(&classNum)
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		fmt.Printf("请输入第%d班的学生个数:\n", j)
		fmt.Scanln(&stuNum)
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩。\n", j, i)
			fmt.Scanln(&score)
			//累计得分
			sum += score
			//判断分数是否及格
			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("第%d个班级的平均分是%v\n", j, sum/float64(stuNum))
		//将各个班级的总成绩累积到totalSum
		totalSum += sum
	}
	fmt.Printf("各个班级的总成绩%v 所有班级的平均分是%v\n", totalSum, totalSum/float64(stuNum*classNum))
	fmt.Printf("及格人数为%v", passCount)
}
