package perf

import (
	"testing"
	"time"

	"github.com/scheibo/calc"
)

func TestCalcM(t *testing.T) {
	tests := []struct {
		dur                string
		d, gr, h, expected float64
	}{
		{"16m09s", 4800, 0.08125, (108 + 498) / 2, 500},   // OLH 13m26s
		{"12m49s", 4800, 0.08125, (108 + 498) / 2, 1000},  // OLH 13m26s
		{"39m04s", 13910, 0.0790, (733 + 1832) / 2, 1000}, // Huez 39m01s (36m50s)
		{"27m06s", 9880, 0.0799, (268 + 1059) / 2, 1000},  // Gibraltr 27m12s
		{"31m54s", 13100, 0.0668, (41 + 916) / 2, 1000},   // Madone 35m36 (29m40s)
		{"36m40s", 16380, 0.0571, (223 + 1159) / 2, 1000}, // Diablo (S) 38m36s
		{"23m29s", 8900, 0.0746, (727 + 1353) / 2, 1000},  // Ax-3 23m14s (22m57s)
		{"47m22s", 15650, 0.0874, (528 + 1870) / 2, 1000}, // Ventoux 48m35s (45m47s)
		{"38m54s", 10100, 0.1188, (546 + 1674) / 2, 1000}, // Zoncolan 39m04s
	}
	for _, tt := range tests {
		actual := CalcM(duration(t, tt.dur), tt.d, tt.gr, tt.h)
		if !calc.Eqf(actual, tt.expected) {
			t.Errorf("CalcM(%s, %.3f, %.3f, %.3f): got: %.3f, want: %.3f",
				tt.dur, tt.d, tt.gr, tt.h, actual, tt.expected)
		}
	}
}

func TestCalcF(t *testing.T) {
	tests := []struct {
		dur                string
		d, gr, h, expected float64
	}{
		{"18m39s", 4800, 0.08125, (108 + 498) / 2, 500},   // OLH 17m12s
		{"14m48s", 4800, 0.08125, (108 + 498) / 2, 1000},  // OLH 17m12s
		{"44m44s", 13910, 0.0790, (733 + 1832) / 2, 1000}, // Huez 49m46s
	}
	for _, tt := range tests {
		actual := CalcF(duration(t, tt.dur), tt.d, tt.gr, tt.h)
		if !calc.Eqf(actual, tt.expected) {
			t.Errorf("CalcF(%s, %.3f, %.3f, %.3f): got: %.3f, want: %.3f",
				tt.dur, tt.d, tt.gr, tt.h, actual, tt.expected)
		}
	}
}

func duration(t *testing.T, s string) float64 {
	dur, err := time.ParseDuration(s)
	if err != nil {
		t.Errorf("Failed to parse duration %s", s)
	}
	return float64(dur / time.Second)
}
