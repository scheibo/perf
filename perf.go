// Package perf implements the 'Performance Equivalency Rating Formula' for
// scoring road cycling performances on climbs for ranking and comparison purposes.
package perf

import (
	"math"

	"github.com/scheibo/calc"
)

const mb = 8.0

const mrM = 67.0
const mrF = 53.0

const cdaM = 0.325 // 1.80m
const cdaF = 0.300 // 1.67m

// CalcM calculates the PERF score for a performance of duration t on a climb
// of distance d in metres, gradient gr (rise/run) and median elevation h in
// metres for a male rider.
func CalcM(t, d, gr, h float64) float64 {
	return score(power(t, d, gr, h, mrM, cdaM), calc.AltitudeAdjust(CpM(t), h))
}

// CalcF calculates the PERF score for a performance of duration t on a climb
// of distance d in metres, gradient gr (rise/run) and median elevation h in
// metres for a female rider.
func CalcF(t, d, gr, h float64) float64 {
	return score(power(t, d, gr, h, mrF, cdaF), calc.AltitudeAdjust(CpF(t), h))
}

// CpM returns the expected power maintable for a male rider during a world
// record level performance of duration t.
//
//		MALE (67 kg)
//		----------------
//		1 min: 11.5 W/kg
//		5min:  7.60 W/kg
//		60min: 6.40 W/kg
//		----------------
//		770 W/509 W/429 W
func CpM(t float64) float64 {
	return 422.58 + 23296.7801287949/t
}

// CpF returns the expected power maintable for a female rider during a world
// record level performance of duration t.
//
//		FEMALE (53 kg)
//		----------------
//		1 min: 9.29 W/kg
//		5min:  6.61 W/kg
//		60min: 5.69 W/kg
//		----------------
//		492 W/350 W/302 W
func CpF(t float64) float64 {
	return 298.29 + 13499.9080036799/t
}

// CalcTimeM calculates the duration of a performance with PERF score s on a climb
// of distance d in metres, gradient gr (rise/run) and median elevation h in
// metres for a male rider.
func CalcTimeM(s, d, gr, h float64) float64 {
	return time_(CalcPowerM(s, d, gr, h), d, gr, h, mrM, cdaM)
}

// CalcTimeF calculates the duration of a performance with PERF score s on a climb
// of distance d in metres, gradient gr (rise/run) and median elevation h in
// metres for a female rider.
func CalcTimeF(s, d, gr, h float64) float64 {
	return time_(CalcPowerF(s, d, gr, h), d, gr, h, mrF, cdaF)
}

// CalcPowerM calculates the power required for a performance with PERF score s on a
// climb of distance d in metres, gradient gr (rise/run) and median elevation h in
// metres for the normalized model male rider.
func CalcPowerM(s, d, gr, h float64) float64 {
	return rscore(s, d, gr, h, mrM, cdaM, CpM)
}

// CalcPowerF calculates the power required for a performance with PERF score s on a
// climb of distance d in metres, gradient gr (rise/run) and median elevation h in
// metres for the normalized model female rider.
func CalcPowerF(s, d, gr, h float64) float64 {
	return rscore(s, d, gr, h, mrF, cdaF, CpF)
}

func rscore(s, d, gr, h, mr, cda float64, cp func(float64) float64) float64 {
	// epsilon is some small value that determines when we will stop the search
	const epsilon = 1e-6
	// max is the maxmium number of iterations of the search
	const max = 100

	p := 0.0
	tl, tm, th := 0.0, 3600.0, 7200.0
	for j := 0; j < max; j++ {
		p = power(tm, d, gr, h, mr, cda)
		s1 := score(p, calc.AltitudeAdjust(cp(tm), h))

		if calc.Eqf(s1, s, epsilon) {
			break
		}

		if s1 > s {
			tl = tm
		} else {
			th = tm
		}

		tm = (th + tl) / 2.0
	}

	return p
}

func power(t, d, gr, h, mr, cda float64) float64 {
	vg := d / t
	return calc.Psimp(calc.Rho(h, calc.G), cda, calc.Crr, vg, vg, gr, mr+mb, calc.G, calc.Ec, calc.Fw)
}

func time_(p, d, gr, h, mr, cda float64) float64 {
	return calc.Time(p, d, calc.Rho(h, calc.G), cda, calc.Crr, 0, 0, 0, gr, mr+mb, calc.G, calc.Ec, calc.Fw)
}

func score(p, wr float64) float64 {
	return 1000 * math.Pow(p/wr, 1.8)
}
