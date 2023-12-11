package main

import (
	"bufio"
	"fmt"
	. "github.com/tt0th/adventofcode2023/src/utils"
	"os"
)

func main() {
	//path := "src/day11/test-input-1.txt"
	path := "src/day11/input.txt"
	var universe = parseInput(path)

	galaxies := collectGalaxies(universe)
	galaxies = translateGalaxiesWithExpansion(universe, galaxies)

	var sum int
	for _, from := range galaxies {
		for _, to := range galaxies {
			sum += from.DistanceTo(to)
		}
	}

	fmt.Printf("sum: %d\n", sum/2)
}

func translateGalaxiesWithExpansion(universe [][]rune, galaxies []C) []C {
	expansionRate := 999999
	var translatedGalaxies []C
	emptyRows := collectEmptyRows(universe)
	emptyColumns := collectEmptyColumns(universe)
	for _, galaxy := range galaxies {
		galaxyCopy := galaxy
		for _, rowIndex := range emptyRows {
			if rowIndex < galaxy.I {
				galaxyCopy.I += expansionRate
			}
		}
		for _, columnsIndex := range emptyColumns {
			if columnsIndex < galaxy.J {
				galaxyCopy.J += expansionRate
			}
		}
		translatedGalaxies = append(translatedGalaxies, galaxyCopy)
	}
	return translatedGalaxies
}

func collectEmptyColumns(universe [][]rune) []int {
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
	return emptyColumns
}

func collectEmptyRows(universe [][]rune) []int {
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
	return emptyRows
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
