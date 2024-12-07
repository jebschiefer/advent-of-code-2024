package day07

import (
	"aoc2024/testing_utils"
	"aoc2024/utils"
	"testing"
)

func TestDay07Part1Sample(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day07_sample.txt")

	got := Day07Part1(lines)
	want := 3749

	testing_utils.CompareInts(t, got, want)
}

func TestDay07Part1(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day07.txt")

	got := Day07Part1(lines)
	want := 4122618559853

	testing_utils.CompareInts(t, got, want)
}

func TestDay07Part2Sample(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day07_sample.txt")

	got := Day07Part2(lines)
	want := 11387

	testing_utils.CompareInts(t, got, want)
}

func TestDay07Part2(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day07.txt")

	got := Day07Part2(lines)
	want := 227615740238334

	testing_utils.CompareInts(t, got, want)
}
