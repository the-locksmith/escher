// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package cmplx

import (
	"math/cmplx"

	"github.com/hoijui/escher/pkg/be"
	cir "github.com/hoijui/escher/pkg/circuit"
)

// Polar
type Polar struct{}

func (Polar) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	return nil
}

func (Polar) CognizeComplex(eye *be.Eye, v interface{}) {
	r, theta := cmplx.Polar(v.(complex128))
	eye.Show("Polar", cir.New().Grow("R", r).Grow("Theta", theta))
}

func (Polar) CognizePolar(eye *be.Eye, v interface{}) {
	x := v.(cir.Circuit)
	eye.Show("Complex", cmplx.Rect(x.FloatAt("R"), x.FloatAt("Theta")))
}
