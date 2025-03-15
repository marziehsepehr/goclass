package main

import (
	"fmt"
	"math"
)

// func main() {
// 	//
// 	var i = 0
// 	fmt.Println(i)
// 	//
// 	sum := 0

// 	for i := 0; i < 10; i += 2 {
// 		sum += i
// 	}

// 	fmt.Println(sum)
// 	//

// 	j := 0

// 	for j < 100 {
// 		j = j + 1
// 		fmt.Println(j)

// 	}

// }

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main() {

	fmt.Println(sqrt(4), sqrt(-6))
}
