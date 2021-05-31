package parse

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-routing-language/syntax"
)

// ParseLink parses formulas of the form link(CID:STRING) to a libp2p cid.
func ParseLink(ctx *ParseCtx, src syntax.Node) (cid.Cid, error) {
	p, ok := src.(syntax.Predicate)
	if !ok {
		return cid.Cid{}, fmt.Errorf("not a predicate")
	}
	if p.Tag != "link" {
		return cid.Cid{}, fmt.Errorf("tag is not link")
	}
	if len(p.Positional) != 1 {
		return cid.Cid{}, fmt.Errorf("expecting one argument")
	}
	s, ok := p.Positional[0].(syntax.String)
	if !ok {
		return cid.Cid{}, fmt.Errorf("expecting a string argument")
	}
	c, err := cid.Decode(s.Value)
	return c, err
}
