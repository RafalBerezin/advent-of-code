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
