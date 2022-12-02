package main

import (
	"fmt"

	"github.com/gondyb/adventofcode/2022/day_02/input"
)

func main() {
	rounds, _ := input.ParseInput()

	points := 0

	for _, round := range rounds {
		points += round.GetPoints()
	}

	fmt.Printf("You won %d points\n", points)
}
