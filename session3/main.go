// package main

// import "fmt"

// type Point struct {
// 	lat, long float64
// }

// // global map decleration with key as string and value a Point struct
// var points map[string]Point

// func main() {

// 	if points == nil {

// 		fmt.Println("not init yet!")
// 	}
// 	// initialize the map
// 	points = make(map[string]Point)
// 	if points == nil {

// 		fmt.Println("not init yet!")
// 	} else {
// 		fmt.Println("initialized!")
// 	}
// 	points["Home"] = Point{34.555, 43.222}
// 	points["Home2"] = Point{34.555, 43.222}
// 	points["Home3"] = Point{34.555, 43.222}
// 	fmt.Println(points["Home"])

// 	fmt.Println("All points are:", points)
// 	var k = "Home3"
// 	value, ok := points[k]
// 	if ok {
// 		fmt.Printf("key: %s exist in points and the value is lat: %0.3f long:%0.3f\n", k, value.lat, value.long)
// 		delete(points, k)
// 		fmt.Printf("after deleteing %s\n", k)
// 		fmt.Println("All points are:", points)
// 	} else {
// 		fmt.Printf("key: %s not exist in points", k)
// 	}

// }

package main

import "gocasts.ir/go-fundamentals/repo/sessions/session3/student"

func main() {
	s := student.Student{}
	s.ID = 1
	s.Name = "Ali"
	// s.age = 20

	s.PrintName()
	// s.printID()

}
