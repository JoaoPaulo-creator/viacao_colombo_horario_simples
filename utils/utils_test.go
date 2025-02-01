package utils

import (
	"testing"
)

func TestIfStrContainsSubStr(t *testing.T) {
	bd := []string{"Sunday", "Tuesday", "Friday"}
	subStr := "Friday"

	contains := Contains(bd, subStr)

	if !contains {
		t.Errorf("contains deveria ser true, mas recebeu o valor false")
	}
}

func TestIfStrNotContainsSubStr(t *testing.T) {
	bd := []string{"Sunday", "Tuesday", "Friday"}
	subStr := "Monday"

	contains := Contains(bd, subStr)

	if contains {
		t.Errorf("contains deveria ser false, mas recebeu o valor true")
	}
}
