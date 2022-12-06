package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/charlie2clarke/advent-of-code/datastructures"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	subSetCountPt1 := 0
	intersectionCountPt2 := 0

	for scanner.Scan() {
		elfOneSet, elfTwoSet := datastructures.Set[int]{}, datastructures.Set[int]{}
		elfOneStart, elfOneStop, elfTwoStart, elfTwoStop := 0, 0, 0, 0

		parses, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &elfOneStart, &elfOneStop, &elfTwoStart, &elfTwoStop)
		if err != nil || parses != 4 {
			panic("error parsing input")
		}

		for i := elfOneStart; i <= elfOneStop; i++ {
			elfOneSet.Add(i)
		}

		for i := elfTwoStart; i <= elfTwoStop; i++ {
			elfTwoSet.Add(i)
		}

		isSubSet := elfOneSet.IsSubset(elfTwoSet) || elfTwoSet.IsSubset(elfOneSet)
		if isSubSet {
			subSetCountPt1 += 1
		}

		intersection := elfOneSet.Intersection(elfTwoSet)
		if intersection.Len() > 0 {
			intersectionCountPt2 += 1
		}
	}

	fmt.Printf("Part 1: %v\n", subSetCountPt1)
	fmt.Printf("Part 2: %v\n", intersectionCountPt2)
}
