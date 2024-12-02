package main

import (
	"reflect"
	"testing"
)

type report struct {
	levels []int
	safe   bool
}

func TestSafeLevels(t *testing.T) {
	var tests = []report{
		{
			levels: []int{7, 6, 4, 2, 1},
			safe:   true,
		},
		{
			levels: []int{1, 2, 7, 8, 9},
			safe:   false,
		},
		{
			levels: []int{9, 7, 6, 2, 1},
			safe:   false,
		},
		{
			levels: []int{1, 3, 2, 4, 5},
			safe:   false,
		},
		{
			levels: []int{8, 6, 4, 4, 1},
			safe:   false,
		},
		{
			levels: []int{1, 3, 6, 7, 9},
			safe:   true,
		},
	}

	for i, report := range tests {
		isSafe := SafeLevels(report.levels)

		if isSafe != report.safe {
			t.Errorf("Test %d: Does not match safety - Expected: %t, Actual: %t", i, report.safe, isSafe)
		}
	}
}

func TestCountSafeReports(t *testing.T) {
	reports := ParseFile("test.txt")
	count := CountSafeReports(reports)
	expected := 2

	if count != expected {
		t.Errorf("Actual does not match expected: %d != %d\n", count, expected)
	}
}

func TestParseFile(t *testing.T) {
	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	actual := ParseFile("test.txt")

	if !reflect.DeepEqual(expected, actual) {
		t.Error("Actual does not match expected")
	}
}
