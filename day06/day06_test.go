package day06

import (
	"aoc2024/testing_utils"
	"aoc2024/utils"
	"testing"
)

func TestDay06Part1Sample(t *testing.T) {
	grid := utils.ReadFileToGrid("../inputs/day06_sample.txt")

	got := Day06(grid)
	want := 41

	testing_utils.CompareInts(t, got, want)
}

func TestDay06Part1(t *testing.T) {
	grid := utils.ReadFileToGrid("../inputs/day06.txt")

	got := Day06(grid)
	want := 4988

	testing_utils.CompareInts(t, got, want)
}
