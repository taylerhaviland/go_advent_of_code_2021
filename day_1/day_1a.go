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

	fmt.Printf("Creating a scanner")
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
	var numpointer *int = nil
	for _, read_ln := range text {
		if numpointer == nil {
			num := read_ln
			numpointer = &num
			fmt.Printf("%d - first number\n", num)
		} else {
			if read_ln > *numpointer {
				increase++
				fmt.Printf("%d - increased\n", read_ln)
				*numpointer = read_ln
			} else if read_ln < *numpointer {
				decrease++
				fmt.Printf("%d - decreased\n", read_ln)
				*numpointer = read_ln
			} else {
				fmt.Printf("Equal number\n")
				*numpointer = read_ln
			}
		}
	}

	fmt.Printf("Increase: %d\n", increase)
	fmt.Printf("Decrease: %d\n", decrease)
}
