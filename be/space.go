// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package be

import (
	// "fmt"
	"log"
	"strings"

	. "github.com/gocircuit/escher/image"
	"github.com/gocircuit/escher/see"
	"github.com/gocircuit/escher/understand"
)

// Space encloses the entire faculty tree of an Escher program, starting from the root.
type Space understand.Faculty

func (x Space) Faculty() understand.Faculty {
	return understand.Faculty(x)
}

// Lookup looks up the design for a root name; name should be a textual path.
func (x Space) Lookup(within understand.Faculty, name string) (d interface{}) {
	if strings.Index(name, ".") >= 0 {
		panic(7)
	}
	return within[""].(*understand.Circuit).Peer[name].Design
}

// Materialize creates the reflex, described by the absolute path walk.
func (x Space) Materialize(walk ...string) Reflex {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Problem materializing %s (%v)", strings.Join(walk,"."), r)
		}
	}()
	within_, term := x.Faculty().Walk(walk...)
	if gate, ok := term.(Gate); ok {
		return gate.Materialize()
	}
	within, cir := within_.(understand.Faculty), term.(*understand.Circuit)
	// println(cir.Print("	", "\t"))
	peers := make(map[string]Reflex)
	for _, peer := range cir.Peer {
		name, ok := peer.Name.(string)
		if !ok {
			log.Fatalf("circuit peers cannot have non-textual names, in circuit %s at peer %v", cir.Name, peer.Name)
		}
		if name == "" { // skip the super peer of this circuit
			continue
		}
		switch t := peer.Design.(type) {
		case see.RootPath:
			peers[name] = x.Materialize([]string(t)...)
		case see.Path: // e.g. “hello.who.is.there”
			peers[name] = x.materializePath(within, []string(t))
		case string, int, float64, complex128:
			peers[name] = NewNounReflex(t) // materialize builtin gates
		case Image:
			peers[name] = NewNounReflex(t.Copy()) // materialize images
		default:
			panicf("unknown design: %T/%v", t, t)
		}
	}
	// Connect/attach all reflex memories
	super := make(Reflex)
	for _, p := range cir.Peer {
		name := p.Name.(string)
		if name == "" {
			continue
		}
		for _, v := range p.Valve {
			m1 := peers[name][v.Name]
			if m1 == nil {
				continue
			}
			delete(peers[name], v.Name)
			if v.Matching.Of.Name == "" {
				if _, ok := super[v.Matching.Name]; ok {
					panic(6)
				}
				super[v.Matching.Name] = m1
			} else {
				qp, qv := v.Matching.Of.Name.(string), v.Matching.Name
				m2 := peers[qp][qv]
				delete(peers[qp], qv)
				go Merge(m1, m2)
			}
		}
	}
	// Check for unmatched valves
	for pname, p := range peers {
		for vname, _ := range p {
			panicf("%s.%s not matched", pname, vname)
		}
	}
	return super
}

func (x Space) materializePath(within understand.Faculty, parts []string) Reflex {
	unfold := x.Lookup(within, parts[0])
	switch t := unfold.(type) {
	case string, int, float64, complex128:
		return NewNounReflex(t) // materialize builtin gates
	case Image:
		return NewNounReflex(t.Copy())
	case see.Path:
		parts = append([]string(t), parts[1:]...)
		return Space(within).Materialize(parts...)
	case see.RootPath:
		parts = append([]string(t), parts[1:]...)
		return x.Materialize(parts...)
	case nil:
		return nil
	}
	panic("unknown design")
}

func (x Space) Interpret(understood *understand.Circuit) (fresh *understand.Circuit) {
	return x.Faculty().Interpret(understood)
}

func (x Space) Forget(name string) (forgotten interface{}) {
	switch t := x.Faculty().Forget(name).(type) {
	case understand.Faculty:
		return Space(t)
	default:
		return t
	}
	panic(0)
}

func (x Space) Walk(walk ...string) (parent, child interface{}) {
	parent, child = x.Faculty().Walk(walk...)
	if fac, ok := parent.(understand.Faculty); ok {
		parent = Space(fac)
	}
	if fac, ok := child.(understand.Faculty); ok {
		child = Space(fac)
	}
	return
}

func (x Space) Roam(walk ...string) (parent, child interface{}) {
	parent, child = x.Faculty().Roam(walk...)
	if fac, ok := parent.(understand.Faculty); ok {
		parent = Space(fac)
	}
	if fac, ok := child.(understand.Faculty); ok {
		child = Space(fac)
	}
	return
}