package main

import "fmt"

func main() {
	fmt.Printf(hello(""))
}

func hello(s string) string {
	return "hello, " + s
}
