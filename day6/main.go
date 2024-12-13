package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/6
// UNSOLVED
func main() {
	part1()
}

type location struct {
	row int
	col int
}

// ^ guard
// # obstacle
// guard moves forward one step unless obstacle, then right 90 degrees
// determine how many positions visited before leaving
func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		matrix        []string
		i             int
		guardLocation location
	)
	for scanner.Scan() {
		line := scanner.Text()

		guardIndex := strings.Index(line, "^")
		if guardIndex != -1 {
			guardLocation = location{
				row: i,
				col: guardIndex,
			}
		}

		i++
		matrix = append(matrix, line)
	}
	copiedMatrix := make([]string, len(matrix))
	copy(copiedMatrix, matrix)

	direction := "UP"
	for {
		var nextPos location
		switch direction {
		case "UP":
			nextPos.row = guardLocation.row - 1
			nextPos.col = guardLocation.col
		case "DOWN":
			nextPos.row = guardLocation.row + 1
			nextPos.col = guardLocation.col
		case "LEFT":
			nextPos.row = guardLocation.row
			nextPos.col = guardLocation.col - 1
		case "RIGHT":
			nextPos.row = guardLocation.row
			nextPos.col = guardLocation.col + 1
		}

		// We have made it to our last position
		if nextPos.row > len(matrix)-1 || nextPos.col > len(matrix[0])-1 || nextPos.row < 0 || nextPos.col < 0 {
			break
		}

		// We need to turn right
		if matrix[nextPos.row][nextPos.col] == byte('#') {
			switch direction {
			case "UP":
				direction = "RIGHT"
			case "DOWN":
				direction = "LEFT"
			case "LEFT":
				direction = "UP"
			case "RIGHT":
				direction = "DOWN"
			}
			continue
		}

		guardLocation = nextPos

		line := copiedMatrix[guardLocation.row]
		replaced := line[:guardLocation.col] + "X" + line[guardLocation.col+1:]
		copiedMatrix[guardLocation.row] = replaced
	}

	var amtOfSteps int
	for i := range copiedMatrix {
		fmt.Println(string(copiedMatrix[i]))
		amtOfSteps += strings.Count(copiedMatrix[i], "X")
	}

	fmt.Println(amtOfSteps)
}
