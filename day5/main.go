package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/5
// Part 2 UNSOLVED
func main() {
	// part1()
	part2()
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		orderingRules []string
		pageNumbers   []string
	)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.Contains(line, "|"):
			orderingRules = append(orderingRules, line)
		case strings.Contains(line, ","):
			pageNumbers = append(pageNumbers, line)
		}
	}

	ruleMap := make(map[string][]string) // Map contains keys, then everything that must come after that key
	for _, rule := range orderingRules {
		split := strings.Split(rule, "|")
		ruleMap[split[0]] = append(ruleMap[split[0]], split[1])
	}

	var goodSets []string
	for _, row := range pageNumbers {
		split := strings.Split(row, ",")

		var failed bool
		for _, num := range split {
			numRuleSet := ruleMap[num]
			numIndex := strings.Index(row, num)

			// Each of the values in this set must be _after_ the current num's index
		loooop:
			for _, v := range numRuleSet {
				vIndex := strings.Index(row, v)
				switch {
				case vIndex == -1:
					// Do nothing, this means that the number isn't in the row
				case vIndex < numIndex:
					failed = true
					break loooop
				}
			}
			if failed {
				break
			}
		}

		if !failed {
			goodSets = append(goodSets, row)
		}
	}

	var sum int
	for i := range goodSets {
		split := strings.Split(goodSets[i], ",")
		middleIndex := int(len(split) / 2)

		numInt, _ := strconv.Atoi(split[middleIndex])
		sum += numInt
	}

	fmt.Println(sum)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		orderingRules []string
		pageNumbers   []string
	)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.Contains(line, "|"):
			orderingRules = append(orderingRules, line)
		case strings.Contains(line, ","):
			pageNumbers = append(pageNumbers, line)
		}
	}

	ruleMap := make(map[string][]string) // Map contains keys, then everything that must come after that key
	for _, rule := range orderingRules {
		split := strings.Split(rule, "|")
		ruleMap[split[0]] = append(ruleMap[split[0]], split[1])
	}

	var failedSets []string
	for _, row := range pageNumbers {
		split := strings.Split(row, ",")

		var failed bool
		for _, num := range split {
			numRuleSet := ruleMap[num]
			numIndex := strings.Index(row, num)

			// Each of the values in this set must be _after_ the current num's index
		loooop:
			for _, v := range numRuleSet {
				vIndex := strings.Index(row, v)
				switch {
				case vIndex == -1:
					// Do nothing, this means that the number isn't in the row
				case vIndex < numIndex:
					failed = true
					break loooop
				}
			}
			if failed {
				break
			}
		}

		if failed {
			failedSets = append(failedSets, row)
		}
	}

	var corrected []string
	for i := range failedSets {
		start := strings.Split(failedSets[i], ",")

		var sorted bool
		for !sorted {
			start = sort(start)
			sorted = check(start, ruleMap)
		}
		corrected = append(corrected, strings.Join(start, ","))
	}

	var sum int
	for i := range corrected {
		split := strings.Split(corrected[i], ",")
		middleIndex := int(len(split) / 2)

		numInt, _ := strconv.Atoi(split[middleIndex])
		sum += numInt
	}

	fmt.Println(sum)
}

func sort(ss []string) []string {
	// TODO: Write proper sorting algorithm
	return ss
}

func check(ss []string, ruleMap map[string][]string) bool {
	row := strings.Join(ss, ",")

	var failed bool
	for _, num := range ss {
		numRuleSet := ruleMap[num]
		numIndex := strings.Index(row, num)

		// Each of the values in this set must be _after_ the current num's index
	loooop:
		for _, v := range numRuleSet {
			vIndex := strings.Index(row, v)
			switch {
			case vIndex == -1:
				// Do nothing, this means that the number isn't in the row
			case vIndex < numIndex:
				failed = true
				break loooop
			}
		}
		if failed {
			break
		}
	}

	return !failed // Will return true if we did not fail
}
