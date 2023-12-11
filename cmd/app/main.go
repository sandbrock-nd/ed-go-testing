package main

import (
	"errors"
	"fmt"
)

// Factorial calculates the factorial of a given number.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("negative input is not allowed")
	}

	if n == 0 {
		return 1, nil
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result, nil
}

func main() {
	result, err := Factorial(5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial of 5:", result) // Expected: 120
	}

	result, err = Factorial(-3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Factorial of -3:", result) // Expected: Error message
	}
}
