package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"
)

func main() {
	//path := "src/day15/test-input-1.txt"
	path := "src/day15/input.txt"
	var inputs = parseInput(path)

	sum := lo.SumBy(inputs, func(item string) int {
		return hash(item)
	})

	fmt.Printf("sum: %d\n", sum)
}

func hash(input string) int {
	value := 0
	for _, char := range []rune(input) {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}

func parseInput(path string) []string {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var input string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = line
	}

	return strings.Split(input, ",")
}
