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
