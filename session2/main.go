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
		/*
			fmt.Println("error: At least one argument is needed")
			os.Exit(1)
		*/
		// Instead of print in colcole and exit with two line ==>logs.Fatalf
		log.Fatalf("There is no input, Len of args is : %d", len(args))

	}

}
