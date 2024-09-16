package math

import "testing"

const defaultError = "Valor esperado %v, mas o resultado encontrado foi %v"

func TestCalculateMedia(t *testing.T) {
	expected := 6.0
	result := CalculateMedia(4, 6, 8)

	if result != expected {
		t.Errorf(defaultError, expected, result)
	}
}
