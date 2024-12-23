package lib

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func InBounds2D(row, col, height, width int) bool {
	return 0 <= row && row < height && 0 <= col && col < width
}
