package main

import "fmt"

func main() {
	fmt.Printf(hello(""))
}

const hellopre = "hello, "

func hello(s string) string {
	if s == "" {
		s = "world"
	}
	return hellopre + s
}
