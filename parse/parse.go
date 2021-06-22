package parse

import (
	"github.com/libp2p/go-routing-language/syntax"
)

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

// TODO: Missing parser for routing language spec. We know have individual parser, we need
// to build a parser combinator
