package ok

import (
	"errors"
)

// Result — структура для хранения значения или ошибки
type Result[T any] struct {
	Value T
	Err   error
}

// Val — создаёт Result с успешным значением
func Val[T any](value T) Result[T] {
	return Result[T]{Value: value, Err: nil}
}

// Err — создаёт Result с ошибкой
func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{Value: zero, Err: err}
}

// From — создаёт Result из значения и ошибки
func From[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Val(value)
}

// Вспомогательная функция для добавления контекста к ошибке
func wrapContext(err error, context ...string) error {
	if len(context) > 0 {
		return errors.Join(errors.New(context[0]), err)
	}
	return err
}

// Try — для функций, возвращающих Result
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

// TryFrom — для функций, возвращающих (значение, ошибку)
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

// TryVal — для функций, возвращающих только значение
func TryVal[T, U any](r Result[T], cb func(T) U) Result[U] {
	if r.Err != nil {
		return Err[U](r.Err)
	}
	return Val(cb(r.Value))
}

// TryErr — для функций, возвращающих только ошибку
func TryErr[T any](r Result[T], cb func(T) error, context ...string) Result[T] {
	if r.Err != nil {
		return Err[T](wrapContext(r.Err, context...))
	}
	if err := cb(r.Value); err != nil {
		return Err[T](wrapContext(err, context...))
	}
	return r
}

// Unwrap — извлекает значение, вызывает панику в случае ошибки
func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic("Unwrap failed: " + r.Err.Error())
	}
	return r.Value
}

// UnwrapOr — извлекает значение или возвращает значение по умолчанию
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.Err != nil {
		return defaultValue
	}
	return r.Value
}

// UnwrapOrElse — извлекает значение или вызывает функцию для получения значения
func (r Result[T]) UnwrapOrElse(fallback func() T) T {
	if r.Err != nil {
		return fallback()
	}
	return r.Value
}
