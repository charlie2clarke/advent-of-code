package main

import (
	"bufio"
	"fmt"
	"os"
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
		compartment1 := bag[:len(bag)/2]
		compartment2 := bag[len(bag)/2:]
		group = append(group, bag)

		priorityPart1 += alphabet[findDuplicate(compartment1, compartment2)]

		if len(group) == 3 {
			duplicate := findDuplicate2(group[0], group[1], group[2])
			priorityPart2 += alphabet[duplicate]
			group = group[:0]
		}
	}

	fmt.Printf("Part 1: %v\n", priorityPart1)
	fmt.Printf("Part 2: %v\n", priorityPart2)
}

func findDuplicate(compartment1, compartment2 string) string {
	for i := 0; i < len(compartment1); i++ {
		for j := 0; j < len(compartment2); j++ {
			if compartment1[i] == compartment2[j] {
				return string(compartment1[i])
			}
		}
	}
	panic("No duplicate found")
}

func findDuplicate2(bag1, bag2, bag3 string) string {
	for i := 0; i < len(bag1); i++ {
		for j := 0; j < len(bag2); j++ {
			for k := 0; k < len(bag3); k++ {
				if bag1[i] == bag2[j] && bag2[j] == bag3[k] {
					return string(bag1[i])
				}
			}
		}
	}
	panic("No duplicate found")
}
