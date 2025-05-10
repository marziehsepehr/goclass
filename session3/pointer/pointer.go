package main

import "fmt"

func main() {

	var p *int
	i := 43

	fmt.Println("p", p)
	fmt.Println("i", i)

	p = &i
	fmt.Println(p)
	fmt.Println(p)
}
