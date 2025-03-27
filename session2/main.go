//parse args without using flag package

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
		log.Fatalf("There is   no input, Len of args is : %d", len(args))

	}

	name := args[1]
	fmt.Println("Hello", name)
	name2 := args[2]
	fmt.Println("Hello", name2)
	name = strings.Replace(name, "--first_name=", "", 1)
	fmt.Println("Hello", name)

}
