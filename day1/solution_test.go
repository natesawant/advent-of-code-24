package main

import (
	"reflect"
	"testing"
)

func TestTotalDistance(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	ans := TotalDistance(a, b)
	if ans != 11 {
		t.Errorf("TotalDistance(a, b) = %d; want 11", ans)
	}
}

func TestSimilarityScore(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	ans := SimilarityScore(a, b)
	if ans != 31 {
		t.Errorf("SimilarityScore(a, b) = %d; want 31", ans)
	}
}

func TestParseFile(t *testing.T) {
	a := []int{3, 4, 2, 1, 3, 3}
	b := []int{4, 3, 5, 3, 9, 3}

	act_a, act_b := ParseFile("test.txt")

	if !reflect.DeepEqual(a, act_a) || !reflect.DeepEqual(b, act_b) {
		t.Error("ParseFile('test.txt') is different")
	}
}
