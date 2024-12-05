package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

// https://adventofcode.com/2024/day/4
func main() {
	part1()
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	var matrix [][]byte
	for scanner.Scan() {
		matrix = append(matrix, scanner.Bytes())
	}

	var xmasAmt int
	for _, row := range matrix {
		xmasAmt += count(row)
	}

	colMatrix := make([][]byte, len(matrix))
	for i := range matrix {
		for j := range matrix[i] {
			colMatrix[j] = append(colMatrix[j], matrix[i][j])
		}
	}

	for _, column := range colMatrix {
		xmasAmt += count(column)
	}

	var diagonals [][]byte
	diagonals = append(diagonals, ccDiagonals(matrix)...)
	diagonals = append(diagonals, cDiagonals(matrix)...)
	for _, diagonal := range diagonals {
		xmasAmt += count(diagonal)
	}

	fmt.Println(xmasAmt)
}

func ccDiagonals(matrix [][]byte) [][]byte {
	var (
		diagonals [][]byte
		rowAmt    = len(matrix)
		colAmt    = len(matrix[0])
	)
	for k := 0; k < 2*colAmt-1; k++ {
		var list []byte
		for i := 0; i < rowAmt; i++ {
			for j := 0; j < colAmt; j++ {
				if i+j == k {
					list = append(list, matrix[i][j])
				}
			}
		}

		diagonals = append(diagonals, list)
	}

	return diagonals
}

func cDiagonals(matrix [][]byte) [][]byte {
	for i := range matrix {
		slices.Reverse(matrix[i])
	}

	var (
		diagonals [][]byte
		rowAmt    = len(matrix)
		colAmt    = len(matrix[0])
	)
	for k := 0; k < 2*colAmt-1; k++ {
		var list []byte
		for i := 0; i < rowAmt; i++ {
			for j := 0; j < colAmt; j++ {
				if i+j == k {
					list = append(list, matrix[i][j])
				}
			}
		}

		diagonals = append(diagonals, list)
	}

	return diagonals
}

func count(b []byte) int {
	count := bytes.Count(b, []byte("XMAS"))
	slices.Reverse(b)
	count += bytes.Count(b, []byte("XMAS"))

	return count
}
