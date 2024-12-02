package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := ParseFile("input.txt")

	count := CountSafeReports(reports)
	fmt.Printf("Safe Reports = %d\n", count)

	count = CountSoftSafeReports(reports)
	fmt.Printf("Soft Safe Reports = %d\n", count)
}

func CountSafeReports(reports [][]int) int {
	count := 0

	for _, levels := range reports {
		if SafeLevels(levels) {
			count += 1
		}
	}

	return count
}

func CountSoftSafeReports(reports [][]int) int {
	count := 0

	for _, levels := range reports {
		if SoftSafeLevels(levels) {
			count += 1
		}
	}

	return count
}

func SafeLevels(levels []int) bool {
	var prev int = levels[0]

	var isAscending = levels[0] < levels[1]

	for _, curr := range levels[1:] {
		diff := Abs(curr - prev)

		if !checkLevelDiff(diff) {
			return false
		}

		if isAscending && !checkAscending(prev, curr) {
			return false
		}

		if !isAscending && !checkDescending(prev, curr) {
			return false
		}

		prev = curr
	}
	return true
}

func SoftSafeLevels(levels []int) bool {
	for i := range len(levels) {
		copySlice := make([]int, len(levels))
		copy(copySlice, levels)

		removed := remove(copySlice, i)
		if SafeLevels(removed) {
			return true
		}
	}

	return false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func ParseFile(filepath string) [][]int {
	var reports = [][]int{}

	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		levels := []int{}

		nums := strings.Split(scanner.Text(), " ")
		for _, n := range nums {
			val, _ := strconv.Atoi(n)
			levels = append(levels, val)
		}

		reports = append(reports, levels)
	}

	file.Close()

	return reports
}

func checkAscending(a int, b int) bool {
	return a < b
}

func checkDescending(a int, b int) bool {
	return a > b
}

func checkLevelDiff(diff int) bool {
	return diff >= 1 && diff <= 3
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
