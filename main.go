package main

import (
	"egyptian/egyptian"
)

func main() {
	for {
		x, y := egyptian.ReadInput()

		if !egyptian.ValidateInput(x, y) {
			break
		}

		result := egyptian.Calculate(x, y)
		egyptian.DisplayResult(result)

		if !egyptian.DoItAgain() {
			break
		}
	}

	egyptian.Finish()
}
