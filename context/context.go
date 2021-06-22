package context

import (
	"context"

	"github.com/libp2p/go-routing-language/middleware"
	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

// Context defines the chain of middlewares in the client context.
type RoutingCtx interface {
	// Use adds a new middleware level to the chain.
	Use(m middleware.Middleware)
	// Route triggers the processing of the chain of middlewares
	Route(syntax.Predicate) syntax.Predicate
}

type routingCtx struct {
	ctx context.Context
	// NOTE: Let's consider a sequential chain for now.
	// Initially I was implementing a tree of middlewares where
	// middlewares could be triggered in parallel. It added additional
	// complexity in this early stage so I disregarded it for now.
	chain []middleware.Middleware
}

func New(ctx context.Context) RoutingCtx {
	return &routingCtx{ctx: ctx}
}

// Use adds a new middleware level to the chain.
func (c *routingCtx) Use(m middleware.Middleware) {
	c.chain = append(c.chain, m)
}

// Route process routing expression as input.
func (c *routingCtx) Route(in syntax.Predicate) syntax.Predicate {
	for _, m := range c.chain {
		next, err := m.Route(&middleware.MiddlCtx{}, in)
		// If there is an error jump to the next middleware in the chain
		// We assume no dependencies between middlewares at this point.
		if err != nil {
			continue
		}
		// If OK predicate it means the full expression was processed.
		// We don't need to move to the rest of the middlewares in the chain.
		if _, err := parse.ParseOk(&parse.ParseCtx{}, next); err == nil {
			return next
		}

		// The result of the current output becomes the input of the next one
		in = next
	}
	return in
}
