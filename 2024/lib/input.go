package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputFile struct {
	dir string
	file string
}

func LoadFile(day string) *InputFile {
	dir := fmt.Sprintf("./day%v/", day)

	return &InputFile{dir, "input.txt"}
}

func (f *InputFile) Input() *InputFile {
	return &InputFile{f.dir, "input.txt"}
}

func (f *InputFile) Example() *InputFile {
	return &InputFile{f.dir, "example.txt"}
}

func (f *InputFile) Bytes() []byte {
	file, err := os.ReadFile(f.dir + f.file)
	CheckError(err)

	return file
}

func (f *InputFile) ByteGrid() [][]byte {
	file, err := os.Open(f.dir + f.file)
	CheckError(err)
	defer file.Close()

	var rows [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := make([]byte, len(scanner.Bytes()))
		copy(bytes, scanner.Bytes())
		rows = append(rows, bytes)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rows
}

func (f *InputFile) Digits() []int {
	bytes := f.Bytes()

	lastIndex := len(bytes)-1
	if bytes[lastIndex] == '\n' {
		bytes = bytes[:lastIndex]
	}

	digits := make([]int, len(bytes))
	for i, byte := range bytes {
		digits[i] = int(byte - '0')
	}

	return digits
}

func (f *InputFile) DigitGrid() [][]byte {
	byteGrid := f.ByteGrid()
	for i, row := range byteGrid {
		for j, col := range row {
			byteGrid[i][j] = col - '0'
		}
	}
	return byteGrid
}

func (f *InputFile) Strings() []string {
	file, err := os.Open(f.dir + f.file)
	CheckError(err)
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func (f *InputFile) Ints() []int {
	lines := f.Strings()

	var ints []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		CheckError(err)
		ints = append(ints, num)
	}

	return ints
}

func (f *InputFile) Floats() []float64 {
	lines := f.Strings()

	var floats []float64
	for _, line := range lines {
		num, err := strconv.ParseFloat(line, 64)
		CheckError(err)
		floats = append(floats, num)
	}

	return floats
}
