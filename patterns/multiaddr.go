package patterns

import (
	"fmt"

	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
	ma "github.com/multiformats/go-multiaddr"
)

// Multiaddr represents the pattern `multiaddr(MULTIADDR:STRING)` or `multiaddr(MULTIADDR:BYTES)`
type Multiaddr struct {
	Multiaddr ma.Multiaddr
}

func (p *Multiaddr) IsProvider() {}

func (p *Multiaddr) Express() syntax.Node {
	return syntax.Predicate{
		Tag: "multiaddr",
		Positional: syntax.Nodes{
			syntax.String{p.Multiaddr.String()},
		},
	}
}

// MultiaddrParser is a parser for the peer pattern.
type MultiaddrParser struct{}

func (MultiaddrParser) Parse(ctx *parse.ParseCtx, src syntax.Node) (interface{}, error) {
	return ParseMultiaddr(ctx, src)
}

// ParseMultiaddr parses formulas of the form multiaddr(MULTIADDR:STRING) to a libp2p multiaddr.
func ParseMultiaddr(ctx *parse.ParseCtx, src syntax.Node) (*Multiaddr, error) {
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
	if err != nil {
		return nil, err
	}
	return &Multiaddr{Multiaddr: c}, nil
}
