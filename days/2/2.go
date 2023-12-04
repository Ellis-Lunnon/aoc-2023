package main

import (
	"aoc2023/pkg/common"
	"log"
	"strings"
)

var Limits map[string]int

func Part1(games []string) int {
	total := 0
	// For each game
	for _, game := range games {
		if game == "" {
			break
		}
		parts := strings.Split(game, ":")
		id := common.MustParseInt(parts[0][5:])

		draws := strings.Split(parts[1], ";")
		isPossible := true
	Game:
		// For each round
		for _, round := range draws {
			fields := strings.Split(round, ",")

			// For each set of key value pairs
			for _, field := range fields {
				attrs := strings.Split(strings.TrimSpace(field), " ")
				label, val := attrs[1], common.MustParseInt(attrs[0])
				if Limits[label] < val {
					isPossible = false
					break Game
				}
			}
		}

		if isPossible {
			total += id
		}
	}

	return total
}

func Part2(games []string) int {
	total := 0
	// For each game
	for _, game := range games {
		if game == "" {
			break
		}
		parts := strings.Split(game, ":")

		draws := strings.Split(parts[1], ";")
		maxCubes := make(map[string]int)
		// For each round
		for _, round := range draws {
			fields := strings.Split(round, ",")

			// For each set of key value pairs
			for _, field := range fields {
				attrs := strings.Split(strings.TrimSpace(field), " ")
				label, val := attrs[1], common.MustParseInt(attrs[0])
				if maxCubes[label] < val {
					maxCubes[label] = val
				}
			}
		}

		pow := 1
		for _, v := range maxCubes {
			pow *= v
		}
		total += pow
	}

	return total
}

func main() {
	input := common.MustReadFile(2)
	games := strings.Split(input, "\n")
	log.Println("Part 1", Part1(games))
	log.Println("Part 2", Part2(games))
}

func init() {
	Limits = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

}
