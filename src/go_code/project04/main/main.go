package main

import (
	"fmt"
	"go_code/project04/model"
)

func main() {

	var num int

	fmt.Printf("num的地址=%v\n", &num)
	fmt.Println(model.HeroName)
}
