package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := ReadInputFile("input.txt")
	result := ParseCorruptString(input)

	fmt.Printf("Multiplication Sum = %d\n", result)
	result = ParseCorruptStringEnabled(input)

	fmt.Printf("Multiplication Sum Enabled = %d\n", result)
}

func ReadInputFile(filepath string) string {
	file, err := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	return string(file)
}

func ParseCorruptString(input string) int {
	result := 0

	re := regexp.MustCompile(`(?i)mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, v := range matches {
		n1, err1 := strconv.Atoi(v[1])
		n2, err2 := strconv.Atoi(v[2])

		if err1 != nil {
			panic(err1)
		}
		if err2 != nil {
			panic(err2)
		}

		result += n1 * n2
	}

	return result
}

func ParseCorruptStringEnabled(input string) int {
	result := 0
	enabled := true

	const ENABLE = "do()"
	const DISABLE = "don't()"

	re := regexp.MustCompile(`(?i)mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, v := range matches {
		switch v[0] {
		case ENABLE:
			enabled = true
		case DISABLE:
			enabled = false
		default:
			if !enabled {
				break
			}

			n1, err1 := strconv.Atoi(v[1])
			n2, err2 := strconv.Atoi(v[2])

			if err1 != nil {
				panic(err1)
			}
			if err2 != nil {
				panic(err2)
			}

			result += n1 * n2
		}
	}

	return result
}
