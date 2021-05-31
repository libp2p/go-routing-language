package parse

import (
	"fmt"

	"github.com/libp2p/go-routing-language/syntax"
	ma "github.com/multiformats/go-multiaddr"
)

// ParseMultiaddr parses formulas of the form multiaddr(MULTIADDR:STRING) to a libp2p multiaddr.
func ParseMultiaddr(ctx *ParseCtx, src syntax.Node) (ma.Multiaddr, error) {
	p, ok := src.(syntax.Predicate)
	if !ok {
		return nil, fmt.Errorf("not a predicate")
	}
	if p.Tag != "multiaddr" {
		return nil, fmt.Errorf("tag is not multiaddr")
	}
	if len(p.Positional) != 1 {
		return nil, fmt.Errorf("expecting one argument")
	}
	s, ok := p.Positional[0].(syntax.String)
	if !ok {
		return nil, fmt.Errorf("expecting a string argument")
	}
	c, err := ma.NewMultiaddr(s.Value)
	return c, err
}
