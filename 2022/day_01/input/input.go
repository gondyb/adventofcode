package input

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/gondyb/adventofcode/2022/day_01/model"
)

//go:embed input.txt
var inputData string

func ParseInput() ([]model.Elf, error) {
	scanner := bufio.NewScanner(strings.NewReader(inputData))

	var elfs []model.Elf
	var currentElf model.Elf
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfs = append(elfs, currentElf)
			currentElf = model.Elf{}
			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("unable to parse input %s %w", line, err)
		}

		meal := model.Meal{Calories: cals}
		currentElf.Meals = append(currentElf.Meals, meal)
	}

	return elfs, nil
}
