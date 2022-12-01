package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
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

	sort.IntSlice(eachCalories).Sort()

	firstHighest, secondHighest, thirdHighest := eachCalories[len(eachCalories)-1], eachCalories[len(eachCalories)-2], eachCalories[len(eachCalories)-3]

	fmt.Printf("Part 1: %v\n", firstHighest)
	fmt.Printf("Part 2: %v", firstHighest+secondHighest+thirdHighest)
}
