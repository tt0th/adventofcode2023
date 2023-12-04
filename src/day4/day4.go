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
	numberOfHits   int
	count          int
}

func main() {
	//path := "src/day4/test-input-1.txt"
	path := "src/day4/input.txt"
	var cards = parseCards(path)

	points := calculatePoints(cards)
	count := countCards(cards)
	fmt.Printf("Points: %d, Cards: %d\n", points, count)
}

func calculatePoints(cards []Card) int {
	var sum = 0
	for _, card := range cards {
		var score int
		if card.numberOfHits > 0 {
			score = 1 << (card.numberOfHits - 1)
		} else {
			score = 0
		}
		sum += score
	}
	return sum
}

func countCards(cards []Card) int {
	for i := 0; i < len(cards); i++ {
		for j := 1; j < cards[i].numberOfHits+1; j++ {
			cards[i+j].count += cards[i].count
		}
	}
	return lo.SumBy(cards, func(card Card) int {
		return card.count
	})
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
	card.numberOfHits = len(lo.Intersect(card.myNumbers, card.winningNumbers))
	card.count = 1
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
