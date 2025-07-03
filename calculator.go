// Comentario a nivel de paquete
//
// Package calculator provides simple arithmetic operations.
//
// This package is designed to be easy to use and provides
// basic mathematical functions.
//
// Usage:
//
//	sum := calculator.Add(5, 3)
//	diff := calculator.Subtract(10, 4)
//
// For more details, refer to the individual function documentations.
package calculator

import (
	"errors" // Comentario para la importación (opcional, pero útil)
)

// ErrDivideByZero is returned when an attempt is made to divide by zero.
var ErrDivideByZero = errors.New("cannot divide by zero")

// Add returns the sum of two integers.
//
// It takes two integer arguments, a and b, and returns their sum.
// Example:
//
//	result := Add(10, 5) // result will be 15
//
// See also: Subtract
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers.
//
// It subtracts the second argument (b) from the first (a).
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers.
//
// This function performs a simple multiplication.
func Multiply(a, b int) int {
	return a * b
}

// Divide divides two integers and returns the result.
// It also returns an error if the divisor (b) is zero.
//
// Errors:
//   - ErrDivideByZero: If the divisor 'b' is 0.
//
// Example:
//
//	result, err := Divide(10, 2) // result = 5, err = nil
//	_, err = Divide(10, 0)     // result = 0, err = ErrDivideByZero
//
// Note: Integer division truncates the remainder.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// Calculator provides methods for performing arithmetic operations.
// It's a simple struct to demonstrate method documentation.
type Calculator struct {
	Name string // Name of the calculator instance.
}

// NewCalculator creates a new Calculator instance with a given name.
func NewCalculator(name string) *Calculator {
	return &Calculator{Name: name}
}

// PerformAddition adds two numbers using the Calculator instance.
// This is an example of a method documentation.
func (c *Calculator) PerformAddition(a, b int) int {
	return Add(a, b) // Delegates to the Add function
}

// internalHelper is a non-exported (private) function and will not appear
// in the GoDoc output by default, unless specifically requested.
func internalHelper(x int) int {
	return x * 2
}
