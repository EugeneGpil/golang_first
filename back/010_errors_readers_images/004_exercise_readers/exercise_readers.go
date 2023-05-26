package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	len := len(b)
	for i := 0; i < len; i++ {
		b[i] = byte('A')
	}
	return len, nil
}

func main() {
	reader.Validate(MyReader{})
}
