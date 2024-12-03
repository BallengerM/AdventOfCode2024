package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/3
func main() {
	//part1()
	part2()
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	var overallTotal int
	for scanner.Scan() {
		line := scanner.Text()

		r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
		if err != nil {
			panic(err)
		}

		matches := r.FindAllString(line, -1)

		for _, command := range matches {
			command = strings.TrimPrefix(command, "mul(")
			command = strings.TrimSuffix(command, ")")

			nums := strings.Split(command, ",")

			num1, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}

			num2, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}

			overallTotal += num1 * num2
		}
	}
	fmt.Println(overallTotal)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	var full string
	for scanner.Scan() {
		line := scanner.Text()

		full += line
	}

	r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	if err != nil {
		panic(err)
	}

	matches := r.FindAllString(full, -1)

	var (
		process      = true // Will be true until finding a 'don't'
		overallTotal int
	)
	for i := range matches {
		switch matches[i] {
		case "don't()":
			process = false
		case "do()":
			process = true
		default:
			if process {
				overallTotal += commandResult(matches[i])
			}
		}
	}

	fmt.Println(overallTotal)
}

func commandResult(command string) int {
	command = strings.TrimPrefix(command, "mul(")
	command = strings.TrimSuffix(command, ")")

	nums := strings.Split(command, ",")

	num1, err := strconv.Atoi(nums[0])
	if err != nil {
		panic(err)
	}

	num2, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}

	return num1 * num2
}
