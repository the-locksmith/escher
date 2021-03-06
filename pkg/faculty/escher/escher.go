// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package escher

import (
	"github.com/hoijui/escher/pkg/be"
	"github.com/hoijui/escher/pkg/faculty"
)

func init() {
	faculty.Register(be.NewMaterializer(&System{}), "e", "Materialize")
	faculty.Register(be.NewMaterializer(Index{}), "e", "Index")
	faculty.Register(be.NewMaterializer(Parse{}), "e", "Parse")
	faculty.Register(be.NewMaterializer(Breakpoint{}), "e", "Breakpoint")
	faculty.Register(be.NewMaterializer(&Help{}), "e", "help")
	faculty.Register(be.NewMaterializer(&Help{}), "e", "Help")
}
