package main

import "fmt"

// import (
// 	"fmt"
// 	"math"
// )

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

// func sqrt(x float64) string {
// 	if x < 0 {
// 		return sqrt(-x) + "i"
// 	}

// 	return fmt.Sprint(math.Sqrt(x))
// }

// // func main() {

// // 	fmt.Println(sqrt(4), sqrt(-6))
// // }

// // func main() {

// // 	y := float64(4)

// // 	if x := sqrt(y); x == "2" {

// // 		fmt.Println("hhhh")
// // 	}

// // 	// : undefined: x
// // 	// fmt.Println(x)
// // }

// // func main() {
// // 	a := 2
// // 	if a > 3 {
// // 		fmt.Println("hey")
// // 	} else {
// // 		fmt.Println("no")
// // 	}
// // }

// func main() {
// 	a := 100
// 	if a > 100 {
// 		fmt.Println("hey")
// 	} else if a > 50 && a < 100 {
// 		fmt.Println("no")
// 	} else {
// 		fmt.Println("Boom")
// 	}

// 	switch {
// 	case a > 100:
// 		fmt.Println("hey")
// 	case a > 50 && a < 100:
// 		fmt.Println("no")
// 	default:
// 		fmt.Println("Boom")

// 	}

// 	s := "hello"
// 	switch s {
// 	case "helll":
// 		fmt.Println("nop1")
// 	case "helloooo":
// 		fmt.Println("nop")
// 	case "hello":
// 		fmt.Println("Yesss")

// 	}
// 	a = 60
// 	switch {
// 	case a > 100:
// 		fmt.Println("hey")
// 	case a > 50 && a < 100:
// 		fmt.Println("yesss")
// 		fallthrough

// 	case a > 50 && a < 70:
// 		fmt.Println("yess2")
// 	default:
// 		fmt.Println("Boom")

// 	}

// }

// func main() {

// 	type Student struct {
// 		ID             uint
// 		NationalNumber string
// 		Age            uint
// 	}
// 	// declare a variable of type Student with Zero value for its fields
// 	var u Student
// 	fmt.Println(u.Age)
// 	fmt.Println(u.NationalNumber)

// 	var j = Student{

// 		Age: 23,
// 	}
// 	fmt.Println(j.Age)
// 	fmt.Println(j.NationalNumber)

// 	var z = Student{10, "00555", 22}
// 	fmt.Println(z.ID)
// 	fmt.Println(z.Age)
// 	fmt.Println(z.NationalNumber)
// 	fmt.Println("Z Id is:", z.ID, "Z Age is:", z.Age)
// 	fmt.Printf("Z Id is: %d Z Age is: %d National is %s", z.ID, z.Age, z.NationalNumber)

// }

func main() {
	// Arrays

	var a [6]int
	a = [6]int{1, 3, 2, 5, 7, 8}
	for index, value := range a {
		fmt.Println(index, value)
	}

	// Slice(Dynamic Array)
	var b []int
	// builtin function
	b = append(b, 5)
	fmt.Println(cap(b))
	fmt.Println(len(b))

	c := make([]int, 5, 10)
	fmt.Println(cap(c))
	fmt.Println(len(c))

}
