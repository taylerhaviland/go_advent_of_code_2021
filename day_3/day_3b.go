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

	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	bitcount := 0
	og_rating := ""
	co2s_rating := ""
	for bitcount < 12 {
		count_0 := 0
		count_1 := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			binary_val := strings.Split(scanner.Text(), "")
			int_val, _ := strconv.Atoi(binary_val[bitcount])
			if int_val == 0 {
				count_0++
			} else {
				count_1++
			}
		}

		if count_0 > count_1 {
			fmt.Printf("Bit index: %d bit 0 count: %d, bit 1 count: %d\n", bitcount, count_0, count_1)
			og_rating += "0"
			co2s_rating += "1"
		} else {
			fmt.Printf("Bit index: %d bit 0 count: %d, bit 1 count: %d\n", bitcount, count_0, count_1)
			og_rating += "1"
			co2s_rating += "0"
		}

		fmt.Printf("og_rating: %s\n", og_rating)
		fmt.Printf("co2_rating: %s\n", co2s_rating)

		_, err := file.Seek(0, 0)
		check(err)
		bitcount++
	}

	ogr_val, _ := strconv.ParseInt(og_rating, 2, 64)
	co2sr_val, _ := strconv.ParseInt(co2s_rating, 2, 64)

	fmt.Printf("OG Rating: %d\n", ogr_val)
	fmt.Printf("CO2 Rating: %d\n", co2sr_val)
	fmt.Printf("Life Support Rating: %d\n", ogr_val*co2sr_val)

}
