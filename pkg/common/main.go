package common

import (
	"fmt"
	"os"
	"strconv"
)

func MustReadFile(day int) string {
	data, err := os.ReadFile(fmt.Sprintf("inputs/day%d.txt", day))
	if err != nil {
		panic(err)
	}
	return string(data)
}

func MustParseInt(val string) int {
	out, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return out
}
