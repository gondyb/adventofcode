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

type Outcome int64

const (
	Win Outcome = iota
	Lose
	Draw
)

var moveMappings = map[string]Move{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var outcomeMappigs = map[string]Outcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

var points = map[Move]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

type Game map[Move]Move

var outcomes = map[Outcome]Game{
	Win: map[Move]Move{
		Scissors: Rock,
		Rock:     Paper,
		Paper:    Scissors,
	},
	Lose: map[Move]Move{
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	},
	Draw: map[Move]Move{
		Scissors: Scissors,
		Rock:     Rock,
		Paper:    Paper,
	},
}

type Round struct {
	OpponentMove Move
	MyMove       Move
}

func RoundFromLine(line string) Round {
	m := strings.Split(line, " ")

	action := outcomeMappigs[m[1]]
	opponentMove := moveMappings[m[0]]
	myMove := outcomes[action][opponentMove]

	return Round{
		OpponentMove: opponentMove,
		MyMove:       myMove,
	}
}

func (r Round) GetPoints() int {
	movePoints := points[r.MyMove]
	winPoints := 0
	if outcomes[Lose][r.MyMove] == r.OpponentMove {
		winPoints = 6
	} else if r.MyMove == r.OpponentMove {
		winPoints = 3
	} else {
		winPoints = 0
	}

	return winPoints + movePoints
}
