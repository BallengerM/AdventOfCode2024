package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1
func main() {
	//part1()
	part2()
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		left  []int
		right []int
	)
	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Split(line, "   ")

		// I'm assuming that we are guaranteed to have two numbers
		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left = append(left, leftNum)

		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right = append(right, rightNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	var totalDistance int
	for i := range left {
		distance := math.Abs(float64(left[i] - right[i]))
		totalDistance += int(distance)
	}

	fmt.Println(totalDistance)
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		left  []int
		right []int
	)
	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Split(line, "   ")

		// I'm assuming that we are guaranteed to have two numbers
		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left = append(left, leftNum)

		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right = append(right, rightNum)
	}

	var totalScore int
	for _, value := range left {
		similarity := count(value, right)

		totalScore += value * similarity
	}

	fmt.Println(totalScore)
}

func count(x int, s []int) int {
	var count int
	for _, value := range s {
		if value == x {
			count++
		}
	}

	return count
}
