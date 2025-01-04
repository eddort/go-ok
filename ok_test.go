package ok

import (
	"errors"
	"testing"
)

func TestVal(t *testing.T) {
	result := Val(42)
	if result.Value != 42 || result.Err != nil {
		t.Errorf("Val failed: expected value 42 with no error, got %v and error %v", result.Value, result.Err)
	}
}

func TestErr(t *testing.T) {
	testErr := errors.New("test error")
	result := Err[int](testErr)
	if result.Value != 0 || result.Err != testErr {
		t.Errorf("Err failed: expected value 0 with error %v, got value %v and error %v", testErr, result.Value, result.Err)
	}
}

func TestFrom(t *testing.T) {
	// Testing successful value creation
	result := From(42, nil)
	if result.Value != 42 || result.Err != nil {
		t.Errorf("From failed: expected value 42 with no error, got %v and error %v", result.Value, result.Err)
	}

	// Testing error creation
	testErr := errors.New("test error")
	result = From[int](0, testErr)
	if result.Err != testErr {
		t.Errorf("From failed: expected error %v, got %v", testErr, result.Err)
	}
}

func TestTry(t *testing.T) {
	initialResult := Val(10)
	result := Try(initialResult, func(x int) Result[int] {
		return Val(x * 2)
	})

	if result.Value != 20 || result.Err != nil {
		t.Errorf("Try failed: expected value 20 with no error, got %v and error %v", result.Value, result.Err)
	}

	initialResult = Err[int](errors.New("initial error"))
	result = Try(initialResult, func(x int) Result[int] {
		return Val(x * 2)
	}, "context")

	if result.Err == nil || result.Value != 0 {
		t.Errorf("Try failed with error handling: expected 0 value with error, got %v and error %v", result.Value, result.Err)
	}
}

func TestUnwrap(t *testing.T) {
	// Test case 1: Successful value
	result := Val(123)
	v, err := result.Unwrap()
	if v != 123 || err != nil {
		t.Errorf("Unwrap failed: expected 123 and nil error, got value: %v, error: %v", v, err)
	}

	// Test case 2: Error case
	result = Err[int](errors.New("unwrap error"))
	v, err = result.Unwrap()
	if v != 0 || err == nil { // Expect zero value and non-nil error
		t.Errorf("Unwrap failed: expected zero value and non-nil error, got value: %v, error: %v", v, err)
	}
}

// Test TryFrom - handles (value, error) return from callback
func TestTryFrom(t *testing.T) {
	initialResult := Val(10)
	result := TryFrom(initialResult, func(x int) (int, error) {
		return x * 2, nil
	})

	if result.Value != 20 || result.Err != nil {
		t.Errorf("TryFrom failed: expected value 20 with no error, got %v and error %v", result.Value, result.Err)
	}

	initialResult = Err[int](errors.New("initial error"))
	result = TryFrom(initialResult, func(x int) (int, error) {
		return x * 2, nil
	}, "context failed")

	if result.Err == nil || result.Value != 0 {
		t.Errorf("TryFrom failed with error handling: expected 0 value with error, got %v and error %v", result.Value, result.Err)
	}
}

// Test TryVal - only value return from callback
func TestTryVal(t *testing.T) {
	initialResult := Val(5)
	result := TryVal(initialResult, func(x int) int {
		return x * 2
	})

	if result.Value != 10 || result.Err != nil {
		t.Errorf("TryVal failed: expected value 10 with no error, got %v and error %v", result.Value, result.Err)
	}

	initialResult = Err[int](errors.New("initial error"))
	result = TryVal(initialResult, func(x int) int {
		return x * 2
	})

	if result.Err == nil || result.Value != 0 {
		t.Errorf("TryVal failed with error handling: expected error, got %v and error %v", result.Value, result.Err)
	}
}

// Test TryErr - only error return from callback
func TestTryErr(t *testing.T) {
	initialResult := Val(8)
	result := TryErr(initialResult, func(x int) error {
		return nil
	})

	if result.Err != nil {
		t.Errorf("TryErr failed: expected no error, got error %v", result.Err)
	}

	result = TryErr(initialResult, func(x int) error {
		return errors.New("new error")
	}, "failure occurred")
	// || result.Value != 8
	if result.Err == nil {
		t.Errorf("TryErr failed with error generation: expected new error, got value %v and error %v", result.Value, result.Err)
	}

	initialResult = Err[int](errors.New("initial error"))
	result = TryErr(initialResult, func(x int) error {
		return nil
	})

	if result.Err == nil {
		t.Errorf("TryErr failed with initial error handling: expected initial error, got no error")
	}
}
