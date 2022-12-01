package model

type Meal struct {
	Calories int
}

type Elf struct {
	Meals []Meal
}

func (e Elf) GetTotalCalories() int {
	var sum int
	for _, meal := range e.Meals {
		sum += meal.Calories
	}

	return sum
}
