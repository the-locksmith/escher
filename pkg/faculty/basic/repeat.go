// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"github.com/hoijui/escher/pkg/be"
	cir "github.com/hoijui/escher/pkg/circuit"
)

type Repeat struct{}

func (Repeat) Spark(*be.Eye, cir.Circuit, ...interface{}) cir.Value {
	return nil
}

func (Repeat) CognizeValue(eye *be.Eye, value interface{}) {
	for {
		eye.Show(cir.DefaultValve, value)
	}
}

func (Repeat) Cognize(eye *be.Eye, value interface{}) {}
