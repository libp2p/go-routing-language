package parse

import (
	"github.com/libp2p/go-routing-language/syntax"
)

// Parser represents an algorithm that recognizes a set of syntactic expressions
// and returns a Go representation of the parsed information.
// This interface enables the implementation of generic compositions of parsing rules,
// also known as "parser combinators".
type Parser interface {
	Parse(ctx *ParseCtx, src syntax.Node) (interface{}, error)
}

type ParseCtx struct {
	keys map[interface{}]interface{}
}

func (x *ParseCtx) Set(key interface{}, value interface{}) interface{} {
	v := x.keys[key]
	x.keys[key] = value
	return v
}

func (x *ParseCtx) Get(key interface{}) (interface{}, bool) {
	v, ok := x.keys[key]
	return v, ok
}
