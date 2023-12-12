package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	. "github.com/tt0th/adventofcode2023/src/utils"
	"os"
	"strings"
	"time"
)

const DAMAGED rune = '#'
const OPERATIONAL rune = '.'
const UNKNOWN rune = '?'

type Row struct {
	springs              []rune
	damagedLengths       []int
	expectedDamagedCount int
}

func main() {
	//path := "src/day12/test-input-1.txt"
	path := "src/day12/input.txt"
	var rows = parseInput(path)

	start := time.Now()

	var sum int
	for _, solutionCount := range findSolutionsForRows(rows) {
		sum += solutionCount
	}
	var sum2 int
	for _, solutionCount := range findSolutionsForRows(unfoldSprings(rows)) {
		sum2 += solutionCount
	}

	fmt.Printf("Time %s\n", time.Since(start))

	fmt.Printf("sum: %d, sum2: %d\n", sum, sum2)
}

func unfoldSprings(rows []Row) []Row {
	return lo.Map(rows, func(row Row, _ int) Row {
		return Row{
			springs:              lo.Flatten([][]rune{row.springs, {UNKNOWN}, row.springs, {UNKNOWN}, row.springs, {UNKNOWN}, row.springs, {UNKNOWN}, row.springs}),
			damagedLengths:       lo.Flatten([][]int{row.damagedLengths, row.damagedLengths, row.damagedLengths, row.damagedLengths, row.damagedLengths}),
			expectedDamagedCount: row.expectedDamagedCount * 5,
		}
	})
}

func findSolutionsForRows(rows []Row) []int {
	return lop.Map(rows, func(row Row, index int) int {
		solutionCount := findSolutionsForRowByPartitioning(row)
		println("solutions for #", index, string(row.springs), solutionCount)
		return solutionCount
	})
}

func findSolutionsForRowByPartitioning(row Row) int {
	indexOfFirstOperational := lo.IndexOf(row.springs, OPERATIONAL)
	if indexOfFirstOperational == -1 {
		return findSolutionsForRow(row)
	}
	firstPart := row.springs[0:indexOfFirstOperational]
	otherPart := row.springs[indexOfFirstOperational+1:]

	var solutionCounter int
	for i := 0; i < len(row.damagedLengths)+1; i++ {
		row1 := Row{
			springs:              firstPart,
			damagedLengths:       row.damagedLengths[0:i],
			expectedDamagedCount: lo.Sum(row.damagedLengths[0:i]),
		}
		row2 := Row{
			springs:              otherPart,
			damagedLengths:       row.damagedLengths[i:],
			expectedDamagedCount: lo.Sum(row.damagedLengths[i:]),
		}
		firstPartSolution := findSolutionsForRowByPartitioning(row1)
		if firstPartSolution > 0 {
			solutionCounter += firstPartSolution * findSolutionsForRowByPartitioning(row2)
		}
	}
	return solutionCounter
}

func findSolutionsForRow(row Row) int {
	if isRowSolved(row) {
		return 1
	}
	if isRowUnsolvable(row) {
		return 0
	}

	unknownIndexes := lo.FilterMap(row.springs, func(item rune, index int) (int, bool) {
		return index, item == UNKNOWN
	})
	indexToChange := unknownIndexes[len(unknownIndexes)/2]
	springs1 := append([]rune{}, row.springs...)
	springs2 := append([]rune{}, row.springs...)
	springs1[indexToChange] = DAMAGED
	springs2[indexToChange] = OPERATIONAL
	rowCopy1 := Row{
		springs:              springs1,
		damagedLengths:       row.damagedLengths,
		expectedDamagedCount: row.expectedDamagedCount,
	}
	rowCopy2 := Row{
		springs:              springs2,
		damagedLengths:       row.damagedLengths,
		expectedDamagedCount: row.expectedDamagedCount,
	}

	return findSolutionsForRowByPartitioning(rowCopy1) + findSolutionsForRowByPartitioning(rowCopy2)
}

func isRowUnsolvable(row Row) bool {
	unknownCount := lo.Count(row.springs, UNKNOWN)
	if unknownCount == 0 {
		return true
	}

	damagedCount := lo.Count(row.springs, DAMAGED)
	if damagedCount > row.expectedDamagedCount {
		return true
	}

	maxPossibleDamagedCount := damagedCount + unknownCount
	if maxPossibleDamagedCount < row.expectedDamagedCount {
		return true
	}

	lengths := collectGroupLengthsBeforeUnknown(row)
	for i := 0; i < len(lengths); i++ {
		if lengths[i] != row.damagedLengths[i] {
			return true
		}
	}

	return false
}

func collectGroupLengthsBeforeUnknown(row Row) []int {
	state := "waitingForDamaged"
	var length int
	var lengths []int
	for i := 0; i < len(row.springs); i++ {
		if row.springs[i] == UNKNOWN {
			return lengths
		}
		switch state {
		case "waitingForDamaged":
			if row.springs[i] == DAMAGED {
				state = "collecting"
				length = 1
			}
		case "collecting":
			if row.springs[i] == DAMAGED {
				length++
			} else {
				state = "waitingForDamaged"
				lengths = append(lengths, length)
			}
		}
	}
	if state == "collecting" {
		lengths = append(lengths, length)
	}

	return lengths
}

func isRowSolved(row Row) bool {
	if lo.IndexOf(row.springs, UNKNOWN) != -1 {
		return false
	}

	lengths := collectGroupLengthsBeforeUnknown(row)
	if len(lengths) != len(row.damagedLengths) {
		return false
	}
	for i := 0; i < len(lengths); i++ {
		if lengths[i] != row.damagedLengths[i] {
			return false
		}
	}
	return true
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
		damagedLengths := StringToIntArrayBy(parts[1], ",")
		input := Row{
			springs:              []rune(parts[0]),
			damagedLengths:       damagedLengths,
			expectedDamagedCount: lo.Sum(damagedLengths),
		}
		inputs = append(inputs, input)
	}

	return inputs
}
