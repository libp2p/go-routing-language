package patterns

import (
	"github.com/libp2p/go-routing-language/syntax"
)

// Pattern represents Go implementations of patterns from the Routing Language Spec:
// https://docs.google.com/document/d/1bGQ3-1u1XgfcXrb0FqbPLtelaOFpuRpsKTpyrgTtfDs/edit
type Pattern interface {
	// Express returns the syntactic representation of a pattern.
	Express() syntax.Node
}
