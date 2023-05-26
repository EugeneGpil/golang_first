package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	readed, err := r.r.Read(b)

	len := len(b)
	for i := 0; i < len; i++ {
		b[i] = getUncoded13(b[i])
	}

	return readed, err
}

func getUncoded13(b byte) byte {
	if isNeedAdd13(b) {
		return b + 13
	}

	if isNeedReduseBy13(b) {
		return b - 13
	}

	return b
}

func isNeedAdd13(b byte) bool {
	isPlus13Uppercase := b >= byte('A') && b <= byte('M')
	isPlus13Lowcase := b >= byte('a') && b <= byte('m')

	return isPlus13Lowcase || isPlus13Uppercase
}

func isNeedReduseBy13(b byte) bool {
	isMinus13Uppercase := b >= byte('N') && b <= byte('Z')
	isMinus13Lowcase := b >= byte('n') && b <= byte('z')

	return isMinus13Lowcase || isMinus13Uppercase
}

func main() {
	r := rot13Reader{
		r: strings.NewReader("Lbh penpxrq gur pbqr!"),
	}
	io.Copy(os.Stdout, &r)
	fmt.Println("")
}
