package syntax

import (
	"bytes"
	"math/big"
	"testing"

	cbor "github.com/ipld/go-ipld-prime/codec/dagcbor"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

func TestBytesIPLD(t *testing.T) {
	b := Bytes{[]byte("test")}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e IPLD Bytes transformation failed", b, bo)
	}
}

func TestBoolIPLD(t *testing.T) {
	b := Bool{true}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e IPLD Bool transformation failed", b, bo)
	}
}

func TestStringIPLD(t *testing.T) {
	b := String{"testing"}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e IPLD String transformation failed", b, bo)
	}
}

func TestIntIPLD(t *testing.T) {
	b := Int{big.NewInt(123)}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e IPLD Int transformation failed", b, bo)
	}
}

func TestFloatIPLD(t *testing.T) {
	b := Float{big.NewFloat(123.123).SetPrec(64)}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e IPLD Float transformation failed", b, bo)
	}
}

func TestListIPLD(t *testing.T) {
	n1 := Bool{true}
	n2 := String{"testing!"}
	n3 := Bytes{[]byte("test")}
	n4 := Int{big.NewInt(567)}

	b := List{
		Elements: []Node{n1, n2, n3, n4},
	}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e IPLD List transformation failed", b, bo)
	}
}

func TestDictIPLD(t *testing.T) {
	b := Dict{
		Pairs: Pairs{
			{Bool{true}, Int{big.NewInt(567)}},
		},
	}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e Dict List transformation failed", b, bo)
	}

}

func TestPredicateIPLD(t *testing.T) {
	b := Predicate{
		Tag: "tag",
		Positional: Nodes{
			Bool{true},
			Int{big.NewInt(567)},
		},
		Named: Pairs{
			{Bool{true}, Int{big.NewInt(567)}},
		},
	}
	bi, err := b.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	bo, err := FromIPLD(bi)
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(b, bo) {
		t.Fatal("e2e Predicate transformation failed", b, bo)
	}

}

func TestCBORMarshalCBOR(t *testing.T) {
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
		},
	}
	in, err := n.ToIPLD()
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	err = cbor.Encode(in, &buf)
	if err != nil {
		t.Fatal(err)
	}
	b := buf.Bytes()

	noipld := xipld.Type.Dict_IPLD.NewBuilder()
	err = cbor.Decode(noipld, bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	no, err := FromIPLD(noipld.Build())
	if err != nil {
		t.Fatal(err)
	}
	if !IsEqual(n, no) {
		t.Fatal("Marshalled Bytes nodes not equal")
	}

}
func TestWrongTypes(t *testing.T) {
	// TODO:
}

func TestToIPLD_Nodes(t *testing.T) {
	// TODO:
}
