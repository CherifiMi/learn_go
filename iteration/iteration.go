package main

import "fmt"

func main() {
	fmt.Printf(Repeat("b", 7))
}

func Repeat(s string, repeatedCount int) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += s
	}
	return repeated
}
