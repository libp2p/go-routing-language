package syntax

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPretty(t *testing.T) {
	n := Dict{
		Pairs: Pairs{
			{String{"bar"}, String{"baz"}},
			{String{"bar2"}, NewInt64(567)},
		},
	}
	var w bytes.Buffer
	n.WritePretty(&w)
	fmt.Println(w.String())
}
