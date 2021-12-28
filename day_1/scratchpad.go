package main

import "fmt"

func main() {
	// var num int = 10
	var point_1 *int = nil

	if point_1 == nil {
		fmt.Printf("point_1 is nil\n")
	} else {
		fmt.Printf("point_1 is not nil\n")
	}

	// fmt.Printf("num value: %d\n", num)
	// fmt.Printf("point_1 value: %d\n", *point_1)
}
