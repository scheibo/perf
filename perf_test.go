package perf

import (
	"testing"

	"github.com/scheibo/calc"
)

func TestCalcM(t *testing.T) {
	tests := []struct {
		t, d, gr, expected float64
	}{
		{965, 4800, 0.08125, 500},   // 16m05
		{766, 4800, 0.08125, 1000},  // 12m46
		{2260, 13910, 0.0790, 1000}, // 37m40
	}
	for _, tt := range tests {
		actual := CalcM(tt.t, tt.d, tt.gr)
		if !calc.Eqf(actual, tt.expected) {
			t.Errorf("CalcM(%.3f, %.3f, %.3f): got: %.3f, want: %.3f",
				tt.t, tt.d, tt.gr, actual, tt.expected)
		}
	}
}

func TestCalcF(t *testing.T) {
	tests := []struct {
		t, d, gr, expected float64
	}{
		{1113, 4800, 0.08125, 500},  // 18m33
		{884, 4800, 0.08125, 1000},  // 14m44
		{2579, 13910, 0.0790, 1000}, // 42m59

	}
	for _, tt := range tests {
		actual := CalcF(tt.t, tt.d, tt.gr)
		if !calc.Eqf(actual, tt.expected) {
			t.Errorf("CalcF(%.3f, %.3f, %.3f): got: %.3f, want: %.3f",
				tt.t, tt.d, tt.gr, actual, tt.expected)
		}
	}
}
