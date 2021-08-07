package patterns

import (
	"fmt"

	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// Provider represents the pattern PROVIDER = MULTIADDRESS | PEER
type Provider interface {
	Pattern
	IsProvider()
}

// ParseProvider parses a provider pattern from a syntactic representation.
func ParseProvider(ctx *parse.ParseCtx, src syntax.Node) (Provider, error) {
	multi, err := ParseMultiaddr(ctx, src)
	if err == nil {
		return multi, nil
	}
	peer, err := ParsePeer(ctx, src)
	if err == nil {
		return peer, nil
	}
	return nil, fmt.Errorf("not a recognizable provider")
}

// Providers is a list of providers.
type Providers []Provider

// Express returns the syntactic representation of the list of providers.
func (p Providers) Express() syntax.Node {
	n := make(syntax.Nodes, len(p))
	for i, u := range p {
		n[i] = u.Express()
	}
	return syntax.List{Elements: n}
}

// ParseProviders parses a list of providers, ignoring unrecognizable provider patterns.
func ParseProviders(ctx *parse.ParseCtx, src syntax.Node) (Providers, error) {
	list, ok := src.(syntax.List)
	if !ok {
		return nil, fmt.Errorf("not a list")
	}
	prov := Providers{}
	for _, e := range list.Elements {
		if p, err := ParseProvider(ctx, e); err == nil {
			prov = append(prov, p)
		}
	}
	return prov, nil
}
