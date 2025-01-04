package main

import (
	"errors"
	"fmt"

	"example.com/ok/ok"
)

func parseStr(v int) (string, error) {
	if v < 0 {
		return "", errors.New("negative value")
	}
	return fmt.Sprintf("Value: %d", v), nil
}

func main() {
	fmt.Println("Привет, мир!")
	numbers := ok.From(-1223, nil)

	parsed := ok.TryFrom(numbers, parseStr, "conversion failed")

	fmt.Println(parsed.Err)
	fmt.Println(parsed.Value)
}
