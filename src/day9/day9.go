package main

import (
	"bufio"
	"fmt"
	"github.com/tt0th/adventofcode2023/src/utils"
	"os"
	"slices"
)

func main() {
	//path := "src/day9/test-input-1.txt"
	path := "src/day9/input.txt"
	var inputs = parseInput(path)

	var sum int
	var sum2 int
	for i := 0; i < len(inputs); i++ {
		sum += getNextNumber(inputs[i])
		slices.Reverse(inputs[i])
		sum2 += getNextNumber(inputs[i])
	}

	fmt.Printf("sum: %d, sum2: %d\n", sum, sum2)
}

func getNextNumber(numbers []int) int {
	if allNumbersAreZeros(numbers) {
		return 0
	}

	var diffs []int
	for i := 0; i < len(numbers)-1; i++ {
		diffs = append(diffs, numbers[i+1]-numbers[i])
	}
	lastDiff := getNextNumber(diffs)
	return numbers[len(numbers)-1] + lastDiff
}

func allNumbersAreZeros(numbers []int) bool {
	for _, number := range numbers {
		if number != 0 {
			return false
		}
	}
	return true
}

func parseInput(path string) [][]int {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var inputs [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, utils.StringToIntArray(line))
	}

	return inputs
}
