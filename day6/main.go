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

	for i := windowLen; i < len(signal); i++ {
		for _, c := range signal[i-windowLen : i] {
			set.Add(string(c))
		}

		if set.Len() == windowLen {
			return i
		}

		set.Clear()
	}

	return 0
}
