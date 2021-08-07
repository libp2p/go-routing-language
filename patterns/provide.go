package patterns

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// FindCid is the Go representation of the `provide(cid=CID:CID, fetch=FETCH:FETCH)`
// pattern from the Routing Language Spec.
type ProvideCid struct {
	Cid   cid.Cid
	Fetch Fetch
}

// Express returns the syntactic representation of the provide expression.
func (p *ProvideCid) Express() syntax.Node {
	return syntax.Predicate{
		Tag: "provide",
		Named: syntax.Pairs{{
			Key:   syntax.String{Value: "fetch"},
			Value: p.Fetch.Express(),
		}},
	}
}

type ProvideCidParser struct{}

func (ProvideCidParser) Parse(ctx *parse.ParseCtx, src syntax.Node) (interface{}, error) {
	return ParseProvideCid(ctx, src)
}

// ParseProvideCid parses formulas of the form `provide(cid=CID:CID, fetch=FETCH:FETCH)`.
func ParseProvideCid(ctx *parse.ParseCtx, src syntax.Node) (*ProvideCid, error) {
	p0, ok := src.(syntax.Predicate)
	if !ok {
		return nil, fmt.Errorf("not a predicate")
	}
	if p0.Tag != "provide" {
		return nil, fmt.Errorf("tag is not provide")
	}
	// parse cid
	cid0 := p0.Named.ValueOf(syntax.String{Value: "cid"})
	if cid0 == nil {
		return nil, fmt.Errorf("cid argument missing")
	}
	cid1, err := ParseLink(ctx, cid0)
	if err != nil {
		return nil, fmt.Errorf("parsing cid (%v)", err)
	}
	// parse fetch
	fetch0 := p0.Named.ValueOf(syntax.String{Value: "fetch"})
	if fetch0 == nil {
		return nil, fmt.Errorf("fetch argument missing")
	}
	fetch1, err := ParseFetchCid(ctx, fetch0)
	if err != nil {
		return nil, fmt.Errorf("parsing fetch cid (%v)", err)
	}
	return &ProvideCid{Cid: cid1.Cid, Fetch: fetch1}, nil
}
