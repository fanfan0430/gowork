package demo

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(x, y int) int {
	z := x + y
	printLog(strconv.Itoa(z))
	return z
}

func printLog(str string) {
	fmt.Println(str)
	fmt.Println(strings.EqualFold("a", "b"))
}
