package egyptian

import (
	"fmt"
	"os"
)

func Calculate(x, y int) int {
	z := 0

	fmt.Printf("%d * %d\n", x, y)

	for x > 1 {
		if IsEven(x) {
			x /= 2
			y *= 2
		} else {
			x--
			z += y
		}

		if z > 0 {
			fmt.Printf("= %d * %d + %d\n", x, y, z)
		} else {
			fmt.Printf("= %d * %d\n", x, y)
		}
	}

	return x*y + z
}

func DisplayResult(result int) {
	fmt.Printf("= %d\n", result)
}

func DoItAgain() bool {
	var response string

	fmt.Print("Continue? (y/N): ")

	_, err := fmt.Scanln(&response)

	if err != nil {
		return false
	}

	return response == "y" || response == "Y"
}

func IsEven(number int) bool {
	return number%2 == 0
}

func IsOdd(number int) bool {
	return number%2 == 1
}

func ReadInput() (int, int) {
	var x, y int

	fmt.Print("🚀 Type two numbers and press Enter: ")
	_, err := fmt.Scanf("%d %d\n", &x, &y) // See https://go.dev/doc/effective_go#blank

	if err != nil { // See https://google.github.io/styleguide/go/decisions.html
		return 0, 0
	}

	return x, y
}

func ValidateInput(x, y int) {
	if x < 1 || y < 1 {
		fmt.Printf("💥 Check your numbers : [x => %d], [y => %d]", x, y)
		os.Exit(666)
	}
}
