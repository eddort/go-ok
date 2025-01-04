package main

import (
	"errors"
	"fmt"

	"github.com/eddort/go-ok"
)

func multiplyByTwo(x int) ok.Result[int] {
	return ok.Val(x * 2)
}

func subtractFive(x int) ok.Result[int] {
	if x < 5 {
		return ok.Err[int](errors.New("result after subtraction must be non-negative"))
	}
	return ok.Val(x - 5)
}

func main() {
	initial := ok.Val(10)
	result := ok.Try(initial, multiplyByTwo)
	result = ok.Try(result, subtractFive, "Error after subtraction")

	if result.Err != nil {
		fmt.Printf("Error: %s\n", result.Err)
	} else {
		fmt.Printf("Final Result: %d\n", result.Value)
	}
}
