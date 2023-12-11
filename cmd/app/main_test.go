package main

import "testing"

// TestFactorial tests the Factorial function with various inputs.
func TestFactorialPositive(t *testing.T) {
	result, err := Factorial(5)
	if err != nil {
		t.Errorf("Factorial(5) unexpected error %v", err)
	}
	if result != 120 {
		t.Errorf("Factorial(5) = %d; want 120", result)
	}
}

func TestFactorialZero(t *testing.T) {
	result, err := Factorial(0)
	if err != nil {
		t.Errorf("Factorial(0) unexpected error %v", err)
	}
	if result != 1 {
		t.Errorf("Factorial(1) = %d; want 1", result)
	}
}

func TestFactorialNegative(t *testing.T) {
	result, err := Factorial(-5)
	if err == nil {
		t.Errorf("Factorial(-5) expected an error but got none")
	}
	if result != 0 {
		t.Errorf("Factorial(-5) = %d; want 0", result)
	}
}
