package main

import (
	"fmt"

	"github.com/eddort/go-ok"
)

func doubleIfPositive(x int) ok.Result[int] {
	if x <= 0 {
		return ok.Err[int](fmt.Errorf("value must be positive, got %d", x))
	}
	return ok.Val(x * 2)
}

func subtractThreeIfEven(x int) ok.Result[int] {
	if x%2 != 0 {
		return ok.Err[int](fmt.Errorf("value must be even, got %d", x))
	}
	return ok.Val(x - 3)
}

func incrementValue(x int) ok.Result[int] {
	return ok.Val(x + 1)
}

func processValue(x int) ok.Result[int] {
	return ok.Try(
		ok.Try(
			ok.Try(
				ok.Val(x),
				doubleIfPositive,
			),
			subtractThreeIfEven,
		),
		incrementValue,
	)
}

func main() {
	finalResult := processValue(10)
	if finalResult.Err != nil {
		fmt.Printf("Error during processing: %s\n", finalResult.Err)
	} else {
		fmt.Printf("Final Result: %d\n", finalResult.Value)
	}
}
