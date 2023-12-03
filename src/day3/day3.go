package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coordinate struct {
	i int
	j int
}

type PartNumber struct {
	number      int
	coordinates []Coordinate
}

func main() {
	//path := "src/day3/test-input-1.txt"
	path := "src/day3/input.txt"
	var matrix = parseMatrix(path)

	var partNumbers = collectPartNumbers(matrix)
	var validPartNumbers = filterForValidPartNumbers(partNumbers, matrix)

	var sum = 0
	for _, partNumber := range validPartNumbers {
		sum += partNumber.number
	}
	var sum2 = calculateGearStuff(partNumbers, matrix)
	fmt.Printf("Sum: %d, sum2: %d\n", sum, sum2)
}

func calculateGearStuff(partNumbers []PartNumber, matrix [][]rune) int {
	var sum int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '*' {
				var adjacentPartNumbers = getAdjacentPartNumbers(i, j, partNumbers)
				if len(adjacentPartNumbers) == 2 {
					sum += adjacentPartNumbers[0].number * adjacentPartNumbers[1].number
				}
			}
		}
	}
	return sum
}

func getAdjacentPartNumbers(i int, j int, partNumbers []PartNumber) []PartNumber {
	var adjacentPartNumbers []PartNumber
	for _, partNumber := range partNumbers {
		if isPartNumberAdjacent(partNumber, i, j) {
			adjacentPartNumbers = append(adjacentPartNumbers, partNumber)
		}
	}
	return adjacentPartNumbers
}

func filterForValidPartNumbers(partNumbers []PartNumber, matrix [][]rune) []PartNumber {
	var validPartNumbers []PartNumber
	for _, partNumber := range partNumbers {
		if isPartNumberValid(partNumber, matrix) {
			validPartNumbers = append(validPartNumbers, partNumber)
		}
	}
	return validPartNumbers
}

func isPartNumberValid(partNumber PartNumber, matrix [][]rune) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !isRuneDigit(matrix[i][j]) && matrix[i][j] != '.' {
				if isPartNumberAdjacent(partNumber, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func isPartNumberAdjacent(partNumber PartNumber, i int, j int) bool {
	for _, coordinate := range partNumber.coordinates {
		if areCoordinatesAdjacent(coordinate, Coordinate{i, j}) {
			return true
		}
	}
	return false
}

func areCoordinatesAdjacent(coordinate1 Coordinate, coordinate2 Coordinate) bool {
	return Abs(coordinate1.i-coordinate2.i) < 2 && Abs(coordinate1.j-coordinate2.j) < 2
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func collectPartNumbers(matrix [][]rune) []PartNumber {
	var partNumbers []PartNumber
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if isRuneDigit(matrix[i][j]) {
				var coordinates []Coordinate
				var runes []rune
				for ; j < len(matrix[i]) && isRuneDigit(matrix[i][j]); j++ {
					coordinates = append(coordinates, Coordinate{i, j})
					runes = append(runes, matrix[i][j])
				}
				number, _ := strconv.Atoi(string(runes))
				partNumbers = append(partNumbers, PartNumber{number: number, coordinates: coordinates})
			}
		}
	}
	return partNumbers
}

func isRuneDigit(character rune) bool {
	return character >= '0' && character <= '9'
}

func parseMatrix(path string) [][]rune {
	var matrix [][]rune
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	return matrix
}
