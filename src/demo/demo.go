package demo

import (
	"fmt"
	"strconv"
)

func Add(x, y int) int {
	z := x + y
	printLog(strconv.Itoa(z))
	return z
}

func printLog(str string) {
	fmt.Println(str)
}
