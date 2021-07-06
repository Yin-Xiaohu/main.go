package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n1 = 100
	fmt.Printf("n1=%T\n", n1)
	var n2 int32 = 10
	fmt.Printf("n2=%T，n2=%d", n2, unsafe.Sizeof(n2))

}
