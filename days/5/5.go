package main

import (
	"aoc2023/pkg/common"
	"container/heap"
	"sync"
	"log"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	end := len(*h) - 1
	out := (*h)[end]
	*h = (*h)[:end]
	return out
}

type Range struct {
	From, To, Len []int
}

type Results struct {
	sync.Mutex
	data *MinHeap
}

func (m *Range) DoMap(seed int) int {
	for i := 0; i < len(m.From); i++ {
		start := m.From[i]
		len := m.Len[i]
		if start <= seed && seed < start+len {
			return m.To[i] + (seed - start)
		}
	}
	return seed
}

func SplitLines(inp string) []string {
	return strings.Split(strings.TrimSpace(inp), "\n")
}

func ParseInts(inp string) []int {
	var out []int
	strs := strings.Split(strings.TrimSpace(inp), " ")
	for _, v := range strs {
		out = append(out, common.MustParseInt(v))
	}
	return out
}

func ParseSection(sect string) Range {
	lineStrs := SplitLines(sect)[1:]
	var fromIdx []int
	var toIdx []int
	var rng []int
	for _, line := range lineStrs {
		ints := ParseInts(line)
		fromIdx = append(fromIdx, ints[1])
		toIdx = append(toIdx, ints[0])
		rng = append(rng, ints[2])
	}
	return Range{
		From: fromIdx,
		To:   toIdx,
		Len:  rng,
	}
}

func Parse(input string) ([]int, []Range) {
	sectionStrs := strings.Split(strings.TrimSpace(input), "\n\n")
	seedsStr := strings.TrimSpace(sectionStrs[0])
	seedsStart := strings.IndexRune(seedsStr, ' ') + 1
	seeds := ParseInts(seedsStr[seedsStart:])

	var sections []Range
	for _, section := range sectionStrs[1:] {
		sect := ParseSection(section)
		sections = append(sections, sect)
	}
	return seeds, sections
}

func Part1(seeds []int, sections []Range) int {
	seedMap := make(map[int]int)
	h := &MinHeap{}
	heap.Init(h)
	for _, seed := range seeds {
		val := GetLocation(seed, sections)
		heap.Push(h, val)
		seedMap[val] = seed
	}
	return heap.Pop(h).(int)
}

func GetLocation(seed int, maps []Range) int {
	val := seed
	for _, section := range maps {
		val = section.DoMap(val)
	}
	return val
}

func DoWork(group *sync.WaitGroup, outputs *Results, seed chan int, sections []Range) {
	defer group.Done()
	for v := range seed {
		val := GetLocation(v, sections)
		outputs.Lock()
		heap.Push(outputs.data, val)
		outputs.Unlock()
	}
}

func Part2(seeds []int, sections []Range) int {
	// create worker pool
	group := sync.WaitGroup{}
	out := Results{
		data: &MinHeap{},
	}
	seedChan := make(chan int)
	for i := 0; i < 10; i++ {
		group.Add(1)
		go DoWork(&group, &out, seedChan, sections)
	}

	for i := 0; i < len(seeds)/2; i++ {
		start, len := seeds[2*i], seeds[(2*i)+1]
		for j := 0; j < len; j++ {
			seedChan <- start + j
		}
	}

	close(seedChan)
	group.Wait()

	out.Unlock()
	min := heap.Pop(out.data).(int)

	return min
}

func main() {
	input := common.MustReadFile(5)
	seeds, sects := Parse(input)
	log.Println("Part 1", Part1(seeds, sects))
	log.Println("Part 2", Part2(seeds, sects))
}
