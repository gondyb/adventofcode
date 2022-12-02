package input

import (
	"bufio"
	_ "embed"
	"strings"

	"github.com/gondyb/adventofcode/2022/day_02/model"
)

//go:embed input.txt
var inputData string

func ParseInput() ([]model.Round, error) {
	scanner := bufio.NewScanner(strings.NewReader(inputData))

	var rounds []model.Round
	for scanner.Scan() {
		line := scanner.Text()

		r := model.RoundFromLine(line)

		rounds = append(rounds, r)
	}

	return rounds, nil
}
