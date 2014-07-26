// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package see

import (
	"github.com/gocircuit/escher/star"
)

func SeeStar(src *Src) (x *star.Star) {
	defer func() {
		if r := recover(); r != nil {
			x = nil
		}
	}()
	x = star.Make()
	t := src.Copy()
	t.Match("{")
	Space(t)
	for {
		q := t.Copy()
		Space(q)
		name, peer := SeePeer(q)
		if peer == nil {
			break
		}
		Space(q)
		q.TryMatch(",")
		Space(q)
		t.Become(q)
		x.Grow(name, "", peer)
	}
	Space(t)
	t.Match("}")
	src.Become(t)
	return x
}