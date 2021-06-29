package parse

import (
	"github.com/libp2p/go-routing-language/syntax"
)

// MatchAll greedily finds all disjoint matches of a given pattern in a formula and returns them in the form of a list.
func MatchAll(ctx *ParseCtx, match Parser, src syntax.Node) []interface{} {
	m, err := match.Parse(ctx, src)
	if err == nil {
		return []interface{}{m}
	}
	r := []interface{}{}
	switch t := src.(type) {
	case syntax.Dict:
		for _, p := range t.Pairs {
			r = append(r, MatchAll(ctx, match, p.Value)...)
		}
	case syntax.List:
		for _, e := range t.Elements {
			r = append(r, MatchAll(ctx, match, e)...)
		}
	case syntax.Predicate:
		for _, a := range t.Positional {
			r = append(r, MatchAll(ctx, match, a)...)
		}
		for _, p := range t.Named {
			r = append(r, MatchAll(ctx, match, p.Value)...)
		}
	}
	return r
}
