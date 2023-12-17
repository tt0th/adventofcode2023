package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"github.com/tt0th/adventofcode2023/src/utils"
	"os"
	"regexp"
	"strings"
)

type Box struct {
	lenses []Lens
}
type Lens struct {
	label string
	focus int
}

type Step struct {
	label     string
	operation rune
	focus     int
	hash      int
}

const SET = '='
const UNSET = '-'

func main() {
	//path := "src/day15/test-input-1.txt"
	path := "src/day15/input.txt"
	var inputs = parseInput(path)

	sum := lo.SumBy(inputs, func(item string) int {
		return hash(item)
	})
	fmt.Printf("sum: %d\n", sum)

	var steps = parseSteps(inputs)
	var boxes = make([]Box, 256)
	executeSteps(steps, boxes)
	fmt.Printf("sum2: %d\n", sumFocus(boxes))
}

func sumFocus(boxes []Box) int {
	return lo.Sum(lo.Map(boxes, func(box Box, boxIndex int) int {
		return lo.Sum(lo.Map(box.lenses, func(lens Lens, lensIndex int) int {
			return (boxIndex + 1) * (lensIndex + 1) * lens.focus
		}))
	}))
}

func executeSteps(steps []Step, boxes []Box) {
	for _, step := range steps {
		if step.operation == SET {
			setLens(step, boxes)
		} else {
			unsetLens(boxes, step)
		}
	}
}

func setLens(step Step, boxes []Box) {
	lens := Lens{label: step.label, focus: step.focus}
	_, indexOfExistingLens, found := lo.FindIndexOf(boxes[step.hash].lenses, func(lens Lens) bool {
		return lens.label == step.label
	})
	if found {
		boxes[step.hash].lenses[indexOfExistingLens] = lens
	} else {
		boxes[step.hash].lenses = append(boxes[step.hash].lenses, lens)
	}
}

func unsetLens(boxes []Box, step Step) {
	boxes[step.hash].lenses = lo.Filter(boxes[step.hash].lenses, func(lens Lens, _ int) bool {
		return lens.label != step.label
	})
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

var stepParserRegexp = regexp.MustCompile(`^(?P<label>[a-z]+)(?P<operation>[\-=])(?P<focus>[0-9])?`)

func parseSteps(steps []string) []Step {
	return lo.Map(steps, func(step string, _ int) Step {
		return parseStep(step)
	})
}
func parseStep(stepString string) Step {
	matches := stepParserRegexp.FindStringSubmatch(stepString)
	step := Step{}
	for i, name := range stepParserRegexp.SubexpNames() {
		if i != 0 && name != "" {
			if name == "label" {
				step.label = matches[i]
			}
			if name == "operation" {
				step.operation = []rune(matches[i])[0]
			}
			if name == "focus" {
				step.focus = utils.UnsafeParseInt(matches[i])
			}
		}
	}
	step.hash = hash(step.label)
	return step
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
