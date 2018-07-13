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
	return score(t, twr(d, gr, h, mrM, cdaM, CpM))
}

// CalcF calculates the PERF score for a performance of duration t on a climb
// of distance d in metres, gradient gr (rise/run) and median elevation h in
// metres for a female rider.
func CalcF(t, d, gr, h float64) float64 {
	return score(t, twr(d, gr, h, mrF, cdaF, CpF))
}

// CpM returns the expected power maintable for a male rider during a world
// record level performance of duration t.
//
// MALE (67 kg)
// 1 min: 11.5 W/kg
// 5min:  7.60 W/kg
// 60min: 6.40 W/kg
// 770 W/509 W/429 W
func CpM(t float64) float64 {
	return 422.58 + 23296.7801287949/t
}

// CpF returns the expected power maintable for a female rider during a world
// record level performance of duration t.
//
// FEMALE (53 kg)
// 1 min: 9.29 W/kg
// 5min:  6.61 W/kg
// 60min: 5.69 W/kg
// 492 W/350 W/302 W
func CpF(t float64) float64 {
	return 298.29 + 13499.9080036799/t
}

func twr(d, gr, h, mr, cda float64, cp func(float64) float64) float64 {
	// epsilon is some small value that determines when we will stop the search
	const epsilon = 1e-6
	// max is the maxmium number of iterations of the search
	const max = 100

	mt := mr + mb
	tl, tm, th := 0.0, 3600.0, 7200.0
	for j := 0; j < max; j++ {
		vg := d / tm
		p1 := calc.Psimp(calc.Rho(h, calc.G), cda, calc.Crr, vg, vg, gr, mt, calc.G, calc.Ec, calc.Fw)
		p2 := calc.AltitudeAdjust(cp(tm), h)

		if calc.Eqf(p1, p2, epsilon) {
			break
		}

		if p1 > p2 {
			tl = tm
		} else {
			th = tm
		}

		tm = (th + tl) / 2.0
	}

	return tm
}

func score(t, wr float64) float64 {
	return 1000 * math.Pow(wr/t, 2)
}
