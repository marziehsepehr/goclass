// //parse args without using flag package

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// //	comman-line interface
// //
// // gui
// func main() {

// 	args := os.Args
// 	if len(args) > 1 {
// 		name := args[1]
// 		fmt.Println(name)
// 	} else {
// 		/*
// 			fmt.Println("error: At least one argument is needed")
// 			os.Exit(1)
// 		*/
// 		// Instead of print in colcole and exit with two line ==>logs.Fatalf
// 		log.Fatalf("There is   no input, Len of args is : %d", len(args))

// 	}

// 	name := args[1]
// 	fmt.Println("Hello", name)
// 	name2 := args[2]
// 	fmt.Println("Hello", name2)
// 	name = strings.Replace(name, "--first_name=", "", 1)
// 	fmt.Println("Hello", name)

// }

// parse args with flag package

package main

import (
	"fmt"
)

func main() {
	// firstName := flag.String("first_name", "default", "first name")
	// lastName := flag.String("last_name", "default", "last name")

	// flag.Parse()
	// fmt.Println("name", *firstName, *lastName, flag.Args())

	// fmt.Println("len of remaining args", len(flag.Args()))

	// var fname, lname string
	// fmt.Println("Enter first name and last name")

	// fmt.Scanln(&fname, &lname)
	// fmt.Println(fname, lname)

	// get parammeter by echo or csv as input - ex: echo "2025, 03, 27"| go run main.go   or cat input.txt | go run main.go

	var a, b, c int
	fmt.Println("Enter three numbers with comma seperated")

	fmt.Scanf("%d, %d, %d", &a, &b, &c)
	fmt.Println("Your numbers are:")
	fmt.Println(a, b, c)

}
