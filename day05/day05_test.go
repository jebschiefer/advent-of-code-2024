package day05

import (
	"aoc2024/testing_utils"
	"aoc2024/utils"
	"testing"
)

func TestDay05Part1Sample(t *testing.T) {
	input := utils.ReadFile("../inputs/day05_sample.txt")

	got := Day05(input)
	want := 143

	testing_utils.CompareInts(t, got, want)
}

func TestDay05Part1(t *testing.T) {
	input := utils.ReadFile("../inputs/day05.txt")

	got := Day05(input)
	want := 5964

	testing_utils.CompareInts(t, got, want)
}
