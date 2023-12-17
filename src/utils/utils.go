package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func PrintObject(object interface{}) {
	fmt.Printf("%v\n", object)
}

func StringToIntArray(input string) []int {
	return StringToIntArrayBy(input, `\s+`)
}

func StringToIntArrayBy(input string, separators string) []int {
	numbersAsString := regexp.MustCompile(separators).Split(strings.Trim(input, " "), -1)

	var numbers []int
	for _, numberAsString := range numbersAsString {
		numbers = append(numbers, UnsafeParseInt(numberAsString))
	}
	return numbers
}

func UnsafeParseInt(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}

func Copy2DSlice[T any](slice [][]T) [][]T {
	sliceCopy := make([][]T, len(slice))
	for i := 0; i < len(slice); i++ {
		sliceCopy[i] = make([]T, len(slice[0]))
		copy(sliceCopy[i], slice[i])
	}
	return sliceCopy
}
