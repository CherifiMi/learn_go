package main

import "fmt"

func main() {
	fmt.Printf(hello("", ""))
}

const hellopre = "hello, "

func hello(s string, lang string) string {
	if s == "" {
		s = "world"
	}
	if lang == "spanish" {
		return "hola, " + s
	}
	return hellopre + s
}
