package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 10, 2, 5, false},
	{"invalid-data", 10, 0, 0, true},
}

//go test -coverprofile=coverage.out && go tool cover -html=coverage.out

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one")
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}
