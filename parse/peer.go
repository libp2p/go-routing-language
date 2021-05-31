package parse

import (
	"fmt"

	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-routing-language/syntax"
)

// ParsePeer parses formulas of the form peer(MULTIFORMAT:STRING) to a libp2p peer id.
func ParsePeer(ctx *ParseCtx, src syntax.Node) (peer.ID, error) {
	p, ok := src.(syntax.Predicate)
	if !ok {
		return peer.ID(""), fmt.Errorf("not a predicate")
	}
	if p.Tag != "peer" {
		return peer.ID(""), fmt.Errorf("tag is not peer")
	}
	if len(p.Positional) != 1 {
		return peer.ID(""), fmt.Errorf("expecting one argument")
	}
	s, ok := p.Positional[0].(syntax.String)
	if !ok {
		return peer.ID(""), fmt.Errorf("expecting a string argument")
	}
	c, err := peer.Decode(s.Value)
	return c, err
}
