package model

import (
	"strings"
)

type Move int64

const (
	Rock Move = iota
	Paper
	Scissors
)

var mappings = map[string]Move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var points = map[Move]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var wins = map[Move]Move{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

type Round struct {
	OpponentMove Move
	MyMove       Move
}

func RoundFromLine(line string) Round {
	m := strings.Split(line, " ")

	return Round{
		OpponentMove: mappings[m[0]],
		MyMove:       mappings[m[1]],
	}
}

func (r Round) GetPoints() int {
	movePoints := points[r.MyMove]
	winPoints := 0
	if wins[r.MyMove] == r.OpponentMove {
		winPoints = 6
	} else if r.MyMove == r.OpponentMove {
		winPoints = 3
	} else {
		winPoints = 0
	}

	return winPoints + movePoints
}
