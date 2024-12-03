package day03

import (
	"regexp"
	"strconv"
)

func Day03(input string) int {
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
