package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fileStat, err := os.Stat("input.txt")
	check(err)
	fileSize := fileStat.Size()

	fmt.Printf("File size: %d bytes\n", fileSize)

	file, err := os.Open("input.txt")
	check(err)

	fmt.Printf("Creating a scanner\n")
	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)
	var text []int

	defer file.Close()

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		num, _ := strconv.Atoi(scanner.Text())
		text = append(text, num)
	}

	increase := 0
	decrease := 0
	var numpointer_1 *int = nil
	var numpointer_2 *int = nil
	var numpointer_3 *int = nil
	var numpointer_4 *int = nil
	for _, read_ln := range text {
		num := read_ln
		numpointer_4 = numpointer_3
		numpointer_3 = numpointer_2
		numpointer_2 = numpointer_1
		numpointer_1 = &num
		if numpointer_1 != nil && numpointer_2 != nil && numpointer_3 != nil && numpointer_4 != nil {

			total_1 := *numpointer_1 + *numpointer_2 + *numpointer_3
			total_2 := *numpointer_2 + *numpointer_3 + *numpointer_4
			if total_1 > total_2 {
				increase++
				fmt.Printf("%d - increased (%d + %d + %d)\n", total_1, *numpointer_1, *numpointer_2, *numpointer_3)

			} else if total_1 < total_2 {
				decrease++
				fmt.Printf("%d - decreased (%d + %d + %d)\n", total_1, *numpointer_1, *numpointer_2, *numpointer_3)
			} else {
				fmt.Printf("Equal number\n")
			}
		}
	}

	fmt.Printf("Increase: %d\n", increase)
	fmt.Printf("Decrease: %d\n", decrease)
}
