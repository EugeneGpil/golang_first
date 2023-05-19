package main

import "golang.org/x/tour/wc"
import "strings"

func WordCount(s string) map[string]int {
	keys := strings.Fields(s)

	result := make(map[string]int)

	for _, key := range keys {
		result[key] = result[key] + 1
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
