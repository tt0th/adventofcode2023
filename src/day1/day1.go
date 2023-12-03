package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//path := "src/day1/test-input-1.txt"
	//path := "src/day1/test-input-2.txt"
	path := "src/day1/input.txt"
	file, _ := os.Open(path)
	defer file.Close()

	var sum = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += extractNumber(line)
	}

	fmt.Printf("Sum: %d\n", sum)
}

func extractNumber(line string) int {
	firstDigitCharacter := getFirstDigitCharacter(replaceWrittenNumbersWithDigits(line))
	lastDigitCharacter := getLastDigitCharacter(replaceWrittenNumbersWithDigitsFromTheBack(line))

	print(line, " ")
	print(string(firstDigitCharacter), " ", string(lastDigitCharacter), " ")
	extractedNumber, _ := strconv.Atoi(string(firstDigitCharacter) + string(lastDigitCharacter))
	print("\n")
	return extractedNumber
}

func getLastDigitCharacter(line string) rune {
	var lastDigitCharacter rune
	for _, character := range line {
		if character >= '0' && character <= '9' {
			lastDigitCharacter = character
		}
	}
	return lastDigitCharacter
}

func getFirstDigitCharacter(line string) rune {
	var firstDigitCharacter rune
	for _, character := range line {
		if character >= '0' && character <= '9' {
			firstDigitCharacter = character
			break
		}
	}
	return firstDigitCharacter
}

func replaceWrittenNumbersWithDigits(line string) string {
	replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
	return replacer.Replace(line)
}

func replaceWrittenNumbersWithDigitsFromTheBack(line string) string {
	replacer := strings.NewReplacer("eno", "1", "owt", "2", "eerht", "3", "ruof", "4", "evif", "5", "xis", "6", "neves", "7", "thgie", "8", "enin", "9")
	return reverseString(replacer.Replace(reverseString(line)))
}

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
