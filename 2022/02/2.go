package main

import (
	"bufio"
	"fmt"
	"os"
)

type Part1Round struct {
	opponent Move
	mine     Move
}

func (r Part1Round) Score() int {
	const (
		LOSS = 0
		DRAW = 3
		WIN  = 6
	)
	if r.opponent == r.mine {
		return DRAW + r.mine.MoveScore()
	} else if r.opponent == ROCK && r.mine == SCISSORS {
		return LOSS + r.mine.MoveScore()
	} else if r.opponent == PAPER && r.mine == ROCK {
		return LOSS + r.mine.MoveScore()
	} else if r.opponent == SCISSORS && r.mine == PAPER {
		return LOSS + r.mine.MoveScore()
	} else {
		return WIN + r.mine.MoveScore()
	}
}

type Part2Round struct {
	opponent Move
	mine     Outcome
}

func (r Part2Round) Score() int {
	if r.mine == WIN {
		return WIN.OutcomeScore() + r.opponent.WinningMove().MoveScore()
	} else if r.mine == DRAW {
		return DRAW.OutcomeScore() + r.opponent.MoveScore()
	} else if r.mine == LOSS {
		return LOSS.OutcomeScore() + r.opponent.LosingMove().MoveScore()
	} else {
		panic("Invalid outcome")
	}
}

type Move string

const (
	ROCK     Move = "R"
	PAPER    Move = "P"
	SCISSORS Move = "S"
)

func (d Move) TranslateMove() Move {
	if d == "A" || d == "X" {
		return ROCK
	} else if d == "B" || d == "Y" {
		return PAPER
	} else if d == "C" || d == "Z" {
		return SCISSORS
	} else {
		panic("Invalid draw")
	}
}

func (m Move) MoveScore() int {
	if m == ROCK {
		return 1
	} else if m == PAPER {
		return 2
	} else if m == SCISSORS {
		return 3
	} else {
		panic("Invalid move")
	}
}

func (m Move) WinningMove() Move {
	if m == ROCK {
		return PAPER
	} else if m == PAPER {
		return SCISSORS
	} else if m == SCISSORS {
		return ROCK
	} else {
		panic("Invalid move")
	}
}

func (m Move) LosingMove() Move {
	if m == ROCK {
		return SCISSORS
	} else if m == PAPER {
		return ROCK
	} else if m == SCISSORS {
		return PAPER
	} else {
		panic("Invalid move")
	}
}

type Outcome string

const (
	WIN  Outcome = "W"
	DRAW Outcome = "D"
	LOSS Outcome = "L"
)

func (o Outcome) TranslateOutcome() Outcome {
	if o == "Z" {
		return WIN
	} else if o == "Y" {
		return DRAW
	} else if o == "X" {
		return LOSS
	} else {
		panic("Invalid outcome")
	}
}

func (o Outcome) OutcomeScore() int {
	if o == WIN {
		return 6
	} else if o == DRAW {
		return 3
	} else if o == LOSS {
		return 0
	} else {
		panic("Invalid outcome")
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	part1Score, part2Score := 0, 0

	for scanner.Scan() {
		part1Score += part1(scanner.Text())
		part2Score += part2(scanner.Text())
	}

	fmt.Printf("Part 1: %v\n", part1Score)
	fmt.Printf("Part 2: %v\n", part2Score)
}

func part1(input string) int {
	var (
		opponent Move
		mine     Move
	)

	linesParsed, err := fmt.Sscanf(input, "%s %s", &opponent, &mine)
	if err != nil || linesParsed != 2 {
		panic("Invalid input")
	}

	round := Part1Round{opponent.TranslateMove(), mine.TranslateMove()}
	return round.Score()
}

func part2(input string) int {
	var (
		opponent Move
		mine     Outcome
	)

	linesParsed, err := fmt.Sscanf(input, "%s %s", &opponent, &mine)
	if err != nil || linesParsed != 2 {
		panic("Invalid input")
	}

	round := Part2Round{opponent.TranslateMove(), mine.TranslateOutcome()}
	return round.Score()
}
