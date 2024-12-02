package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left, right := ParseFile("input.txt")

	distance := TotalDistance(left, right)
	fmt.Printf("Total Distance = %d\n", distance)

	similarity := SimilarityScore(left, right)
	fmt.Printf("Similarity Score = %d\n", similarity)
}

func TotalDistance(left []int, right []int) int {
	result := 0

	if len(left) != len(right) {
		panic("Arrays must be same length")
	}

	sort.Ints(left)
	sort.Ints(right)

	for i := range len(left) {
		dist := left[i] - right[i]
		if dist < 0 {
			dist *= -1
		}
		result += dist
	}

	return result
}

func SimilarityScore(left []int, right []int) int {
	result := 0

	count := map[int]int{}

	for _, num := range right {
		count[num] += 1
	}

	for _, num := range left {
		result += num * count[num]
	}

	return result
}

func ParseFile(filepath string) ([]int, []int) {
	var left []int
	var right []int

	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		a, _ := strconv.Atoi(nums[0])
		b, _ := strconv.Atoi((nums[1]))

		left = append(left, a)
		right = append(right, b)
	}

	file.Close()

	return left, right
}
