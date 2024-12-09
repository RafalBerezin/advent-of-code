package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RafalBerezin/advent-of-code/2024/day1"
	"github.com/RafalBerezin/advent-of-code/2024/day2"
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
	Part1 func(* lib.InputFile) any
	Part2 func(* lib.InputFile) any
}

var days = map[string]*day {
	"1": {day1.Part1, day1.Part2},
	"2": {day2.Part1, day2.Part2},
	"3": {day3.Part1, day3.Part2},
	"4": {day4.Part1, day4.Part2},
	"5": {day5.Part1, day5.Part2},
	"6": {day6.Part1, day6.Part2},
	"7": {day7.Part1, day7.Part2},
	"8": {day8.Part1, day8.Part2},
	"9": {day9.Part1, day9.Part2},
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

	day := days[dayStr]
	if day == nil {
		fmt.Printf("No solution for day: '%v'\n", dayStr)
		return
	}

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

	fmt.Printf("Part should be either '1' or '2' (or don't specify to run both)\n\n" + usageInfo)
}

func runPart(day *day, part rune, input *lib.InputFile) {
	fmt.Printf("┌ Running part %v\n", string(part))

	defer func(){
		if r := recover(); r != nil {
			fmt.Printf("└ ❌ An error occured:\n%v\n", r)
		}
	}()

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
