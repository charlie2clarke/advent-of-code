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

	alphabet := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
		"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52,
	}
	priorityPart1, priorityPart2 := 0, 0
	group := []string{}

	for scanner.Scan() {
		bag := scanner.Text()
		compartment1, compartment2 := datastructures.NewSet[string](), datastructures.NewSet[string]()

		// add the first half of the letters in bag to compartment1
		for i := 0; i < len(bag)/2; i++ {
			compartment1.Add(string(bag[i]))
		}

		// add the second half of the letters in bag to compartment2
		for i := len(bag) / 2; i < len(bag); i++ {
			compartment2.Add(string(bag[i]))
		}

		priorityPart1 += alphabet[findIntersectionPt1(compartment1, compartment2)]

		group = append(group, bag)

		if len(group) == 3 {
			bag1, bag2, bag3 := datastructures.NewSet[string](), datastructures.NewSet[string](), datastructures.NewSet[string]()
			for i := 0; i < len(group[0]); i++ {
				bag1.Add(string(group[0][i]))
			}

			for i := 0; i < len(group[1]); i++ {
				bag2.Add(string(group[1][i]))
			}

			for i := 0; i < len(group[2]); i++ {
				bag3.Add(string(group[2][i]))
			}

			priorityPart2 += alphabet[findIntersectionPt2(bag1, bag2, bag3)]
			group = group[:0]
		}
	}

	fmt.Printf("Part 1: %v\n", priorityPart1)
	fmt.Printf("Part 2: %v\n", priorityPart2)
}

func findIntersectionPt1(compartment1, compartment2 datastructures.Set[string]) string {
	intersection := compartment1.Intersection(compartment2)
	if intersection.Len() != 1 {
		panic("invalid intersection found")
	}

	// return the only element in the intersection
	for k := range intersection {
		return k
	}

	panic("No intersection found")
}

func findIntersectionPt2(bag1, bag2, bag3 datastructures.Set[string]) string {
	// find the intersection of bag1, bag2, and bag3
	intersection := bag1.Intersection(bag2).Intersection(bag3)
	if intersection.Len() != 1 {
		panic("invalid intersection found")
	}

	// return the only element in the intersection
	for k := range intersection {
		return k
	}

	panic("No intersection found")
}
