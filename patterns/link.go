package patterns

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// Link represents the pattern `link(CID:STRING)`
type Link struct {
	Cid cid.Cid
}

func (l *Link) Express() syntax.Node {
	return syntax.Predicate{
		Tag: "link",
		Positional: syntax.Nodes{
			syntax.String{Value: l.Cid.String()},
		},
	}
}

// LinkParser is a Parser for the link pattern.
type LinkParser struct{}

func (LinkParser) Parse(ctx *parse.ParseCtx, src syntax.Node) (interface{}, error) {
	return ParseLink(ctx, src)
}

// ParseLink parses formulas of the form `link(CID:STRING)`.
func ParseLink(ctx *parse.ParseCtx, src syntax.Node) (*Link, error) {
	p, ok := src.(syntax.Predicate)
	if !ok {
		return nil, fmt.Errorf("not a predicate")
	}
	if p.Tag != "link" {
		return nil, fmt.Errorf("tag is not link")
	}
	if len(p.Positional) != 1 {
		return nil, fmt.Errorf("expecting one argument")
	}
	s, ok := p.Positional[0].(syntax.String)
	if !ok {
		return nil, fmt.Errorf("expecting a string argument")
	}
	c, err := cid.Decode(s.Value)
	return &Link{Cid: c}, err
}
