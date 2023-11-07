package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/charlie2clarke/advent-of-code/datastructures"
)

type Monkey struct {
	items     datastructures.Queue[int]
	op        func(int) (newWorry int)
	test      func(int) (throwTo int)
	inspected int
}

func (m *Monkey) inspect() int {
	m.inspected++
	return m.items.Dequeue()
}

func loadMonkeys(fileString string) ([]*Monkey, int) {
	monkeySplit := strings.Split(fileString, "\n\n")
	monkeys := make([]*Monkey, len(monkeySplit))
	lcm := 1

	for i, monkey := range monkeySplit {
		var items, op string
		var monkeyNum, opValue, test, trueOp, falseOp int
		parses, err := fmt.Sscanf(strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(monkey),
			`Monkey %d:
				Starting items: %s
				Operation: new = old %s %d
				Test: divisible by %d
					If true: throw to monkey %d
					If false: throw to monkey %d`,
			&monkeyNum, &items, &op, &opValue, &test, &trueOp, &falseOp)
		if parses != 7 || err != nil {
			panic("couldn't parse monkeys")
		}

		monkeys[i] = &Monkey{}
		if err := json.Unmarshal([]byte("["+items+"]"), &monkeys[i].items); err != nil {
			panic(err)
		}

		monkeys[i].op = map[string]func(int) int{
			"*": func(worry int) int { return worry * opValue },
			"+": func(worry int) int { return worry + opValue },
			"^": func(worry int) int { return worry * worry },
		}[op]

		monkeys[i].test = func(worry int) (throwTo int) {
			if worry%test == 0 {
				return trueOp
			}
			return falseOp
		}
		lcm *= test
	}

	return monkeys, lcm
}

func run(monkeys []*Monkey, rounds int, worryManage func(m []*Monkey, w int) int) {
	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			if monkey.items.IsEmpty() {
				continue
			}

			for worryLevel := monkey.inspect(); ; worryLevel = monkey.inspect() {
				worryLevel = monkey.op(worryLevel)
				worryLevel = worryManage(monkeys, worryLevel)
				monkeys[monkey.test(worryLevel)].items.Enqueue(worryLevel)

				if monkey.items.IsEmpty() {
					break
				}
			}
		}
	}
}

func calcMonkeyBusiness(monkeys []*Monkey) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	return monkeys[0].inspected * monkeys[1].inspected
}

func part1(monkeys []*Monkey) int {
	run(monkeys, 20, func(m []*Monkey, w int) int { return w / 3 })
	return calcMonkeyBusiness(monkeys)
}

func part2(monkeys []*Monkey, lcm int) int {
	run(monkeys, 10000, func(m []*Monkey, w int) int { return w % lcm })
	return calcMonkeyBusiness(monkeys)
}

func main() {
	file, _ := os.ReadFile("input.txt")
	monkeys1, _ := loadMonkeys(string(file))
	monkeys2, lcm := loadMonkeys(string(file))

	fmt.Printf("Part 1: %d\n", part1(monkeys1))
	fmt.Printf("Part 2: %d\n", part2(monkeys2, lcm))
}
