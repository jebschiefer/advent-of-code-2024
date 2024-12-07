package day07

import (
	"strconv"
	"strings"
)

func Day07Part1(lines []string) int {
	operators := []string{"+", "*"}
	return sumTrueEquations(lines, operators)
}

func Day07Part2(lines []string) int {
	operators := []string{"+", "*", "||"}
	return sumTrueEquations(lines, operators)
}

func getValues(line string) (int, []int) {
	nums := []int{}

	first := strings.Split(line, ":")
	firstNum, _ := strconv.Atoi(first[0])

	second := strings.Split(strings.Trim(first[1], " "), " ")
	for _, value := range second {
		num, _ := strconv.Atoi(value)
		nums = append(nums, num)
	}

	return firstNum, nums
}

func sumTrueEquations(lines []string, operators []string) int {
	total := 0
	for _, line := range lines {
		testValue, numbers := getValues(line)
		end := len(numbers) - 1

		if isTrueEquation(testValue, numbers, operators, "", 0, end, numbers[0]) {
			total += testValue
		}
	}

	return total
}

func isTrueEquation(testValue int, numbers []int, operators []string, operator string, currentPosition int, end int, total int) bool {
	total = doOperation(total, numbers[currentPosition], operator)

	if currentPosition == end {
		return testValue == total
	}

	currentPosition++

	for _, operator := range operators {
		if isTrueEquation(testValue, numbers, operators, operator, currentPosition, end, total) {
			return true
		}
	}

	return false
}

func doOperation(firstNumber, secondNumber int, operator string) int {
	result := firstNumber

	if operator == "+" {
		result += secondNumber
	} else if operator == "*" {
		result *= secondNumber
	} else if operator == "||" {
		resultString := strconv.Itoa(result)
		secondNumberString := strconv.Itoa(secondNumber)
		result, _ = strconv.Atoi(resultString + secondNumberString)
	}

	return result
}
