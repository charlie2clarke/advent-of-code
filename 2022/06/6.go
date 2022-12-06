package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charlie2clarke/advent-of-code/datastructures"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	signal := strings.TrimSpace(string(file))

	pt1, pt2 := 0, 0

	pt1 += findMarker(signal, 4)
	pt2 += findMarker(signal, 14)

	fmt.Printf("Part 1: %v\n", pt1)
	fmt.Printf("Part 2: %v\n", pt2)
}

func findMarker(signal string, windowLen int) int {
	for i := windowLen; i < len(signal); i++ {
		set := datastructures.NewSet[string]()
		for _, c := range signal[i-windowLen : i] {
			set.Add(string(c))
		}

		if set.Len() == windowLen {
			return i
		}
	}

	return 0
}
