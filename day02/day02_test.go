package day02

import (
	"aoc2024/testing_utils"
	"aoc2024/utils"
	"testing"
)

func TestDay02Part1Sample(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day02_sample.txt")

	got := Day02(lines, IsReportSafe)
	want := 2

	testing_utils.CompareInts(t, got, want)
}

func TestDay02Part1(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day02.txt")

	got := Day02(lines, IsReportSafe)
	want := 670

	testing_utils.CompareInts(t, got, want)
}

func TestDay02Part2Sample(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day02_sample.txt")

	got := Day02(lines, IsReportSafeWithDampening)
	want := 5

	testing_utils.CompareInts(t, got, want)
}

func TestDay02Part2(t *testing.T) {
	lines := utils.ReadFileToLines("../inputs/day02.txt")

	got := Day02(lines, IsReportSafeWithDampening)
	want := 700

	testing_utils.CompareInts(t, got, want)
}
