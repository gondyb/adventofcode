package main

import (
	"fmt"
	"sort"

	"github.com/gondyb/adventofcode/2022/day_01/input"
)

func main() {
	elfs, err := input.ParseInput()
	if err != nil {
		panic(err)
	}

	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].GetTotalCalories() < elfs[j].GetTotalCalories()
	})

	fmt.Printf("first three elfs have %d cals\n", last(elfs, 1).GetTotalCalories()+last(elfs, 2).GetTotalCalories()+last(elfs, 3).GetTotalCalories())
}

func last[E any](s []E, i int) E {
	if len(s) == 0 {
		var zero E
		return zero
	}
	return s[len(s)-i]
}
