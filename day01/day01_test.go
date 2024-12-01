package day01

import (
	"aoc2024/testing_utils"
	"aoc2024/utils"
	"testing"
)

func TestDay01Part1Sample(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day01_sample1.txt")

	got := Day01(lines, Part1)
	expected := 11
	testing_utils.CompareInts(t, got, expected)
}

func TestDay01Part1(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day01_1.txt")

	got := Day01(lines, Part1)
	expected := 1590491
	testing_utils.CompareInts(t, got, expected)
}

func TestDay01Part2Sample(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day01_sample1.txt")

	got := Day01(lines, Part2)
	expected := 31
	testing_utils.CompareInts(t, got, expected)
}

func TestDay01Part2(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day01_1.txt")

	got := Day01(lines, Part2)
	expected := 22588371
	testing_utils.CompareInts(t, got, expected)
}
