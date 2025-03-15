package main

import "fmt"

func main() {
	var i = 0
	fmt.Println(i)

	sum := 0

	for i := 0; i < 10; i += 2 {
		sum += i
	}

	fmt.Println(sum)

}
