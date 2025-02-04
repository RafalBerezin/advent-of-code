package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RafalBerezin/advent-of-code/2024/day1"
	"github.com/RafalBerezin/advent-of-code/2024/day10"
	"github.com/RafalBerezin/advent-of-code/2024/day11"
	"github.com/RafalBerezin/advent-of-code/2024/day12"
	"github.com/RafalBerezin/advent-of-code/2024/day13"
	"github.com/RafalBerezin/advent-of-code/2024/day14"
	"github.com/RafalBerezin/advent-of-code/2024/day15"
	"github.com/RafalBerezin/advent-of-code/2024/day16"
	"github.com/RafalBerezin/advent-of-code/2024/day17"
	"github.com/RafalBerezin/advent-of-code/2024/day18"
	"github.com/RafalBerezin/advent-of-code/2024/day19"
	"github.com/RafalBerezin/advent-of-code/2024/day2"
	"github.com/RafalBerezin/advent-of-code/2024/day20"
	"github.com/RafalBerezin/advent-of-code/2024/day21"
	"github.com/RafalBerezin/advent-of-code/2024/day22"
	"github.com/RafalBerezin/advent-of-code/2024/day23"
	"github.com/RafalBerezin/advent-of-code/2024/day24"
	"github.com/RafalBerezin/advent-of-code/2024/day25"
	"github.com/RafalBerezin/advent-of-code/2024/day3"
	"github.com/RafalBerezin/advent-of-code/2024/day4"
	"github.com/RafalBerezin/advent-of-code/2024/day5"
	"github.com/RafalBerezin/advent-of-code/2024/day6"
	"github.com/RafalBerezin/advent-of-code/2024/day7"
	"github.com/RafalBerezin/advent-of-code/2024/day8"
	"github.com/RafalBerezin/advent-of-code/2024/day9"
	"github.com/RafalBerezin/advent-of-code/2024/lib"
)

type day struct {
	Part1 func(file *lib.InputFile) any
	Part2 func(file *lib.InputFile) any
}

var days = []*day {
	{day1.Part1, day1.Part2},
	{day2.Part1, day2.Part2},
	{day3.Part1, day3.Part2},
	{day4.Part1, day4.Part2},
	{day5.Part1, day5.Part2},
	{day6.Part1, day6.Part2},
	{day7.Part1, day7.Part2},
	{day8.Part1, day8.Part2},
	{day9.Part1, day9.Part2},
	{day10.Part1, day10.Part2},
	{day11.Part1, day11.Part2},
	{day12.Part1, day12.Part2},
	{day13.Part1, day13.Part2},
	{day14.Part1, day14.Part2},
	{day15.Part1, day15.Part2},
	{day16.Part1, day16.Part2},
	{day17.Part1, day17.Part2},
	{day18.Part1, day18.Part2},
	{day19.Part1, day19.Part2},
	{day20.Part1, day20.Part2},
	{day21.Part1, day21.Part2},
	{day22.Part1, day22.Part2},
	{day23.Part1, day23.Part2},
	{day24.Part1, day24.Part2},
	{day25.Part1, day25.Part2},
}

var usageInfo = "Usage: go run main.go <day> [<part>] [-e]\nUse '-e' flag to use the example input\n\nExample: 'go run main.go 3 2 -e'\n - runs day 3 part 2 with example input"

func main() {
	args := os.Args[1:]
	argsCount := len(args)

	if (argsCount == 0) {
		log.Fatal(usageInfo)
	}

	dayStr := args[0]
	dayNum, err := strconv.Atoi(dayStr)
	if err != nil || dayNum < 1 || dayNum > 25 {
		fmt.Println("Day should be a number between 1 and 25")
		return
	}
	dayNum--

	if dayNum >= len(days) {
		fmt.Printf("No solution for day: '%v'\n", dayStr)
		return
	}
	day := days[dayNum]

	title := fmt.Sprintf("Running ❄︎ Advent of Code ❄︎ 2024 day %v", dayStr)
	titleBorder := strings.Repeat("─", len(title))
	fmt.Printf("%v\n   %v\n%v\n", titleBorder, title, titleBorder)

	input := lib.LoadFile(dayStr)
	if argsCount > 1 && args[len(args)-1] == "-e" {
		fmt.Println("  Using example input")
		input = input.Example()
		argsCount--
	}

	if argsCount < 2 {
		runPart(day, '1', input)
		runPart(day, '2', input)
		return
	}

	part := args[1][0]
	if part == '1' || part == '2' {
		runPart(day, rune(part), input)
		return
	}

	fmt.Println("Part should be either '1' or '2' (or don't specify to run both)\n\n" + usageInfo)
}

func runPart(day *day, part rune, input *lib.InputFile) {
	fmt.Printf("┌ Running part %v\n", string(part))

	// defer func(){
	// 	if r := recover(); r != nil {
	// 		fmt.Printf("└ ❌ An error occured:\n%v\n", r)
	// 	}
	// }()

	startTime := time.Now()

	var solution any
	if part == '1' {
		solution = day.Part1(input)
	} else {
		solution = day.Part2(input)
	}

	elapsed := time.Since(startTime)

	fmt.Printf("├ ✔ Solution: %v\n", solution)
	fmt.Printf("└ 🕒 Time: %s\n", elapsed)
}
