package ok

import (
	"errors"
)

// Result represents a container for either a value or an error.
type Result[T any] struct {
	Value T
	Err   error
}

// Val creates a Result with a successful value.
func Val[T any](value T) Result[T] {
	return Result[T]{Value: value, Err: nil}
}

// Err creates a Result containing an error.
// It initializes the value to the zero value of the specified type.
func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{Value: zero, Err: err}
}

// From creates a Result based on a value and an error.
// If the error is non-nil, it wraps the error in a Result.
func From[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Val(value)
}

// wrapContext adds context to an error by joining it with additional information.
func wrapContext(err error, context ...string) error {
	if len(context) > 0 {
		return errors.Join(errors.New(context[0]), err)
	}
	return err
}

// Try applies a function that returns a Result and propagates errors if present.
// It wraps errors with additional context if needed.
func Try[T, U any](r Result[T], cb func(T) Result[U], context ...string) Result[U] {
	if r.Err != nil {
		return Err[U](wrapContext(r.Err, context...))
	}
	res := cb(r.Value)
	if res.Err != nil {
		return Err[U](wrapContext(res.Err, context...))
	}
	return res
}

// TryFrom applies a function that returns a value and an error.
// It propagates errors and wraps them with context if necessary.
func TryFrom[T, U any](r Result[T], cb func(T) (U, error), context ...string) Result[U] {
	if r.Err != nil {
		return Err[U](wrapContext(r.Err, context...))
	}
	value, err := cb(r.Value)
	if err != nil {
		return Err[U](wrapContext(err, context...))
	}
	return Val(value)
}

// TryVal applies a function that returns only a value.
// Errors are propagated without modifying the result.
func TryVal[T, U any](r Result[T], cb func(T) U) Result[U] {
	if r.Err != nil {
		return Err[U](r.Err)
	}
	return Val(cb(r.Value))
}

// TryErr applies a function that returns only an error.
// It propagates errors with context if they occur.
func TryErr[T any](r Result[T], cb func(T) error, context ...string) Result[T] {
	if r.Err != nil {
		return Err[T](wrapContext(r.Err, context...))
	}
	if err := cb(r.Value); err != nil {
		return Err[T](wrapContext(err, context...))
	}
	return r
}

// Unwrap extracts the value, panicking if an error is present.
func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic("Unwrap failed: " + r.Err.Error())
	}
	return r.Value
}

// UnwrapOr extracts the value or returns a default value if an error is present.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.Err != nil {
		return defaultValue
	}
	return r.Value
}

// UnwrapOrElse extracts the value or calls a fallback function to produce a default value if an error is present.
func (r Result[T]) UnwrapOrElse(fallback func() T) T {
	if r.Err != nil {
		return fallback()
	}
	return r.Value
}
