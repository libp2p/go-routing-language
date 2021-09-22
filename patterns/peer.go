package patterns

import (
	"fmt"

	peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// Peer represents the pattern `peer(MULTIFORMAT:STRING)` or `peer(MULTIFORMAT:BYTES)`
type Peer struct {
	ID peer.ID
}

func (p *Peer) IsProvider() {}

func (p *Peer) Express() syntax.Node {
	return syntax.Predicate{
		Tag: "peer",
		Positional: syntax.Nodes{
			syntax.String{Value: p.ID.String()},
		},
	}
}

// PeerParser is a parser for the peer pattern.
type PeerParser struct{}

func (PeerParser) Parse(ctx *parse.ParseCtx, src syntax.Node) (interface{}, error) {
	return ParsePeer(ctx, src)
}

// ParsePeer parses formulas of the form `peer(MULTIFORMAT:STRING)` to a libp2p peer id.
func ParsePeer(ctx *parse.ParseCtx, src syntax.Node) (*Peer, error) {
	p, ok := src.(syntax.Predicate)
	if !ok {
		return nil, fmt.Errorf("not a predicate")
	}
	if p.Tag != "peer" {
		return nil, fmt.Errorf("tag is not peer")
	}
	if len(p.Positional) != 1 {
		return nil, fmt.Errorf("expecting one argument")
	}
	s, ok := p.Positional[0].(syntax.String)
	if !ok {
		return nil, fmt.Errorf("expecting a string argument")
	}
	c, err := peer.Decode(s.Value)
	if err != nil {
		return nil, err
	}
	return &Peer{ID: c}, nil
}
