package main

import (
	"fmt"
)

func myPrint(value int) {
	fmt.Println(value)
}
func myFprint(value int) {
	fmt.Printf("The Value is :%d \n", value)
}

func Power(x, y int, f func(int)) int {
	sum := x
	for i := 0; i < y; i++ {
		sum *= x
	}
	f(sum)
	return sum
}

func main() {

	Power(2, 3, myPrint)

	Power(2, 3, myFprint)

}
