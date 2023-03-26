package server

import (
	"testing"
)

// Unit tests to test the calculate method.
func TestLegs(t *testing.T) {

	tests := []struct {
		trip          [][]string
		expectedStart string
		expectedEnd   string
		expectError   bool
	}{
		{
			[][]string{
				{"SFO", "EWR"},
			}, "SFO", "EWR", false},
		{
			[][]string{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
			}, "SFO", "EWR", false},
		{
			[][]string{
				{"IND", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			}, "SFO", "EWR", false},
	}

	for i, tt := range tests {
		actualStart, actualEnd, err := calculateRoute(tt.trip)
		if actualStart != tt.expectedStart {
			t.Fatalf("tests[%d] - wrong start. expected=%q, got=%q",
				i, tt.expectedStart, actualStart)
		}
		if actualEnd != tt.expectedEnd {
			t.Fatalf("tests[%d] - wrong end. expected=%q, got=%q",
				i, tt.expectedEnd, actualEnd)
		}

		if err != nil && !tt.expectError {
			t.Fatalf("tests[%d] - unexpected error %q",
				i, err)
		}

		if err == nil && tt.expectError {
			t.Fatalf("tests[%d] - expect error  but got nil", i)
		}

	}
}
