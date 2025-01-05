# go-ok

`go-ok` is a lightweight Go library for error handling inspired by Rust's `Result<T, E>` type. It brings the safety and convenience of Rust-style error management to Go, making it easier to handle errors in complex workflows without sacrificing readability or correctness.

---

## Why Use `go-ok`?

Error handling in Go traditionally relies on multiple `if err != nil` checks after every function call. While this approach is explicit, it can lead to:

1. **Repetitive Boilerplate Code:** Each operation requires explicit error handling, which can clutter the code.
2. **Human Error:** Go does not enforce error checking, making it easy to accidentally ignore errors.
3. **Nested Logic:** Complex workflows often result in deeply nested code, reducing readability.

### The `go-ok` Advantage:

`go-ok` addresses these issues by introducing a `Result` type, similar to Rust's `Result<T, E>`, and a set of utilities for propagating errors seamlessly. It enforces error handling through chaining while keeping code flat and readable. This approach combines Go's simplicity with Rust's safety guarantees.

---

## Installation

To install the library, run:

```bash
go get github.com/eddort/go-ok
```

---

## Example: Traditional Go vs. `go-ok`

### Common Functions

```go
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
```

### Without `go-ok`

```go
func processOrder(productId int) (float64, error) {
    id, err := validateProduct(productId)
    if err != nil {
        return 0, err
    }

    taxedPrice, err := calculateTax(id)
    if err != nil {
        return 0, err
    }

    discountedPrice, err := applyDiscount(taxedPrice)
    if err != nil {
        return 0, err
    }

    _, err = processPayment(discountedPrice)
    if err != nil {
        return 0, err
    }

    return discountedPrice, nil
}
```

Issues:

- Errors are manually checked after each step.
- Repetitive error handling increases boilerplate code.
- Nested logic reduces readability and makes workflows harder to maintain.

### With `go-ok`

```go
func processOrder(productId int) (float64, error) {
	id := ok.From(validateProduct(productId))
	taxedPrice := ok.TryFrom(id, calculateTax)
	discountedPrice := ok.TryFrom(taxedPrice, applyDiscount)
	payment := ok.TryErr(discountedPrice, processPayment)
	return discountedPrice.Value, payment.Err
}
```

Advantages:

- **Flat structure:** No deep nesting, even with multiple steps.
- **Automatic error propagation:** Errors stop the chain immediately, returning the first encountered error.
- **Safety enforced by type system:** Results must be explicitly handled.

---

## Key Features

### 1. Rust-Like `Result` Type

Encapsulates either a value (`Value`) or an error (`Err`), enforcing explicit error handling without boilerplate.

### 2. `Try` for Seamless Chaining

Simplifies error propagation by automatically passing results or short-circuiting on errors.

### 3. Functional Composition

Supports higher-order functions, enabling flexible and modular workflows.

---

## API Reference

### Constructors

- **`ok.Val(value T)`** - Creates a successful `Result`.
- **`ok.Err(err error)`** - Creates an error `Result`.

### Core Methods

- **`ok.Try(result, cb)`** - Chains operations, propagating errors automatically.
- **`ok.TryFrom(result, cb)`** - Supports functions returning `(value, error)`.
- **`ok.TryVal(result, cb)`** - For pure value-returning functions.
- **`ok.TryErr(result, cb)`** - For functions returning only errors.

### Accessors

- **`result.Unwrap()`** - Extracts the value and error without panicking.
- **`result.UnwrapOr(defaultValue)`** - Returns the value or a default.
- **`result.UnwrapOrElse(func)`** - Calls a fallback function if an error exists.

---

## When Should You Use `go-ok`?

- **Complex workflows:** Chains of dependent operations where each step might fail.
- **Readable code:** When you need flat, declarative error handling without nesting.
- **Safety enforcement:** To avoid accidentally ignoring errors, making your code more robust.

If you value the safety guarantees of Rust and want to bring similar patterns into Go, `go-ok` is an excellent fit for your project.

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve `go-ok`.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

