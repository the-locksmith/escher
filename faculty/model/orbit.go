// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package model

import (
	"container/list"
	// "fmt"
	"log"

	"github.com/gocircuit/escher/faculty"
	. "github.com/gocircuit/escher/circuit"
	"github.com/gocircuit/escher/be"
)

/*
	Orbit traverses the hierarchy of circuits induced by a given top-level/valveless circuit.

	Start = {
		Circuit Circuit
		Vector Vector
	}

	View = {
		Circuit Circuit // Current circuit in the exploration sequence
		Vector Vector
		Index int // Index of this circuit within exploration sequence, 0-based
		Depth int
		Dir string
		Series string // Loop
	}
*/
type Orbit struct{}

func CognizeView(*be.Eye, interface{}) {}

func CognizeStart(eye *be.Eye, dv interface{}) {
	var in = dv.(Circuit)
	var start = view{
		Circuit: in.CircuitAt("Circuit"),
		Vector: Vector(in.CircuitAt("Vector")),
		Index: 0,
		Depth: 0,
	}
	var v = start
	var memory list.List
	for {
		eye.Show("View", v.Circuitize()) // yield current view

		switch t := v.Circuit.At(v.Vector.Gate()).(type) { // next gate
		case Address: // Down
			if memory.Len() > 100 {
				log.Fatalf("memory overload")
				// memory.Remove(memory.Front())
			}
			memory.PushFront(v) // remember
			//
			_, lookup := faculty.Root.LookupAddress(t.String())
			v.Circuit = lookup.(Circuit) // transition to next circuit
			v.Vector = v.Circuit.Follow(v.Vector)
			v.Depth++

		case Super: // Up
			e := memory.Front() // backtrack
			if e == nil {
				log.Fatalf("short memory")
			}
			u := e.Value.(view)
			memory.Remove(e)
			//
			v.Circuit = u.Circuit
			v.Vector = v.Circuit.Follow(NewVector(u.Vector.Gate(), v.Vector.Valve()))
			v.Depth--

		default:
			panic("unknown gate meaning")
		}
		v.Index++
		//
		if Same(v.Circuit, start.Circuit) && Same(v.Vector, start.Vector) {
			eye.Show("View", v.Circuitize().Grow("Series", "Loop")) // yield current view
			break
		}
	}
}

type view struct {
	Circuit Circuit
	Vector Vector
	Index int
	Depth int
}

func (v *view) Dir() string {
	if _, ok := v.Circuit.At(v.Vector.Gate()).(Super); ok {
		return "Up"
	}
	return "Down"
}

func (v *view) Circuitize() Circuit {
	return New().
		Grow("Circuit", v.Circuit).
		Grow("Vector", Circuit(v.Vector)).
		Grow("Index", v.Index).
		Grow("Depth", v.Depth).
		Grow("Dir", v.Dir())
}