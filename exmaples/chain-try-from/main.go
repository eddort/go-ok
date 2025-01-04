package main

import (
	"fmt"

	"github.com/eddort/go-ok"
)

func addTen(x int) (int, error) {
	return x + 10, nil
}

func divideByTwo(x int) (int, error) {
	if x%2 != 0 {
		return 0, fmt.Errorf("value must be even to divide by two")
	}
	return x / 2, nil
}

func main() {
	initial := ok.Val(20)
	result := ok.TryFrom(initial, addTen, "Failed to add ten")
	result = ok.TryFrom(result, divideByTwo, "Failed to divide by two")

	if result.Err != nil {
		fmt.Printf("Error: %s\n", result.Err)
	} else {
		fmt.Printf("Final Result: %d\n", result.Value)
	}
}
