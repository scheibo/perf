package perf

import (
	"testing"

	"github.com/scheibo/calc"
)

func TestCalcM(t *testing.T) {
	tests := []struct {
		t, d, gr, expected float64
	}{
		{965, 4800, 0.08125, 500},   // OLH: 16m05
		{766, 4800, 0.08125, 1000},  // OLH: 12m46 (13m26)
		{2260, 13910, 0.0790, 1000}, // Huez: 37m40 (36m50)
		{1605, 9880, 0.0799, 1000},  // Gibraltr: 26m45 (27m12)
		{1904, 13100, 0.0668, 1000}, // Madone: 31m44 (29m40)
		{2184, 16380, 0.0571, 1000}, // Diablo (S): 36m24 (38m36)
		{1375, 8900, 0.0746, 1000},  // Ax-3: 22m55 (22m57)
		{2741, 15650, 0.0874, 1000}, // Ventoux: 45m41 (45m47)
		{2244, 10100, 0.1188, 1000}, // Zoncolan: 37m24 (39m04)
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
		{1113, 4800, 0.08125, 500},  // OLH: 18m33
		{884, 4800, 0.08125, 1000},  // OLH: 14m44 (17m12)
		{2579, 13910, 0.0790, 1000}, // Huez: 42m59 (49m46)
	}
	for _, tt := range tests {
		actual := CalcF(tt.t, tt.d, tt.gr)
		if !calc.Eqf(actual, tt.expected) {
			t.Errorf("CalcF(%.3f, %.3f, %.3f): got: %.3f, want: %.3f",
				tt.t, tt.d, tt.gr, actual, tt.expected)
		}
	}
}
