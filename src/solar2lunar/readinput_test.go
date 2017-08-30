package main

import (
	"testing"
)

func TestIsValidMonth(t *testing.T) {
	tests := []struct {
		description string
		inputMonth  int
		expStatus   bool
	}{
		{
			description: "Invalid month 0",
			inputMonth:  0,
			expStatus:   false,
		},
		{
			description: "valid month 1",
			inputMonth:  1,
			expStatus:   true,
		},
		{
			description: "valid month 12",
			inputMonth:  12,
			expStatus:   true,
		},
		{
			description: "invalid month 13",
			inputMonth:  13,
			expStatus:   false,
		},
	}

	for _, tc := range tests {
		t.Log("Test:", tc.description)
		status := isValidMonth(tc.inputMonth)
		if status != tc.expStatus {
			t.Errorf("Func returned wrong value: got %v want %v", status, tc.expStatus)
		}
	}

}
