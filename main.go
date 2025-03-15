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

// func main() {

// 	fmt.Println(sqrt(4), sqrt(-6))
// }

// func main() {

// 	y := float64(4)

// 	if x := sqrt(y); x == "2" {

// 		fmt.Println("hhhh")
// 	}

// 	// : undefined: x
// 	// fmt.Println(x)
// }

// func main() {
// 	a := 2
// 	if a > 3 {
// 		fmt.Println("hey")
// 	} else {
// 		fmt.Println("no")
// 	}
// }

func main() {
	a := 100
	if a > 100 {
		fmt.Println("hey")
	} else if a > 50 && a < 100 {
		fmt.Println("no")
	} else {
		fmt.Println("Boom")
	}

	switch {
	case a > 100:
		fmt.Println("hey")
	case a > 50 && a < 100:
		fmt.Println("no")
	default:
		fmt.Println("Boom")

	}

	s := "hello"
	switch s {
	case "helll":
		fmt.Println("nop1")
	case "helloooo":
		fmt.Println("nop")
	case "hello":
		fmt.Println("Yesss")
	}

}
