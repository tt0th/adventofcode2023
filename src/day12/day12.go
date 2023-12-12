package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	. "github.com/tt0th/adventofcode2023/src/utils"
	"os"
	"regexp"
	"strings"
)

const DAMAGED rune = '#'
const OPERATIONAL rune = '.'
const UNKNOWN rune = '?'

type Row struct {
	springs        []rune
	damagedLengths []int
}

func main() {
	//path := "src/day12/test-input-1.txt"
	path := "src/day12/input.txt"
	var rows = parseInput(path)

	var solutions [][][]rune
	solutions = findSolutionsForRows(rows)

	var sum int
	for _, solution := range solutions {
		sum += len(solution)
	}

	fmt.Printf("sum: %d\n", sum)
}

func findSolutionsForRows(rows []Row) [][][]rune {
	var solutions [][][]rune
	for _, row := range rows {
		solution := findSolutionsForRow(row)
		solutions = append(solutions, solution)
	}
	return solutions
}

func findSolutionsForRow(row Row) [][]rune {
	if isRowSolved(row) {
		return [][]rune{row.springs}
	}
	if isRowUnsolvable(row) {
		return [][]rune{}
	}

	rowCopy1 := row
	rowCopy2 := row
	rowCopy1.springs = lo.Replace(row.springs, UNKNOWN, DAMAGED, 1)
	rowCopy2.springs = lo.Replace(row.springs, UNKNOWN, OPERATIONAL, 1)
	return append(findSolutionsForRow(rowCopy1), findSolutionsForRow(rowCopy2)...)
}

func isRowUnsolvable(row Row) bool {
	damagedCount := lo.Count(row.springs, DAMAGED)
	unknownCount := lo.Count(row.springs, UNKNOWN)
	maxPossibleDamagedCount := damagedCount + unknownCount
	expectedDamagedCount := lo.Sum(row.damagedLengths)

	return damagedCount > expectedDamagedCount || maxPossibleDamagedCount < expectedDamagedCount || unknownCount == 0
}

func isRowSolved(row Row) bool {
	unknownCount := lo.Count(row.springs, UNKNOWN)
	if unknownCount > 0 {
		return false
	}

	m1 := regexp.MustCompile(`\.+`)
	stringOutcome := strings.Trim(m1.ReplaceAllString(string(row.springs), "."), ".")
	expectedStringOutcome := strings.Join(lo.Map(row.damagedLengths, func(length int, _ int) string {
		return strings.Repeat("#", length)
	}), ".")
	return stringOutcome == expectedStringOutcome
}

func parseInput(path string) []Row {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var inputs []Row

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		input := Row{
			springs:        []rune(parts[0]),
			damagedLengths: StringToIntArrayBy(parts[1], ","),
		}
		inputs = append(inputs, input)
	}

	return inputs
}
