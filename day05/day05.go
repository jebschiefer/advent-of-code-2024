package day05

import (
	"aoc2024/utils"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Rules = map[int][]int
type Updates = [][]int
type PartFunc = func(rules Rules, updates Updates) Updates

func Day05(input string, part PartFunc) int {
	parts := strings.Split(input, "\n\n")

	rules := getRules(parts[0])
	updates := getUpdates(parts[1])
	updatesPart := part(rules, updates)
	middlePages := getMiddlePages(updatesPart)

	return utils.Sum(middlePages)
}

func Part1(rules Rules, updates Updates) Updates {
	inOrderUpdates, _ := categorizeUpdates(rules, updates)
	return inOrderUpdates
}

func Part2(rules Rules, updates Updates) Updates {
	_, wrongOrderUpdates := categorizeUpdates(rules, updates)
	wrongOrderUpdates = sortUpdates(rules, wrongOrderUpdates)

	return wrongOrderUpdates
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

func categorizeUpdates(rules Rules, updates Updates) (Updates, Updates) {
	inOrderUpdates := Updates{}
	wrongOrderUpdates := Updates{}

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
		} else {
			wrongOrderUpdates = append(wrongOrderUpdates, pages)
		}
	}

	return inOrderUpdates, wrongOrderUpdates
}

func getMiddlePages(updates Updates) []int {
	middlePages := []int{}

	for _, update := range updates {
		middle := len(update) / 2
		middlePages = append(middlePages, update[middle])
	}

	return middlePages
}

func sortUpdates(rules Rules, updates Updates) Updates {
	sorted := Updates{}

	for _, pages := range updates {
		sort.Slice(pages, func(i, j int) bool {
			found := false
			rule, ok := rules[pages[i]]

			if ok {
				for _, r := range rule {
					if r == pages[j] {
						found = true
					}
				}
			}

			return found
		})

		sorted = append(sorted, pages)
	}

	return sorted
}
