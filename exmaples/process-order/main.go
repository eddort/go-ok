package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/eddort/go-ok"
)

func checkAvailability(productId int) ok.Result[int] {
	if rand.Intn(10) < 2 { // 20% chance that the product is not available
		return ok.Err[int](fmt.Errorf("product %d is not available", productId))
	}
	return ok.Val(productId) // return productId for next step use
}

func calculateTax(price float64, state string) ok.Result[float64] {
	var taxRate float64
	switch state {
	case "CA":
		taxRate = 0.075
	case "NY":
		taxRate = 0.085
	default:
		taxRate = 0.05
	}
	totalPrice := price + price*taxRate
	return ok.Val(totalPrice) // return totalPrice for next step use
}

func applyDiscount(totalPrice float64, code string) ok.Result[float64] {
	if code == "DISCOUNT20" {
		return ok.Val(totalPrice * 0.8) // 20% discount
	}
	return ok.Val(totalPrice) // return discounted price for next step use
}

func processPayment(discountedPrice float64) ok.Result[bool] {
	if discountedPrice > 500 {
		return ok.Err[bool](errors.New("payment declined, amount too high"))
	}
	return ok.Val(true) // return payment success for the next step
}

func confirmOrder(success bool) ok.Result[bool] {
	if success {
		fmt.Println("Order confirmed")
		return ok.Val(true)
	}
	return ok.Err[bool](fmt.Errorf("order confirmation failed"))
}

func processOrder(productId int, price float64, state, discountCode string) ok.Result[bool] {
	availability := checkAvailability(productId)
	taxedPrice := ok.Try(availability, func(_ int) ok.Result[float64] {
		return calculateTax(price, state)
	})
	discountedPrice := ok.Try(taxedPrice, func(price float64) ok.Result[float64] {
		return applyDiscount(price, discountCode)
	})
	payment := ok.Try(discountedPrice, processPayment)
	finalConfirmation := ok.Try(payment, confirmOrder)
	return finalConfirmation
}

func main() {
	order := processOrder(10, 300, "CA", "DISCOUNT20")
	if order.Err != nil {
		fmt.Printf("Error processing order: %s\n", order.Err)
	} else {
		fmt.Println("Order processed successfully")
	}
}
