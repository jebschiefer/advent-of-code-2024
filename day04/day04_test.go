package day04

import (
	"aoc2024/testing_utils"
	"aoc2024/utils"
	"testing"
)

func TestDay04Part1Sample(t *testing.T) {
	grid := utils.ReadFileToGrid("../inputs/day04_sample.txt")

	got := Day04(grid)
	want := 18

	testing_utils.CompareInts(t, got, want)
}

func TestDay04Part1(t *testing.T) {
	grid := utils.ReadFileToGrid("../inputs/day04.txt")

	got := Day04(grid)
	want := 2500

	testing_utils.CompareInts(t, got, want)
}
