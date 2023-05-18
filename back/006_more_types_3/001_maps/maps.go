package main

import "fmt"

type Vertex struct {
	Lat, Lon float64
}

func main() {
	m := make(map[string]Vertex)
	m["SOME KEY"] = Vertex{
		40.5,
		30.5,
	}
	fmt.Println(m)
}
