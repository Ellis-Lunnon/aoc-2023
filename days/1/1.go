package main

import (
	"aoc2023/pkg/common"
	"log"
	"strings"
	"sync"
)

var PatternMap map[string]int
var Part1Tokens []string
var Part2Tokens []string

type Counter struct {
	sync.Mutex
	group *sync.WaitGroup
	Val   int
}

func (c *Counter) AsyncParseTokens(tokens []string, str string) {
	defer c.group.Done()
	val := ParseTokens(tokens, str)
	c.Lock()
	defer c.Unlock()
	c.Val += val
}

func ParseTokens(tokens []string, str string) int {
	firstIdx, firstVal := len(str), 0
	lastIdx, lastVal := 0, 0
	hasLastVal := false
	for _, token := range tokens {
		idx := strings.Index(str, token)
		if idx != -1 && idx < firstIdx {
			firstIdx = idx
			firstVal = PatternMap[token]
		}

		idx = strings.LastIndex(str, token)
		if idx != -1 && idx > lastIdx {
			lastIdx = idx
			lastVal = PatternMap[token]
			hasLastVal = true
		}
	}
	if !hasLastVal {
		lastVal = firstVal
	}
	return 10*firstVal + lastVal
}

func Part1(lines []string) int {
	total := Counter{
		group: &sync.WaitGroup{},
	}
	for _, line := range lines {
		total.group.Add(1)
		go total.AsyncParseTokens(Part1Tokens, line)
	}

	total.group.Wait()
	return total.Val
}

func Part2(lines []string) int {
	total := Counter{
		group: &sync.WaitGroup{},
	}
	for _, line := range lines {
		total.group.Add(1)
		go total.AsyncParseTokens(Part2Tokens, line)
	}
	total.group.Wait()
	return total.Val
}

func init() {
	// Set up the tokens for parts 1 and two
	PatternMap = make(map[string]int)
	Part1Tokens = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	Part2Tokens = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 1; i < 10; i++ {
		PatternMap[Part1Tokens[i-1]] = i
		PatternMap[Part2Tokens[i-1]] = i
	}
	Part2Tokens = append(Part2Tokens, Part1Tokens...)
}

func main() {
	input := common.MustReadFile(1)
	lines := strings.Split(input, "\n")

	log.Println("Part 1", Part1(lines))
	log.Println("Part 2", Part2(lines))
}
