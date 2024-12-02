package utils

import (
	"aoc2024/testing_utils"
	"testing"
)

func TestGetLines(t *testing.T) {
	input := `This is
	a multiline string
	with a few lines`

	got := GetLines(input)
	want := []string{"This is", "a multiline string", "with a few lines"}

	testing_utils.DeepCompareStrings(t, got, want)
}

func TestStringToInts(t *testing.T) {
	input := "1 2 3 4"

	got := StringToInts(input)
	want := []int{1, 2, 3, 4}

	testing_utils.DeepCompareInts(t, got, want)
}
