package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/charlie2clarke/advent-of-code-2022/datastructures"
)

var input = [][]string{
	{"G", "T", "R", "W"},
	{"G", "C", "H", "P", "M", "S", "V", "W"},
	{"C", "L", "T", "S", "G", "M"},
	{"J", "H", "D", "M", "W", "R", "F"},
	{"P", "Q", "L", "H", "S", "W", "F", "J"},
	{"P", "J", "D", "N", "F", "M", "S"},
	{"Z", "B", "D", "F", "G", "C", "S", "J"},
	{"R", "T", "B"},
	{"H", "N", "W", "L", "C"},
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	stacksPt1 := loadStacks(input)
	stacksPt2 := loadStacks(input)
	instructionsStarted := false
	pt1, pt2 := "", ""

	for scanner.Scan() {
		if scanner.Text() == "" {
			instructionsStarted = true
			continue
		}

		if instructionsStarted {
			var count, takeFrom, putTo int
			parses, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &takeFrom, &putTo)
			if err != nil || parses != 3 {
				panic("invalid input")
			}

			for i := 0; i < count; i++ {
				crate := stacksPt1[takeFrom-1].Pop()
				stacksPt1[putTo-1].Push(crate)
			}

			cratesTaken := make([]string, count)
			for i := 0; i < count; i++ {
				cratesTaken[i] = stacksPt2[takeFrom-1].Pop()
			}
			for j := len(cratesTaken) - 1; j >= 0; j-- {
				stacksPt2[putTo-1].Push(cratesTaken[j])
			}
		}
	}

	for _, stack := range stacksPt1 {
		if !stack.IsEmpty() {
			pt1 += stack.Pop()
		}
	}

	for _, stack := range stacksPt2 {
		if !stack.IsEmpty() {
			pt2 += stack.Pop()
		}
	}

	fmt.Printf("Part 1: %s\n", pt1)
	fmt.Printf("Part 2: %s\n", pt2)
}

func loadStacks(input [][]string) [9]datastructures.Stack[string] {
	var stacks [9]datastructures.Stack[string]
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			stacks[i].Push(input[i][j])
		}
	}

	return stacks
}
