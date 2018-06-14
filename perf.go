package perf

import (
	_ "github.com/scheibo/calc" // nolint
)

const mb = 8.0

const mrM = 67.0
const mrF = 53.0

const cdaM = 0.325
const cdaF = 0.325 // TODO scale!

// CalcM calculates the PERF score for a performance on a climb of distance d
// in metres and gradient gr (rise/run) for a male rider.
func CalcM(d, gr float64) float64 {
	_ = power(d, gr, mrM, cdaM)
	_ = cpM(0)
	return 0
}

// CalcF calculates the PERF score for a performance on a climb of distance d
// in metres and gradient gr (rise/run) for a female rider.
func CalcF(d, gr float64) float64 {
	_ = power(d, gr, mrF, cdaF)
	return 0
}

// TODO need the formula for time, not power!
func power(d, gr, mr, cda float64) float64 {
	_, _, _, _, _ = d, gr, mr, mb, cda
	_ = cpF(0)
	// NOTE: if vgi = vgf, Pke = 0, so we can set whatever we want for ti and tf provided tf - ti != 0
	//return calc.Ptot(calc.Rho0, cda, calc.Crr, va, vg, gr, mr + mb, calc.R, vg, vg, 0, 1, calc.G, calc.Ec, calc.Fw, calc.I)
	return 0
}

// MALE (67 kg)
// 1 min: 11.5 W/kg
// 5min:  7.60 W/kg
// 60min: 6.40 W/kg
// 770 W/509 W/429 W
func cpM(t float64) float64 {
	return 422.58 + 23296.7801287949/t
}

// FEMALE (53 kg)
// 1 min: 9.29 W/kg
// 5min:  6.61 W/kg
// 60min: 5.69 W/kg
// 492 W/350 W/302 W
func cpF(t float64) float64 {
	return 298.29 + 13499.9080036799/t
}
