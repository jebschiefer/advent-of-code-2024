package day05

import (
	"aoc2024/utils"
	"regexp"
	"strconv"
	"strings"
)

type Rules = map[int][]int
type Updates = [][]int

func Day05(input string) int {
	parts := strings.Split(input, "\n\n")

	rules := getRules(parts[0])
	updates := getUpdates(parts[1])
	inOrderUpdates := findInOrderUpdates(rules, updates)
	middlePages := getMiddlePages(inOrderUpdates)

	return utils.Sum(middlePages)
}

func getRules(input string) Rules {
	rules := Rules{}

	re := regexp.MustCompile(`(\d+)\|(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		before, _ := strconv.Atoi(match[1])
		after, _ := strconv.Atoi(match[2])

		entry, ok := rules[before]

		if !ok {
			entry = []int{}
		}

		entry = append(entry, after)
		rules[before] = entry
	}

	return rules
}

func getUpdates(input string) Updates {
	updates := Updates{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pages := strings.Split(line, ",")
		update := []int{}

		for _, page := range pages {
			number, _ := strconv.Atoi(page)
			update = append(update, number)
		}

		updates = append(updates, update)
	}

	return updates
}

func findInOrderUpdates(rules Rules, updates Updates) Updates {
	inOrderUpdates := Updates{}

	for _, pages := range updates {
		inOrder := true
		for i, page := range pages {
			if i == len(pages)-1 {
				break
			}

			rule, ok := rules[page]

			if ok {
				for _, p2 := range pages[i+1:] {
					found := false
					for _, r := range rule {
						if p2 == r {
							found = true
							break
						}
					}
					if !found {
						inOrder = false
					}
				}

			} else {
				inOrder = false
			}
		}

		if inOrder {
			inOrderUpdates = append(inOrderUpdates, pages)
		}
	}

	return inOrderUpdates
}

func getMiddlePages(updates Updates) []int {
	middlePages := []int{}

	for _, update := range updates {
		middle := len(update) / 2
		middlePages = append(middlePages, update[middle])
	}

	return middlePages
}
