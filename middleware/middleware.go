package middleware

import (
	"context"

	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// Middleware implements and runs different routing protocols in the
// client's context
type Middleware interface {
	// Route receives a routing expression and return a processed expression.
	Route(*MiddlCtx, syntax.Predicate) (syntax.Predicate, error)
}

type MiddlCtx struct {
	ctx  context.Context
	keys map[interface{}]interface{}
}

// FindPassThrough simple bypass middleware. It returns the find expressions as-is
type FindPassThrough struct{}

func (m *FindPassThrough) Route(ctx *MiddlCtx, in syntax.Predicate) (syntax.Predicate, error) {
	return parse.ParseFind(&parse.ParseCtx{}, in)
}
