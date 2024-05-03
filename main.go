package main

import (
	"fmt"
	"main/integers"
	strct "main/struct"
)

func main() {
	fmt.Println("%d", integers.Add(2, 2))
	c := strct.Circle{R: 10.0}
	fmt.Println("", c.Perimeter())
	fmt.Println("", c.Area())
}
