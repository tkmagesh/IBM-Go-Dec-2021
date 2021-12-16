package utils

import "testing"

func TestIsPrime(t *testing.T) {
	//arrange
	var no int = 97
	var expected bool = true

	//act
	actual := IsPrime(no)

	//assert
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
