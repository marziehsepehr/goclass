package student

import "fmt"

type Student struct {
	ID   int16
	Name string
	age  int
}

// PrintName is an exported method

func (s Student) PrintName() {
	fmt.Println("Name is:", s.Name)
}

// printID is an unexported method

func (s Student) printID() {
	fmt.Println("ID is:", s.ID)

}
