package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetLines(input string) []string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}

func ReadFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func ReadFileToLines(path string) []string {
	content := ReadFile(path)
	return GetLines(content)
}

func AbsoluteValue(x int) int {
	if x < 0 {
		return 0 - x
	}

	return x
}

func Sum(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}

func StringToInts(value string) []int {
	numbers := []int{}
	values := strings.Split(value, " ")

	for _, num := range values {
		number, err := strconv.Atoi(num)

		if err == nil {
			numbers = append(numbers, number)
		}
	}

	return numbers
}
