package day02

import (
	"aoc2024/utils"
	"slices"
)

func Day02(lines []string, solve func([]int) bool) int {
	safeCount := 0
	reports := getReports(lines)

	for _, report := range reports {
		if solve(report) {
			safeCount++
		}
	}

	return safeCount
}

func getReports(lines []string) [][]int {
	reports := [][]int{}

	for _, line := range lines {
		report := utils.StringToInts(line)
		reports = append(reports, report)
	}

	return reports
}

func IsReportSafe(levels []int) bool {
	safe := false

	levelChanges := []int{}

	for i, currentLevel := range levels {
		if i > 0 {
			previousLevel := levels[i-1]
			change := currentLevel - previousLevel
			levelChanges = append(levelChanges, change)
		}
	}

	allIncreasing, allDecreasing, allAllowedStepSize := true, true, true

	for _, change := range levelChanges {
		absChange := utils.AbsoluteValue(change)

		if absChange == 0 || absChange > 3 {
			allAllowedStepSize = false
		}

		allIncreasing = allIncreasing && change > 0
		allDecreasing = allDecreasing && change < 0
	}

	safe = (allIncreasing || allDecreasing) && allAllowedStepSize

	return safe
}

func IsReportSafeWithDampening(levels []int) bool {
	safe := IsReportSafe(levels)

	if !safe {
		for i := range levels {
			dampenedLevels := []int{}
			dampenedLevels = append(dampenedLevels, levels...)
			dampenedLevels = slices.Delete(dampenedLevels, i, i+1)

			safe = IsReportSafe(dampenedLevels)

			if safe {
				break
			}
		}
	}

	return safe
}
