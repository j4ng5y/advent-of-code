package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := calculate_calibration_value(scanner.Text())
		acc += i
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	print(acc)
}

func calculate_calibration_value(s string) int {
	var digits []int
	for _, c := range s {
		i, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		digits = append(digits, i)
	}

	dstr := fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
	d, err := strconv.Atoi(dstr)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
