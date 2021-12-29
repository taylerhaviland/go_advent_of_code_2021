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

	scanner := bufio.NewScanner(file)

	defer file.Close()
	m := map[int]int{
		0:  0,
		1:  0,
		2:  0,
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
		9:  0,
		10: 0,
		11: 0,
	}
	for scanner.Scan() {

		binary_val := strings.Split(scanner.Text(), "")
		for i, val := range binary_val {
			int_val, _ := strconv.Atoi(val)
			if int_val == 1 {
				m[i]++
			} else {
				m[i]--
			}
		}

	}

	fmt.Println(m)
	g_rate := ""
	e_rate := ""
	for i := 0; i < len(m); i++ {
		fmt.Println(i, m[i])
		if m[i] >= 0 {
			g_rate += "1"
			e_rate += "0"
		} else {
			g_rate += "0"
			e_rate += "1"
		}
	}

	fmt.Println(g_rate)
	fmt.Println(e_rate)

	gr_value, _ := strconv.ParseInt(g_rate, 2, 64)
	er_value, _ := strconv.ParseInt(e_rate, 2, 64)

	fmt.Println(gr_value)
	fmt.Println(er_value)
	fmt.Println(gr_value * er_value)
}
