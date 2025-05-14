package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func generateUserID() int {
	return rand.Int()
}

var userStorage []User

func main() {
	fmt.Println("Hello TODO app")
	command := flag.String("command", "No command", "Command to run")
	flag.Parse()
	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please Enter next command")
		scanner.Scan()
		*command = scanner.Text()
	}

	fmt.Println("userStorage", userStorage)
}
func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)

	default:
		fmt.Println("Unknown command:", command)
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, dueDate, category string
	fmt.Println("please enter the task title:")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the due date:")
	scanner.Scan()
	dueDate = scanner.Text()

	fmt.Println("please enter the category:")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("task:", name, category, dueDate)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var categoryName string
	fmt.Println("please enter the category name:")
	scanner.Scan()
	categoryName = scanner.Text()

	fmt.Println("please enter the category color:")
	scanner.Scan()
	categoryColor := scanner.Text()

	fmt.Println("category created:", categoryName, "with color:", categoryColor)
}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string
	fmt.Println("please enter the email:")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	newUser := User{
		ID:       generateUserID(),
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, newUser)

	fmt.Println("user registered with email:", email, "and ID:", newUser.ID)
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string
	fmt.Println("please enter your email:")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter your password:")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			fmt.Println("login successful for user with email:", email)

		}
	}

	fmt.Println("login failed: invalid email or password")

	fmt.Println("All registered users:")
	for _, u := range userStorage {
		fmt.Printf("ID: %d, Email: %s\n", u.ID, u.Email)
	}

}
