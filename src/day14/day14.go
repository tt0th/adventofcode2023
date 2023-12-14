package main

import (
	"bufio"
	"fmt"
	"os"
)

const ROCK = 'O'
const EMPTY = '.'

func main() {
	//path := "src/day14/test-input-1.txt"
	path := "src/day14/input.txt"
	var platform = parseInput(path)

	rollToNorth(platform)

	var sum int
	for i, row := range platform {
		for _, column := range row {
			if column == ROCK {
				sum += len(platform) - i
			}
		}
	}

	fmt.Printf("sum: %d\n", sum)
}

func rollToNorth(platform [][]rune) {
	for range platform {
		for i := 0; i < len(platform)-1; i++ {
			for j := 0; j < len(platform[0]); j++ {
				if platform[i][j] == EMPTY && platform[i+1][j] == ROCK {
					platform[i+1][j] = EMPTY
					platform[i][j] = ROCK
				}
			}
		}
	}
}

func parseInput(path string) [][]rune {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var inputs [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, []rune(line))
	}

	return inputs
}
