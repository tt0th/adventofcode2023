package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"regexp"
	"strconv"
)

type Card struct {
	id             int
	myNumbers      []int
	winningNumbers []int
}

func main() {
	//path := "src/day4/test-input-1.txt"
	path := "src/day4/input.txt"
	var cards = parseCards(path)

	points := calculatePoints(cards)
	fmt.Printf("Points: %d\n", points)
}

func calculatePoints(cards []Card) int {
	var sum = 0
	for _, card := range cards {
		intersection := lo.Intersect(card.myNumbers, card.winningNumbers)
		numberOfHits := len(intersection)
		var score int
		if numberOfHits > 0 {
			score = 1 << (numberOfHits - 1)
		} else {
			score = 0
		}
		sum += score
	}
	return sum
}

var cardParserRegexp = regexp.MustCompile(`^Card\s+(?P<id>[0-9]+):\s+(?P<myNumbers>([0-9]\s*)+)\s+\|\s+(?P<winningNumbers>([0-9]\s*)+)`)

func parseCards(path string) []Card {
	var cards []Card
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseCard(line))
	}

	return cards
}

func parseCard(line string) Card {
	matches := cardParserRegexp.FindStringSubmatch(line)
	card := Card{}
	for i, name := range cardParserRegexp.SubexpNames() {
		if i != 0 && name != "" {
			if name == "id" {
				card.id = unsafeParseInt(matches[i])
			}
			if name == "myNumbers" {
				card.myNumbers = stringToIntArray(matches[i])
			}
			if name == "winningNumbers" {
				card.winningNumbers = stringToIntArray(matches[i])
			}
		}
	}
	return card
}

func stringToIntArray(input string) []int {
	numbersAsString := regexp.MustCompile(`\s+`).Split(input, -1)
	var numbers []int
	for _, numberAsString := range numbersAsString {
		numbers = append(numbers, unsafeParseInt(numberAsString))
	}
	return numbers
}

func unsafeParseInt(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}
