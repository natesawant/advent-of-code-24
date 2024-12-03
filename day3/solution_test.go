package main

import "testing"

func TestParseCorruptString(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161

	actual := ParseCorruptString(input)

	if actual != expected {
		t.Errorf("Actual is different from expected: %d != %d", actual, expected)
	}
}

func TestParseCorruptStringEnable(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	expected := 48

	actual := ParseCorruptStringEnabled(input)

	if actual != expected {
		t.Errorf("Actual is different from expected: %d != %d", actual, expected)
	}
}
