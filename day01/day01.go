package day01

import (
	"aoc2024/utils"
	"sort"
	"strconv"
	"strings"
)

func Day01(lines []string, solve func(left, right []int) []int) int {
	left, right := createLeftAndRight(lines)
	numbers := solve(left, right)

	return utils.Sum(numbers)
}

func Part1(left, right []int) []int {
	distances := []int{}

	for i := range left {
		distance := utils.AbsoluteValue(left[i] - right[i])
		distances = append(distances, distance)
	}

	return distances
}

func Part2(left, right []int) []int {
	similarities := []int{}

	for _, lVal := range left {
		count := 0
		for _, rVal := range right {
			if lVal == rVal {
				count++
			}
		}
		similarities = append(similarities, lVal*count)
	}

	return similarities
}

func createLeftAndRight(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}

	for _, line := range lines {
		values := strings.Split(line, "   ")

		leftNum, _ := strconv.Atoi(values[0])
		rightNum, _ := strconv.Atoi(values[1])

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}
