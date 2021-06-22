package context

import (
	"context"
	"testing"

	"github.com/libp2p/go-routing-language/middleware"
	"github.com/libp2p/go-routing-language/syntax"
)

var c = "bafkreiapctlo25cmwulbz4ylpwh442bh5zb4gykm5lry3u4iuvi4rs2s5y"
var f = syntax.Predicate{
	Tag: "find",
	Named: syntax.Pairs{
		syntax.Pair{
			Key: syntax.String{Value: "cid"},
			Value: syntax.List{
				Elements: syntax.Nodes{
					syntax.Predicate{
						Tag: "link",
						Positional: syntax.Nodes{
							syntax.String{Value: c},
						},
					},
				},
			},
		},
	},
}

// FindPassThrough simple bypass middleware. It returns the find expressions as-is
type mockTermination struct{}

func (m *mockTermination) Route(ctx *middleware.MiddlCtx, in syntax.Predicate) (syntax.Predicate, error) {
	return syntax.Predicate{Tag: "ok"}, nil
}

func TestBypassChain(t *testing.T) {
	rCtx := New(context.Background())
	rCtx.Use(&middleware.FindPassThrough{})
	rCtx.Use(&middleware.FindPassThrough{})

	p := rCtx.Route(f)
	if !syntax.IsEqual(f, p) {
		t.Errorf("Passthrough didn't work")
	}

	rCtx = New(context.Background())
	rCtx.Use(&middleware.FindPassThrough{})
	rCtx.Use(&mockTermination{})

	p = rCtx.Route(f)
	if !syntax.IsEqual(syntax.Predicate{Tag: "ok"}, p) {
		t.Errorf("Termination didn't work")
	}

}
