package middleware

import (
	"testing"

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

func TestSimpleBypass(t *testing.T) {
	m := FindPassThrough{}
	p, err := m.Route(&MiddlCtx{}, f)
	if err != nil {
		t.Fatal(err)
	}
	if !syntax.IsEqual(f, p) {
		t.Errorf("Passthrough didn't work")
	}
}
