package main

import (
	"fmt"
	"log"
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
		fmt.Println("error: At least one argument is needed")
		os.Exit(1)
		
	}

}
