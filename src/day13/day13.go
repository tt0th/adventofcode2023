package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"time"
)

type MirrorDirection int

const (
	Horizontal MirrorDirection = iota
	Vertical
)

type Mirror struct {
	direction MirrorDirection
	position  int
}

type Field = [][]rune

const ASH = '.'
const ROCK = '#'

func main() {
	//path := "src/day13/test-input-1.txt"
	path := "src/day13/input.txt"
	var fields = parseInput(path)

	start := time.Now()

	var sum int
	for _, mirror := range findMirrors(fields) {
		if mirror.direction == Vertical {
			sum += mirror.position
		} else {
			sum += mirror.position * 100
		}
	}
	fmt.Printf("Time %s\n", time.Since(start))
	fmt.Printf("Sum: %d\n", sum)
}

func printField(field Field) {
	for _, runes := range field {
		println(string(runes))
	}
}

func findMirrors(fields []Field) []Mirror {
	return lo.Map(fields, func(field Field, _ int) Mirror {
		return findMirror(field)
	})
}

func findMirror(field Field) Mirror {
	// horizontal
	for position := 1; position < len(field); position++ {
		if isHorizontalMirrorCorrect(field, position) {
			return Mirror{
				direction: Horizontal,
				position:  position,
			}
		}
	}
	// vertical
	for position := 1; position < len(field[0]); position++ {
		if isVerticalMirrorCorrect(field, position) {
			return Mirror{
				direction: Vertical,
				position:  position,
			}
		}
	}
	printField(field)
	panic("No mirror o.O")
}

func isVerticalMirrorCorrect(field Field, position int) bool {
	sizeOfReflection := min(position, len(field[0])-position)
	for i := 0; i < sizeOfReflection; i++ {
		originalColumn := position - i - 1
		reflectedColumn := position + i
		for row := 0; row < len(field); row++ {
			if field[row][originalColumn] != field[row][reflectedColumn] {
				return false
			}
		}
	}
	return true
}

func isHorizontalMirrorCorrect(field Field, position int) bool {
	sizeOfReflection := min(position, len(field)-position)
	for i := 0; i < sizeOfReflection; i++ {
		originalRow := position - i - 1
		reflectedRow := position + i
		for column := 0; column < len(field[0]); column++ {
			if field[originalRow][column] != field[reflectedRow][column] {
				return false
			}
		}
	}
	return true
}

func parseInput(path string) []Field {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var inputs []Field
	var field Field

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputs = append(inputs, field)
			field = [][]rune{}
		} else {
			field = append(field, []rune(line))
		}
	}
	inputs = append(inputs, field)

	return inputs
}
