package main

import (
	"fmt"

	"github.com/eddort/go-ok"
)

func square(x int) int {
	return x * x
}

func validatePositive(x int) error {
	if x <= 0 {
		return fmt.Errorf("number must be positive")
	}
	return nil
}

func main() {
	initial := ok.Val(4)
	result := ok.TryVal(initial, square)
	result = ok.TryErr(result, validatePositive, "Validation failed")

	if result.Err != nil {
		fmt.Printf("Error: %s\n", result.Err)
	} else {
		fmt.Printf("Final Result: %d\n", result.Value)
	}
}
