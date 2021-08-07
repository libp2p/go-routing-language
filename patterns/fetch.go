package patterns

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// Fetch represents a FETCH pattern.
type Fetch interface {
	Pattern
	IsFetch()
}

// FetchCid is the Go representation of the routing language pattern
// 	fetch(cid=link(CID:STRING), providers=[PROVIDER])
type FetchCid struct {
	Cid       cid.Cid
	Providers Providers
}

// IsFetch informs the Go type system that FetchCid is a fetch pattern.
func (f *FetchCid) IsFetch() {}

// Express returns the syntactic representation of the fetch cid pattern.
func (f *FetchCid) Express() syntax.Node {
	return syntax.Predicate{
		Tag: "fetch",
		Named: syntax.Pairs{
			{Key: syntax.String{Value: "cid"}, Value: (&Link{Cid: f.Cid}).Express()},
			{Key: syntax.String{Value: "providers"}, Value: f.Providers.Express()},
		},
	}
}

// FetchCidParser is a parser for fetch cid patterns.
type FetchCidParser struct{}

// Parse parses a fetch cid pattern.
func (FetchCidParser) Parse(ctx *parse.ParseCtx, src syntax.Node) (interface{}, error) {
	return ParseFetchCid(ctx, src)
}

// MatchAllFetchCid parses all fetch cid patterns found in src.
func MatchAllFetchCid(ctx *parse.ParseCtx, match parse.Parser, src syntax.Node) []*FetchCid {
	m := FetchCidParser{}
	found := parse.MatchAll(ctx, m, src)
	r := make([]*FetchCid, len(found))
	for i, e := range found {
		r[i] = e.(*FetchCid)
	}
	return r
}

// ParseFetchCid parses formulas of the form `fetch(cid=link(CID:STRING), providers=[PROVIDER])`.
func ParseFetchCid(ctx *parse.ParseCtx, src syntax.Node) (*FetchCid, error) {
	// parse fetch(...)
	p0, ok := src.(syntax.Predicate)
	if !ok {
		return nil, fmt.Errorf("not a predicate")
	}
	if p0.Tag != "fetch" {
		return nil, fmt.Errorf("tag is not find")
	}
	// parse cid=link(...)
	cid0 := p0.GetNamed("cid")
	if cid0 == nil {
		return nil, fmt.Errorf("cid argument missing")
	}
	link, err := ParseLink(ctx, cid0)
	if err != nil {
		return nil, fmt.Errorf("parsing cid (%v)", err)
	}
	// parse providers=[...]
	prov0 := p0.GetNamed("providers")
	if prov0 == nil {
		return nil, fmt.Errorf("providers argument missing")
	}
	prov, err := ParseProviders(ctx, prov0)
	if err != nil {
		return nil, fmt.Errorf("parsing providers (%v)", err)
	}
	return &FetchCid{Cid: link.Cid, Providers: prov}, nil
}
