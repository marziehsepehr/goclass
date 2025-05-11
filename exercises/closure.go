package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() []int {
	// 0, 1, 1, 2, 3, 5, ...

	var fibList []int
	return func() []int {

		if len(fibList) == 0 {
			fibList = append(fibList, 0)

		} else if len(fibList) == 1 {
			fibList = append(fibList, 1)
		} else {
			next := fibList[len(fibList)-1] + fibList[len(fibList)-2]
			fibList = append(fibList, next)
		}

		return fibList
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
