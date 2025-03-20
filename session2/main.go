package main

import (
	"fmt"
	"os"
)

//	comman-line interface
//
// gui
func main() {

	args := os.Args
	if len(args) > 1 {
		name := args[1]
		fmt.Println(name)
	} else {
		fmt.Println("At least one argument is needed")
	}

}
