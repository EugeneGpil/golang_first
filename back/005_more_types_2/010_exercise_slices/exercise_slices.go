package main

import "golang.org/x/tour/pic"

// import "math"

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)

	for row_key := range matrix {
		matrix[row_key] = make([]uint8, dx)
		
		for column_key := range matrix[row_key] {
			matrix[row_key][column_key] = uint8(column_key * row_key)
			// matrix[row_key][column_key] = uint8((column_key + row_key) / 2)
			// matrix[row_key][column_key] = uint8(column_key * 2 + row_key / 2)
			// matrix[row_key][column_key] = uint8(math.Pow(float64(row_key), float64(column_key)))
		}
	}

	return matrix
}
