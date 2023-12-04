package main

import (
	"aoc2023/pkg/common"
	"log"
	"strings"
)

func ParseCardWins(card string) int {
	cnt := 0
	winnersMap := make(map[string]bool)
	cardDataLst := strings.Split(card, ":")
	cardData := strings.TrimSpace(cardDataLst[1])

	cardGame := strings.Split(cardData, "|")
	winningStr, mineStr := strings.TrimSpace(cardGame[0]), strings.TrimSpace(cardGame[1])
	mine := strings.Split(mineStr, " ")
	for _, winNumber := range strings.Split(winningStr, " ") {
		if winNumber != "" {
			winnersMap[winNumber] = true
		}
	}
	for _, no := range mine {
		if winnersMap[no] {
			cnt += 1
		}
	}

	return cnt
}

func Part1(cards []string) int {
	total := 0
	for _, card := range cards {
		if card == "" {
			break
		}
		cnt := ParseCardWins(card)
		score := 0
		if cnt > 0 {
			score = 1
			cnt--
		}
		for ; cnt > 0; cnt-- {
			score *= 2
		}
		total += score
	}

	return total
}

func Part2(cards []string) int {
	cardCounts := map[int]int{}
	for idx, card := range cards {
		if card == "" {
			break
		}
		id := idx + 1
		cnt := ParseCardWins(card)
		// Increment by one to account for the original
		cardCounts[id]++
		numCopies := cardCounts[id]
		for i := 0; i < cnt; i++ {
			cardCounts[id+i+1] += numCopies
		}
	}

	total := 0
	for _, v := range cardCounts {
		total += v
	}
	return total
}

func main() {
	data := common.MustReadFile(4)
	cards := strings.Split(data, "\n")
	log.Println("Part 1", Part1(cards))
	log.Println("Part 2", Part2(cards))

}
