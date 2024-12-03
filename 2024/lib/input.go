package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputFile struct {
	filename string
}

func LoadInputFile(day int) *InputFile {
	file := fmt.Sprintf("./day-%d/input.txt", day)

	return &InputFile{file}
}

func LoadExampleFile(day int) *InputFile {
	file := fmt.Sprintf("./day-%d/example.txt", day)

	return &InputFile{file}
}

func (f *InputFile) Strings() ([]string, error) {
	file, err := os.Open(f.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (f *InputFile) Ints() ([]int, error) {
	lines, err := f.Strings()
	if err != nil {
		return nil, err
	}

	var ints []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if (err != nil) {
			return nil, err
		}
		ints = append(ints, num)
	}

	return ints, nil
}

func (f *InputFile) Floats() ([]float64, error) {
	lines, err := f.Strings()
	if err != nil {
		return nil, err
	}

	var floats []float64
	for _, line := range lines {
		num, err := strconv.ParseFloat(line, 64)
		if (err != nil) {
			return nil, err
		}
		floats = append(floats, num)
	}

	return floats, nil
}
