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
