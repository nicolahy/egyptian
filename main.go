package main

import (
	"egyptian/egyptian"
)

func main() {
	for {
		x, y := egyptian.ReadInput()
		egyptian.ValidateInput(x, y)
		result := egyptian.Calculate(x, y)
		egyptian.DisplayResult(result)
		if !egyptian.DoItAgain() {
			break
		}
	}
}
