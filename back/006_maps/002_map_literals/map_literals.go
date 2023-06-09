package main

import "fmt"

type Vertex struct {
	Lat, Lon float64
}

func main() {
	m := map[string]Vertex{
		"Bell Labs": {
			40.68433,
			-74.39967,
		},
		"Google": {
			37.42202,
			-122.08408,
		},
	}

	fmt.Println(m)
}
