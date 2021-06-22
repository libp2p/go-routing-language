package parse

import (
	"fmt"

	"github.com/libp2p/go-routing-language/syntax"
)

// ParseFind parses formulas to check if it is of the form find
// and transforms it into a predicate to be processed by the client's
// context
func ParseFind(ctx *ParseCtx, src syntax.Node) (syntax.Predicate, error) {
	p, ok := src.(syntax.Predicate)
	if !ok {
		return syntax.Predicate{}, fmt.Errorf("not a predicate")
	}
	if p.Tag != "find" {
		return syntax.Predicate{}, fmt.Errorf("tag is not link")
	}
	// TODO: We could check here if all the required arguments for
	// find are in place. For now, let's make middlewares responsible for
	// this.
	return p, nil
}
