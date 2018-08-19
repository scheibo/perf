// perf provides a CLI for calculating PERF score of an arbitrary performance.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/scheibo/perf"
)

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	flag.PrintDefaults()
	os.Exit(1)
}

func verify(s string, x float64) {
	if x < 0 {
		exit(fmt.Errorf("%s must be non negative but was %f", s, x))
	}
}

func main() {
	var e, gr, t, d, h, score float64
	var x string
	var dur time.Duration

	flag.Float64Var(&e, "e", 0, "total elevation gained in m")
	flag.Float64Var(&gr, "gr", 0, "average grade")
	flag.Float64Var(&h, "h", 0, "median elevation")
	flag.StringVar(&x, "x", "M", "sex of the athlete")

	flag.Float64Var(&d, "d", -1, "distance travelled in m")
	flag.DurationVar(&dur, "t", -1, "duration in minutes and seconds ('12m34s')")

	flag.Parse()

	verify("h", h)

	// error correct in case grade was passed in as a %
	if gr > 1 || gr < -1 {
		gr = gr / 100
	}

	if d <= 0 {
		exit(fmt.Errorf("d must be positive but was %f", d))
	}

	if e > 0 {
		// if both are specified, make sure they agree
		if gr > 0 && ((d*gr != e) || (e/d != gr)) {
			exit(fmt.Errorf("specified both e=%f and gr=%f but they do not agree", e, gr))
		}
		gr = e / d
	}

	if t != -1 {
		verify("t", float64(dur))
		t = float64(dur / time.Second)
		if x == "M" {
			score = perf.CalcM(t, d, gr, h)
		} else {
			score = perf.CalcF(t, d, gr, h)
		}

		fmt.Printf("%s (%.2f km @ %.2f%%) = %.2f\n", fmtDuration(dur), d/1000, gr*100, score)
	} else {
		exit(fmt.Errorf("t must be specified"))
	}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%d:%02d", m, s)
}
