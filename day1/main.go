package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	eachCalories := make([]int, 0)
	currentCalories := 0

	for scanner.Scan() {
		snack, err := strconv.Atoi(scanner.Text())
		currentCalories += snack

		if err != nil {
			eachCalories = append(eachCalories, currentCalories)
			currentCalories = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(eachCalories)))

	firstHighest, secondHighest, thirdHighest := eachCalories[1], eachCalories[2], eachCalories[3]

	fmt.Printf("Part 1: %v\n", firstHighest)
	fmt.Printf("Part 2: %v", firstHighest+secondHighest+thirdHighest)
}
