package day03

import (
	"regexp"
	"strconv"
)

func Day03Part1(input string) int {
	sum := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		first, _ := strconv.Atoi(match[1])
		second, _ := strconv.Atoi(match[2])
		sum += first * second
	}

	return sum
}

func Day03Part2(input string) int {
	sum := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	skip := false

	for _, match := range matches {
		if match[0] == "don't()" {
			skip = true
		}

		if match[0] == "do()" {
			skip = false
		}

		if !skip && match[1] != "" && match[2] != "" {
			first, _ := strconv.Atoi(match[1])
			second, _ := strconv.Atoi(match[2])
			sum += first * second
		}
	}

	return sum
}
