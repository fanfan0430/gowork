package demo

import (
	"fmt"
	"strconv"
	"strings"
	"math/rand"
)

func Add(x, y int) int {
	z := x + y
	printLog(strconv.Itoa(z))
	return z
}

func printLog(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str)
		result := strings.EqualFold(strconv.Itoa(rand.Intn(10)), str)
		if result {
			fmt.Println("result:", result, ", str:" + str)
		} else {
			fmt.Println("false:" + str)
		}
	}
}
