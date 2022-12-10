package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const cycleCount = 40
	file, _ := os.ReadFile("input.txt")
	cycle, register, signalStrength, crt := 0, 1, 0, ""

	tick := func() {
		crt += map[bool]string{true: "#", false: "."}[cycle%40 >= register-1 && cycle%cycleCount <= register+1]
		crt += map[bool]string{true: "\n"}[cycle%cycleCount == cycleCount-1]
		cycle++

		if (cycle+cycleCount/2)%cycleCount == 0 {
			signalStrength += cycle * register
		}
	}

	for _, line := range strings.Split(string(file), "\n") {
		tick()

		if strings.HasPrefix(line, "addx") {
			var value int
			fmt.Sscanf(line, "addx %d", &value)

			tick()
			register += value
		}
	}

	fmt.Printf("Part 1: %d\n", signalStrength)
	fmt.Printf("Part 2: \n%s", crt)
}
