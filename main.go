package main

import (
	"egyptian/egyptian"
)

func main() {
	x, y := egyptian.ReadInput()
	egyptian.ValidateInput(x, y)
	result := egyptian.Calculate(x, y)
	egyptian.DisplayResult(result)
}
