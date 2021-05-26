package syntax

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"
)

func TestPretty(t *testing.T) {
	n := Dict{
		Pairs: Pairs{
			{String{"bar1"}, String{"baz"}},
			{Int{big.NewInt(567)}, String{"baz"}},
			{String{"bar2"}, Int{big.NewInt(567)}},
			{String{"bar3"}, Bytes{[]byte("asdf")}},
			{Bytes{[]byte("asdf")}, Int{big.NewInt(567)}},
			{String{"bar4"}, Dict{
				Pairs: Pairs{
					{Bool{true}, Int{big.NewInt(567)}},
				},
			}},
			{String{"predicate"}, Predicate{
				Tag: "tag",
				Positional: Nodes{
					Bool{true},
					Int{big.NewInt(567)},
				},
				Named: Pairs{
					{Bool{true}, Int{big.NewInt(567)}},
				},
			}},
		},
	}
	var w bytes.Buffer
	n.WritePretty(&w)
	fmt.Println(w.String())
}
