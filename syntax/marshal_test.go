package syntax

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"
)

func TestMarshalJSONAll(t *testing.T) {
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

	// Encode
	b, err := MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	println(string(b))
	// Decode
	out, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}

	if !IsEqual(n, out) {
		fmt.Println(IsEqual(n, out))
		fmt.Println("== IN ==")
		var w bytes.Buffer
		n.WritePretty(&w)
		fmt.Println(w.String())

		fmt.Println("== OUT ==")
		w.Reset()
		out.WritePretty(&w)
		fmt.Println(w.String())
		t.Fatal("error unmarshalling Dict")
	}
}

func TestMarshalJSONString(t *testing.T) {
	n := String{"testing!"}
	b, err := MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	o, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}

	if !IsEqual(n, o) {
		t.Fatal("error unmarshalling string", n, o)
	}
}

func TestMarshalJSONBool(t *testing.T) {
	n := Bool{true}
	b, err := MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	o, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, o) {
		t.Fatal("error unmarshalling string", n, o)
	}

	n = Bool{false}
	b, err = MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	o, err = UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, o) {
		t.Fatal("error unmarshalling string", n, o)
	}
}

func TestMarshalJSONBytes(t *testing.T) {
	n := Bytes{[]byte("testing!")}
	b, err := MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	o, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, o) {
		t.Fatal("error unmarshalling blob", n, o)
	}
}

func TestMarshalJSONNumber(t *testing.T) {
	n := Int{big.NewInt(123)}
	b, err := MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	o, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, o) {
		t.Fatal("error unmarshalling Int", n, o)
	}

	// We must use 64 precision to perform the right comparison.
	// UnmarshalText generates a 64 precision float.
	// Check: https://github.com/golang/go/issues/45309
	f := Float{big.NewFloat(123.123).SetPrec(64)}
	b, err = MarshalJSON(f)
	if err != nil {
		t.Fatal(err)
	}
	of, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(f, of) {
		t.Fatal("error unmarshalling Float", f, of)
	}

}

func TestMarshalJSONDict(t *testing.T) {
	n := Dict{
		Pairs: Pairs{
			{Bool{true}, Int{big.NewInt(567)}},
		},
	}
	b, err := MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	out, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, out) {
		t.Fatal("error unmarshalling Dict")
	}

}

func TestMarshalJSONList(t *testing.T) {
	n1 := Bool{true}
	n2 := String{"testing!"}
	n3 := Int{big.NewInt(567)}

	n := List{
		Elements: []Node{n1, n2, n3},
	}

	b, err := MarshalJSON(n)
	if err != nil {
		t.Fatal(err)
	}
	out, err := UnmarshalJSON(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, out) {
		t.Fatal("error unmarshalling List")
	}

}

func TestMarshalCBORAll(t *testing.T) {
	n := Dict{
		Pairs: Pairs{
			{String{"string"}, String{"baz"}},
			{Int{big.NewInt(567)}, String{"baz"}},
			{String{"int1"}, Int{big.NewInt(567)}},
			{String{"bytes"}, Bytes{[]byte("asdf")}},
			{Bytes{[]byte("int2")}, Int{big.NewInt(567)}},
			{String{"dict"}, Dict{
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
	b, err := MarshalCBOR(n)
	if err != nil {
		t.Fatal(err)
	}
	out, err := UnmarshalCBOR(b)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, out) {
		t.Fatal("error CBOR decoding using IPLD encoders", n, out)
	}
}
