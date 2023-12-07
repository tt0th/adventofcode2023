package main

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Bid struct {
	hand  [5]rune
	bid   int
	value int
}

func main() {
	//path := "src/day7/test-input-1.txt"
	path := "src/day7/input.txt"
	var bids = parseInput(path)

	slices.SortFunc(bids, bidComparator)
	sum := lo.Sum(lo.Map(lo.Reverse(bids), func(bid Bid, index int) int {
		return bid.bid * (index + 1)
	}))
	fmt.Printf("Sum: %d\n", sum)
}

func bidComparator(a, b Bid) int {
	firstPriorityComparation := b.value - a.value
	secondPriorityComparation := strings.Compare(string(b.hand[:]), string(a.hand[:]))
	return firstPriorityComparation*100 + secondPriorityComparation
}

func parseInput(path string) []Bid {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var bids []Bid

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bids = append(bids, parseBid(line))
	}
	return lo.Map(bids, fillBidHandValue)
}

func parseBid(line string) Bid {
	splitByWhitespace := strings.Split(line, " ")
	var runes [5]rune
	for i, char := range splitByWhitespace[0] {
		runes[i] = swapLetterToComparableOne(char)
	}
	return Bid{hand: runes, bid: unsafeParseInt(splitByWhitespace[1])}
}

func swapLetterToComparableOne(character rune) rune {
	switch character {
	case 'T':
		return 'A'
	case 'J':
		return 'B'
	case 'Q':
		return 'C'
	case 'K':
		return 'D'
	case 'A':
		return 'E'
	}
	return character
}

func fillBidHandValue(bid Bid, _ int) Bid {
	groupedRunes := lo.PartitionBy(bid.hand[:], func(item rune) rune {
		return item
	})
	var typeValue int
	// high card
	if len(groupedRunes) == 5 {
		typeValue = 0
	}
	// 1 pair
	if len(groupedRunes) == 4 {
		typeValue = 1
	}
	// 2 pairs
	if len(groupedRunes) == 3 && (len(groupedRunes[0]) == 2 || len(groupedRunes[1]) == 2) {
		typeValue = 2
	}
	// 3 of a kind
	if len(groupedRunes) == 3 && (len(groupedRunes[0]) == 3 || len(groupedRunes[1]) == 3 || len(groupedRunes[2]) == 3) {
		typeValue = 3
	}
	// full
	if len(groupedRunes) == 2 && (len(groupedRunes[0]) == 3 || len(groupedRunes[1]) == 3) {
		typeValue = 4
	}
	// 4 of a kind
	if len(groupedRunes) == 2 && (len(groupedRunes[0]) == 4 || len(groupedRunes[1]) == 4) {
		typeValue = 5
	}
	// 5 of a kind
	if len(groupedRunes) == 1 {
		typeValue = 6
	}

	bid.value = typeValue
	return bid
}

func unsafeParseInt(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}
