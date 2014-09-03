// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package basic

import (
	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/be"
)

func init() {
	faculty.Root.AddTerminal("Merge", Merge{})
}

// Merge
type Merge struct{}

func (Merge) Materialize() be.Reflex {
	xEndo, xExo := be.NewSynapse()
	yEndo, yExo := be.NewSynapse()
	zEndo, zExo := be.NewSynapse()
	go func() {
		h := &merge{
			connected: make(chan struct{}),
			x:         make(chan Image),
			y:         make(chan Image),
		}
		h.z = zEndo.Focus(be.DontCognize)
		close(h.connected)
		xEndo.Focus(h.CognizeX)
		yEndo.Focus(h.CognizeY)
		go h.xyz()
	}()
	return be.Reflex{
		"X": xExo, // write-only
		"Y": yExo, // write-only
		"_": zExo, // read-only
	}
}

type merge struct {
	connected chan struct{}
	x, y      chan Image
	z         *be.ReCognizer
}

func (h *merge) CognizeX(v interface{}) {
	h.x <- v.(Image)
}

func (h *merge) CognizeY(v interface{}) {
	h.y <- v.(Image)
}

func (h *merge) xyz() {
	<-h.connected
	var x = Make()
	var y = Make()
	for {
		select {
		case x = <-h.x:
			if x == nil {
				x = Make()
			}
			h.z.ReCognize(Make().Attach(x).Attach(y))
		case y = <-h.y:
			if y == nil {
				y = Make()
			}
			h.z.ReCognize(Make().Attach(x).Attach(y))
		}
	}
}
