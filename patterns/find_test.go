package patterns

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-routing-language/parse"
	"github.com/libp2p/go-routing-language/syntax"
)

func TestFindCid(t *testing.T) {
	cid1, err := cid.Decode("QmQeJmz16RwLgbb8hq5trFYoPyZ7UjAjieqzEs3JEf6ggD")
	if err != nil {
		t.Fatal(err)
	}
	// f0 = find(link(CID))
	f0 := syntax.Predicate{
		Tag: "find",
		Positional: syntax.Nodes{
			(&Link{Cid: cid1}).Express(),
		},
	}
	pctx := parse.NewParseCtx()
	f1, err := ParseFindCid(pctx, f0)
	if err != nil {
		t.Fatalf("find not parsed (%v)", err)
	}
	f2 := f1.Express()
	if !syntax.IsEqual(f0, f2) {
		t.Errorf("find cid expressions are not the same")
	}
}
