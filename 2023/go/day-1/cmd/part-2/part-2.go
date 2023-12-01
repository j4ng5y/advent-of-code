package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	newstr := s
	newstr = strings.ReplaceAll(newstr, "one", "o1e")
	newstr = strings.ReplaceAll(newstr, "two", "t2o")
	newstr = strings.ReplaceAll(newstr, "three", "t3e")
	newstr = strings.ReplaceAll(newstr, "four", "f4r")
	newstr = strings.ReplaceAll(newstr, "five", "f5v")
	newstr = strings.ReplaceAll(newstr, "six", "s6x")
	newstr = strings.ReplaceAll(newstr, "seven", "s7n")
	newstr = strings.ReplaceAll(newstr, "eight", "e8t")
	newstr = strings.ReplaceAll(newstr, "nine", "n9e")

	for _, c := range newstr {
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

	fmt.Printf("original str: %s, final str: %s, first and last: %d\n", s, newstr, d)

	return d
}
