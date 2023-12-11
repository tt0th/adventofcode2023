package main

import (
	"bufio"
	"fmt"
	. "github.com/tt0th/adventofcode2023/src/utils"
	"os"
	"slices"
)

func main() {
	//path := "src/day11/test-input-1.txt"
	path := "src/day11/input.txt"
	var universe = parseInput(path)

	universe = expandUniverse(universe)

	galaxies := collectGalaxies(universe)

	var sum int
	for _, from := range galaxies {
		for _, to := range galaxies {
			sum += from.DistanceTo(to)
		}
	}

	fmt.Printf("sum: %d\n", sum/2)
}

func expandUniverse(universe [][]rune) [][]rune {
	var emptyRows []int
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[0]); j++ {
			if universe[i][j] != '.' {
				break
			}
			if j == len(universe[0])-1 {
				emptyRows = append(emptyRows, i)
			}
		}
	}

	var emptyColumns []int
	for j := 0; j < len(universe[0]); j++ {
		for i := 0; i < len(universe); i++ {
			if universe[i][j] != '.' {
				break
			}
			if i == len(universe)-1 {
				emptyColumns = append(emptyColumns, j)
			}
		}
	}

	slices.Reverse(emptyRows)
	for _, rowIndex := range emptyRows {
		universe = append(universe[:rowIndex+1], universe[rowIndex:]...)
		universe[rowIndex] = universe[rowIndex+1]
	}

	slices.Reverse(emptyColumns)
	for _, columnIndex := range emptyColumns {
		for i := 0; i < len(universe); i++ {
			universe[i] = append(universe[i][:columnIndex+1], universe[i][columnIndex:]...)
			universe[i][columnIndex] = universe[i][columnIndex+1]
		}
	}

	return universe
}

func collectGalaxies(universe [][]rune) []C {
	var galaxies []C
	for i := 0; i < len(universe); i++ {
		for j := 0; j < len(universe[0]); j++ {
			if universe[i][j] == '#' {
				galaxies = append(galaxies, C{I: i, J: j})
			}
		}
	}
	return galaxies
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
