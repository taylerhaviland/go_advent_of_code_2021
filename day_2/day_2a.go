package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	input_slice := make([]map[string]interface{}, 1)

	defer file.Close()

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		// num, _ := strconv.Atoi(scanner.Text())
		m := make(map[string]interface{})
		s := strings.Split(scanner.Text(), " ")

		distance_int, _ := strconv.Atoi(s[1])

		m["direction"] = s[0]
		m["distance"] = distance_int

		input_slice = append(input_slice, m)
	}

	hdis := 0
	depth := 0
	for _, object := range input_slice {

		switch object["direction"] {
		case "up":
			depth -= object["distance"].(int)
		case "down":
			depth += object["distance"].(int)
		case "forward":
			hdis += object["distance"].(int)
		}
		fmt.Printf("Direction: %s distance: %d\n", object["direction"], object["distance"])
	}

	fmt.Printf("depth: %d hdis: %d\n", depth, hdis)
	fmt.Printf("Total distance: %d\n", depth*hdis)

}
