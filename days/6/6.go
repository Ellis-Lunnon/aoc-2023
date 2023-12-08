package main

import (
	"aoc2023/pkg/common"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ParseLine(line string) []int {
	expr := regexp.MustCompile(`\d+`)
	vals := expr.FindAllString(line, -1)
	ints := make([]int, len(vals))
	for i, s := range vals {
		ints[i] = common.MustParseInt(s)
	}
	return ints
}

func ParseLine2(line string) int {
	nums := strings.Split(line, ":")[1]
	nospace := strings.ReplaceAll(nums, " ", "")
	return common.MustParseInt(nospace)
}

func CalcDist(timeHeld, totalTime int) int {
	speed := timeHeld
	runtime := totalTime - timeHeld
	return runtime * speed
}

func CalcNumWins(totalTime, raceBest int) int {
	cnt := 0
	dists := make([]int, totalTime-1)
	for i := 1; i < totalTime; i++ {
		dists[i-1] = CalcDist(i, totalTime)
		if dists[i-1] > raceBest {
			cnt++
		}
	}
	return cnt
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	times, dists := ParseLine(lines[0]), ParseLine(lines[1])

	numRaces := len(times)
	totals := make([]int, numRaces)
	for i := 0; i < numRaces; i++ {
		totals[i] = CalcNumWins(times[i], dists[i])
	}
	fmt.Printf("totals: %v\n", totals)
	if len(totals) == 0 {
		return 0
	}
	total := 1
	for _, v := range totals {
		total *= v
	}
	return total
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	time, dist := ParseLine2(lines[0]), ParseLine2(lines[1])
	return CalcNumWins(time, dist)
}

func main() {
	input := common.MustReadFile(6)
	log.Println("Part 1:", Part1(input))
	log.Println("Part 2:", Part2(input))
}
