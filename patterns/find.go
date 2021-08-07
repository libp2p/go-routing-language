package patterns

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// FindCid is the Go representation of the `find(link(CID:STRING))` pattern from the Routing Language Spec.
type FindCid struct {
	Cid cid.Cid
}

// Express returns the syntactic representation of the find cid pattern.
func (f *FindCid) Express() syntax.Node {
	return syntax.Predicate{
		Tag: "find",
		Positional: syntax.Nodes{
			(&Link{Cid: f.Cid}).Express(),
		},
	}
}

// FindCidParser is a parser for `find(link(CID:STRING))` patterns.
type FindCidParser struct{}

// Parse parses a fetch cid pattern.
func (FindCidParser) Parse(ctx *parse.ParseCtx, src syntax.Node) (interface{}, error) {
	return ParseFindCid(ctx, src)
}

// MatchAllFindCid parses all find cid patterns found in src.
func MatchAllFindCid(ctx *parse.ParseCtx, match parse.Parser, src syntax.Node) []*FindCid {
	m := FindCidParser{}
	found := parse.MatchAll(ctx, m, src)
	r := make([]*FindCid, len(found))
	for i, e := range found {
		r[i] = e.(*FindCid)
	}
	return r
}

// ParseFindCid parses formulas of the form `find(link(CID:STRING))`.
func ParseFindCid(ctx *parse.ParseCtx, src syntax.Node) (*FindCid, error) {
	p0, ok := src.(syntax.Predicate)
	if !ok {
		return nil, fmt.Errorf("not a predicate")
	}
	if p0.Tag != "find" {
		return nil, fmt.Errorf("tag is not find")
	}
	if len(p0.Positional) != 1 {
		return nil, fmt.Errorf("expecting one argument")
	}
	link, err := ParseLink(ctx, p0.Positional[0])
	if err != nil {
		return nil, fmt.Errorf("parsing link (%v)", err)
	}
	return &FindCid{Cid: link.Cid}, err
}

// FindPath is the Go representation of the `find(path(PATH:STRING))` pattern from the Routing Language Spec.
type FindPath struct {
	Path string
}

// Express returns the syntactic representation of the find path pattern.
func (f *FindPath) Express() syntax.Node {
	return syntax.Predicate{
		Tag: "find",
		Positional: syntax.Nodes{
			syntax.Predicate{
				Tag: "path",
				Positional: syntax.Nodes{
					syntax.String{Value: f.Path},
				},
			},
		},
	}
}
