package main

import (
	"fmt"

	"github.com/eddort/go-ok"
)

func validateProduct(id int) (int, error) {
	if id <= 0 {
		return 0, fmt.Errorf("invalid product ID")
	}
	return id, nil
}

func calculateTax(id int) (float64, error) {
	var taxRate float64
	switch id {
	case 1:
		taxRate = 0.075
	case 2:
		taxRate = 0.085
	default:
		taxRate = 0.05
	}
	return 100 + 100*taxRate, nil
}

func applyDiscount(price float64) (float64, error) {
	if price > 100 {
		return price * 0.8, nil
	}
	return price, nil
}

func processPayment(price float64) error {
	if price > 500 {
		return fmt.Errorf("payment declined")
	}
	return nil
}

func processOrder(productId int) (float64, error) {
	id := ok.From(validateProduct(productId))
	taxedPrice := ok.TryFrom(id, calculateTax)
	discountedPrice := ok.TryFrom(taxedPrice, applyDiscount)
	payment := ok.TryErr(discountedPrice, processPayment)
	return discountedPrice.Value, payment.Err
}

func main() {
	_, err := processOrder(10)
	if err != nil {
		fmt.Printf("Error processing order: %s\n", err)
	} else {
		fmt.Println("Order processed successfully")
	}
}
