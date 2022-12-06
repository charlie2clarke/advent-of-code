package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/charlie2clarke/advent-of-code-2022/datastructures"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pt1, pt2 := 0, 0

	for scanner.Scan() {
		signal := scanner.Text()

		pt1 += findMarker(signal, 4)
		pt2 += findMarker(signal, 14)
	}

	if pt1 == 0 {
		panic("no solution found to part 1")
	}

	if pt2 == 0 {
		panic("no solution found to part 2")
	}

	fmt.Printf("Part 1: %v\n", pt1)
	fmt.Printf("Part 2: %v\n", pt2)
}

func findMarker(signal string, windowLen int) int {
	set := datastructures.NewSet[string]()

	for i := 0; i < len(signal); i++ {
		if i+windowLen > len(signal) {
			break
		}
		window := signal[i : i+windowLen]

		for j := 0; j < len(window); j++ {
			set.Add(string(window[j]))
		}

		if set.Len() == windowLen {
			return i + windowLen
		}

		set.Clear()
	}

	return 0
}
