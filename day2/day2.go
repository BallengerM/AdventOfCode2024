package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/2
func main() {
	//part1()
	part2()
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	var safeCount int
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), " ")

		var (
			previousNum            int
			increasing, decreasing bool
			safe                   = true
		)
	looooooop:
		for i, current := range nums {
			converted, err := strconv.Atoi(current)
			if err != nil {
				panic(err)
			}

			if i == 0 {
				previousNum = converted
				continue
			}

			diff := float64(previousNum - converted)
			previousNum = converted

			if diff == 0 {
				safe = false
				break
			}

			switch {
			// if we are increasing, we expect the diff to be a negative number
			case increasing:
				// if diff is negative, signbit will return true
				if !math.Signbit(diff) {
					safe = false
					break looooooop
				}
			// if we are decreasing, we expect the diff to be a positive number
			case decreasing:
				// if diff is positive, signbit will return false
				if math.Signbit(diff) {
					safe = false
					break looooooop
				}
			// we don't know yet, so set accordingly
			default:
				if math.Signbit(diff) {
					increasing = true
				} else {
					decreasing = true
				}
			}

			if math.Abs(diff) > 3 {
				safe = false
				break looooooop
			}
		}

		if safe {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	var safeCount int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		// Check if the level is safe as is
		if checkLevel(nums) {
			safeCount++
			continue
		}

		for i := 0; i < len(nums); i++ {
			if checkLevel(remove(nums, i)) {
				safeCount++
				break
			}
		}
	}

	fmt.Println(safeCount)
}

func remove(slice []string, index int) []string {
	temp := make([]string, len(slice))
	copy(temp, slice)
	return append(temp[:index], temp[index+1:]...)
}

func checkLevel(nums []string) bool {
	var (
		previousNum            int
		increasing, decreasing bool
	)
	for i, current := range nums {
		converted, err := strconv.Atoi(current)
		if err != nil {
			panic(err)
		}

		if i == 0 {
			previousNum = converted
			continue
		}

		var safeDiff bool
		safeDiff, increasing, decreasing = safe(converted, previousNum, increasing, decreasing)
		previousNum = converted

		if !safeDiff {
			return false
		}
	}
	return true
}

func safe(current, previous int, increasing, decreasing bool) (bool, bool, bool) {
	diff := float64(previous - current)

	if diff == 0 {
		return false, increasing, decreasing
	}

	switch {
	case increasing:
		if !math.Signbit(diff) {
			return false, increasing, decreasing
		}
	case decreasing:
		if math.Signbit(diff) {
			return false, increasing, decreasing
		}
	default:
		if math.Signbit(diff) {
			increasing = true
		} else {
			decreasing = true
		}
	}

	if math.Abs(diff) > 3 {
		return false, increasing, decreasing
	}

	return true, increasing, decreasing
}
