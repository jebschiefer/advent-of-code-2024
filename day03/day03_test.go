package day03

import (
	"aoc2024/testing_utils"
	"aoc2024/utils"
	"testing"
)

func TestDay02Part1Sample(t *testing.T) {
	input := utils.ReadFile("../inputs/day03_sample.txt")

	got := Day03(input)
	want := 161

	testing_utils.CompareInts(t, got, want)
}

func TestDay02Part1(t *testing.T) {
	input := utils.ReadFile("../inputs/day03.txt")

	got := Day03(input)
	want := 166357705

	testing_utils.CompareInts(t, got, want)
}
