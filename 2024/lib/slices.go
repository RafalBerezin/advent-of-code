package lib

func CloneMatrix[T any](pSource *[][]T) *[][]T {
	source := *pSource
	copy := make([][]T, len(source))

	for row, rowData := range source {
		copy[row] = make([]T, len(rowData))
		for col, colData := range rowData {
			copy[row][col] = colData
		}
	}

	return &copy
}

func SetRange[T any](slice []T, start, size int, value T) {
	for i := 0; i < size; i++ {
		slice[start+i] = value
	}
}

type Tuple[L any, R any] struct {
	L L
	R R
}

func Zip[L any, R any](left *[]L, right *[]R) *[]Tuple[L, R] {
	maxLen := len(*left)
	if _r := len(*right); _r < maxLen {
		maxLen = _r
	}

	zipped := make([]Tuple[L, R], maxLen)

	for i := range zipped {
		zipped[i] = Tuple[L, R]{(*left)[i], (*right)[i]}
	}

	return &zipped
}

func Unzip[L any, R any](zipped *[]Tuple[L, R]) (*[]L, *[]R) {
	left := make([]L, 0, len(*zipped))
	right := make([]R, 0, len(*zipped))

	for i, tuple := range *zipped {
		left[i] = tuple.L
		right[i] = tuple.R
	}

	return &left, &right
}

func Transpose[T any](pMatrix *[][]T) *[][]T {
	matrix := *pMatrix
	height := len(matrix)
	width := len(matrix[0])

	for _, row := range matrix {
		if len(row) != width {
			panic("matrix is not a rectangle")
		}
	}

	transposed := make([][]T, width)
	for i := range transposed {
		transposed[i] = make([]T, height)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return &transposed
}

func TransposeInPlace[T any](pMatrix *[][]T) {
	matrix := *pMatrix
	height := len(matrix)
	width := len(matrix[0])

	if width != height {
		panic("matrix is not a square")
	}

	for i := 0; i < height; i++ {
		for j := i + 1; j < width; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}
