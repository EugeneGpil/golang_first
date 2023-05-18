package main

import "golang.org/x/tour/pic"

// import "math"

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)

	for rowKey := range matrix {
		matrix[rowKey] = make([]uint8, dx)
		
		for columnKey := range matrix[rowKey] {
			matrix[rowKey][columnKey] = uint8(columnKey * rowKey)
			// matrix[row_key][column_key] = uint8((column_key + row_key) / 2)
			// matrix[row_key][column_key] = uint8(column_key * 2 + row_key / 2)
			// matrix[row_key][column_key] = uint8(math.Pow(float64(row_key), float64(column_key)))
		}
	}

	return matrix
}
