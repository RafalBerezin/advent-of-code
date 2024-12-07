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
