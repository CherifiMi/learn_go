package main

import "fmt"

func main() {
	fmt.Printf(hello(""))
}

const hellopre = "hello, "

func hello(s string) string {
	return hellopre + s
}
