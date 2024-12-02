package lib

import "fmt"

type QuickLogger struct {
	day int
	part int
}

func NewQuickLogger(day, part int) *QuickLogger {
	return &QuickLogger{day, part}
}

func (ql *QuickLogger) Title() {
	fmt.Printf(" ❄︎ Running day %d part %d ❆ \n", ql.day, ql.part)
}

func (ql *QuickLogger) Solve(solution interface{}) {
	fmt.Printf("Solution to Advent of Code 2024 day %d part %d is: '%v'\n", ql.day, ql.part, solution)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
