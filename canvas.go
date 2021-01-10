package main

import (
	"github.com/schleising/canvas"
)

func main() {
	canvas := canvas.Initialise("canvas")

	canvas.SetFillStyle("#FF8800")

	canvas.FillRect(10, 20, 100, 200)
}
