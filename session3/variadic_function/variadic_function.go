package main

import "fmt"

func prt(values ...string) {
	for _, value := range values {
		fmt.Println(value)

	}
}

func main() {
	prt("Hello", "World", "Go", "is", "awesome")
	prt("hiiiiiii", "sdgfsdzgsd")
}
