package parse

import (
	"fmt"

	"github.com/libp2p/go-routing-language/syntax"
)

// ParseOk parses formulas of the form ok(NODES:EXPR). If the expression is OK
// it returns what is inside Nodes.
func ParseOk(ctx *ParseCtx, src syntax.Node) (syntax.Nodes, error) {
	p, ok := src.(syntax.Predicate)
	if !ok {
		return nil, fmt.Errorf("not a predicate")
	}
	if p.Tag != "ok" {
		return nil, fmt.Errorf("tag is not link")
	}
	if len(p.Positional) == 0 {
		return nil, fmt.Errorf("empty OK predicate")
	}
	return p.Positional, nil
}
