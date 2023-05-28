package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type MyImage struct{
	image [][]uint8
}

func (i MyImage) ColorModel() color.Model {
	return color.NRGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	width := len(i.image)
	height := len(i.image[0])

	return image.Rect(0, 0, width, height)
}

func (i MyImage) At(x, y int) color.Color {
	return color.RGBA{i.image[x][y], i.image[x][y], 255, 255}
}

func getMyPic(width, height int) MyImage {
	matrix := make([][]uint8, height)

	for rowKey := range matrix {
		matrix[rowKey] = make([]uint8, width)
		
		for columnKey := range matrix[rowKey] {
			matrix[rowKey][columnKey] = uint8(columnKey * 4 + columnKey * rowKey / 2 + 1000)
			// matrix[row_key][column_key] = uint8((column_key + row_key) / 2)
			// matrix[row_key][column_key] = uint8(column_key * 2 + row_key / 2)
			// matrix[row_key][column_key] = uint8(math.Pow(float64(row_key), float64(column_key)))
		}
	}

	return MyImage{
		image: matrix,
	}
}

func main() {
	m := getMyPic(256, 512)
	pic.ShowImage(m)
}
