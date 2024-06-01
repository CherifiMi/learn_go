package main

import (
	"fmt"
	"math"
)

const (
	center      = 150
	hour_hand   = 50
	minute_hand = 80
	second_hand = 90
)

func main() {
	/*
			x = r cos  phi + offset
			y = r sin  phi + offset

		-2 pi

		2pi -> 100

		12 -> 100
		6 -> 50%

		phi = (time * 100 / 12)*2pi

	*/

	//h := time.Now().Hour() / 2

	phi := timeToPhi(6, 12)

	x := hour_hand*math.Cos(phi) + 150
	y := hour_hand*math.Sin(phi) - 150

	//m := time.Now().Minute()
	//s := time.Now().Second()

	fmt.Printf("x is %f, y is %f, phi %f", x, y, phi)
}

func timeToPhi(time, i int) float64 {
	pt := float64(time) / float64(i)
	return (pt * 2 * math.Pi) - 1.5708
}
