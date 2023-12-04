package main

import (
	"aoc2023/pkg/common"
	"log"
	"regexp"
	"strings"
)

var SymbolicRunes map[rune]bool

func HasNeighbourPart(mat *common.Mat[rune], x, y int) (bool, []int) {
	var symbolIdx []int
	abs := mat.GetIdx(x, y)
	indices := []int{
		// Row above
		abs - mat.X - 1,
		abs - mat.X,
		abs - mat.X + 1,
		// Left and right
		abs - 1, abs + 1,
		// Row below
		abs + mat.X - 1,
		abs + mat.X,
		abs + mat.X + 1,
	}
	for _, idx := range indices {
		// Check we are still at a valid index
		if idx < 0 || idx >= len(mat.Data) {
			continue
		}
		newX, newY := mat.GetXY(idx)
		// Check if we have wrapped the array
		if x-newX > 1 || x-newX < -1 ||
			y-newY > 1 || y-newY < -1 {
			continue
		}
		char := mat.Get(newX, newY)
		if !SymbolicRunes[char] {
			symbolIdx = append(symbolIdx, mat.GetIdx(newX, newY))
		}
	}
	return len(symbolIdx) > 0, symbolIdx
}

func FindLabels(mat *common.Mat[rune]) [][]int {
	expr := regexp.MustCompile(`\d+`)
	var matches [][]int
	// For each row
	for rowIdx := 0; rowIdx < mat.X; rowIdx++ {
		start := rowIdx * mat.Y
		end := (rowIdx + 1) * mat.Y
		row := mat.Data[start:end]
		// Get matches & convert to absolute idx
		rowMatches := expr.FindAllStringIndex(string(row), -1)
		for i := 0; i < len(rowMatches); i++ {
			absMatch := []int{
				rowMatches[i][0] + start,
				rowMatches[i][1] + start,
			}
			matches = append(matches, absMatch)
		}
	}
	return matches
}

func Part1(mat common.Mat[rune]) int {
	var partLabels []int
	matches := FindLabels(&mat)
	for _, match := range matches {
		start, end := match[0], match[1]
		for idx := start; idx < end; idx++ {
			x, y := mat.GetXY(idx)
			if hasPart, _ := HasNeighbourPart(&mat, x, y); hasPart {
				val := common.MustParseInt(string(mat.Data[start:end]))
				partLabels = append(partLabels, val)
				break
			}
		}
	}
	total := 0
	for _, v := range partLabels {
		total += v
	}
	return total
}

func Part2(mat common.Mat[rune]) int {
	// Build a map of labels to parts
	// Use that to make a map of part to applied labels
	labels := FindLabels(&mat)
	partLabelMap := make(map[int][]int)
	total := 0

	neighbours := make(map[int]bool)
	for _, label := range labels {
		// Clear the neighbours map
		for k := range neighbours {
			delete(neighbours, k)
		}

		// For each char of the label check for adjacent parts
		start, end := label[0], label[1]
		for idx := start; idx < end; idx++ {
			x, y := mat.GetXY(idx)
			// If there's a part then mark it as known and store
			// the value of this label under that parts entry
			_, parts := HasNeighbourPart(&mat, x, y)
			for _, part := range parts {
				if !neighbours[part] {
					val := common.MustParseInt(string(mat.Data[start:end]))
					partLabelMap[part] = append(partLabelMap[part], val)
					neighbours[part] = true
				}
			}
		}
	}

	// Find all the gear parts from that map
	for k, v := range partLabelMap {
		x, y := mat.GetXY(k)
		char := mat.Get(x, y)
		if char == '*' && len(v) == 2 {
			total += v[0] * v[1]
		}
	}
	return total
}

func main() {
	data := common.MustReadFile(3)
	x := len(strings.Split(data, "\n")) - 1
	y := len(strings.Split(data, "\n")[0])
	mat := common.NewMat[rune](x, y, []rune(strings.ReplaceAll(data, "\n", "")))
	log.Println("Part 1", Part1(mat))
	log.Println("Part 2", Part2(mat))
}

func init() {
	SymbolicRunes = map[rune]bool{
		'.': true, '0': true, '1': true,
		'2': true, '3': true, '4': true,
		'5': true, '6': true, '7': true,
		'8': true, '9': true,
	}
}
