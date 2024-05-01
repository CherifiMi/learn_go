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

	pre := hellopre

	switch lang {
	case "spanish":
		pre = "hola, "
	case "franch":
		pre = "oui, "
	}

	return pre + s
}
