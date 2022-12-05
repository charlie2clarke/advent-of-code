package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Stack[T]) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s.Swap(i, j)
	}
}

func (s *Stack[T]) Copy() *Stack[T] {
	s2 := NewStack[T]()
	*s2 = append(*s2, *s...)
	return s2
}

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

func loadStacks(input [][]string) [9]Stack[string] {
	var stacks [9]Stack[string]
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			switch i {
			case 0:
				stacks[0].Push(input[i][j])
			case 1:
				stacks[1].Push(input[i][j])
			case 2:
				stacks[2].Push(input[i][j])
			case 3:
				stacks[3].Push(input[i][j])
			case 4:
				stacks[4].Push(input[i][j])
			case 5:
				stacks[5].Push(input[i][j])
			case 6:
				stacks[6].Push(input[i][j])
			case 7:
				stacks[7].Push(input[i][j])
			case 8:
				stacks[8].Push(input[i][j])
			}
		}
	}

	return stacks
}
