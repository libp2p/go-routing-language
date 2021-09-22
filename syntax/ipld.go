package syntax

import (
	"fmt"
	"math/big"

	"github.com/ipld/go-ipld-prime"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

// ipldTypeTags used in Node_IPLD type
var ipldTypeTags = []string{
	"String_IPLD",
	"Bytes_IPLD",
	"Float_IPLD",
	"Int_IPLD",
	"Bool_IPLD",
	"Dict_IPLD",
	"List_IPLD",
	"Predicate_IPLD",
}

// FromIPLD transforms an IPLD Node into its xr.Node representation
func FromIPLD(n ipld.Node) (Node, error) {
	switch n1 := n.(type) {
	case xipld.Bytes_IPLD:
		b, err := n1.AsBytes()
		if err != nil {
			return nil, err
		}
		return Bytes{b}, nil

	case xipld.Bool_IPLD:
		b, err := n1.AsBool()
		if err != nil {
			return nil, err
		}
		return Bool{b}, nil

	case xipld.String_IPLD:
		b, err := n1.AsString()
		if err != nil {
			return nil, err
		}
		return String{b}, nil

	case xipld.Int_IPLD:
		b, err := n1.AsInt()
		if err != nil {
			return nil, err
		}
		return Int{big.NewInt(b)}, nil

	case xipld.Float_IPLD:
		b, err := n1.AsFloat()
		if err != nil {
			return nil, err
		}
		return Float{big.NewFloat(b).SetPrec(64)}, nil

	case xipld.List_IPLD:
		return fromIPLDToList(n1)

	case xipld.Dict_IPLD:
		return fromIPLDToDict(n1)

	case xipld.Predicate_IPLD:
		return fromIPLDToPredicate(n1)

	case xipld.Node_IPLD:
		for _, k := range ipldTypeTags {
			// Check which type is Node_IPLD to convert into the right IPLD Node
			nt, err := n1.LookupByString(k)
			if err == nil {
				return FromIPLD(nt)
			}
		}
		return nil, fmt.Errorf("Node_IPLD has no valid type inside")
	}

	return nil, fmt.Errorf("IPLD type for xr.Node not found. Can't convert")
}

// Creates a List in XR from List_IPLD
func fromIPLDToList(n xipld.List_IPLD) (List, error) {
	// Get elements
	els := make([]Node, 0)
	li := n.FieldElements().Iterator()
	for !li.Done() {
		_, enode := li.Next()
		n, err := FromIPLD(enode)
		if err != nil {
			return List{}, err
		}
		// Append element
		els = append(els, n)
	}

	return List{Elements: els}, nil
}

// Create Dict in XR from Dict_IPLD
func fromIPLDToDict(n xipld.Dict_IPLD) (Dict, error) {
	// Get pairs
	pairs := make([]Pair, 0)
	li := n.FieldPairs().Iterator()
	for !li.Done() {
		_, enode := li.Next()
		// Get key and convert to xr.Node
		ikey := enode.FieldKey()
		k, err := FromIPLD(ikey)
		if err != nil {
			return Dict{}, err
		}
		// Get value and convert to xr.Node
		ivalue := enode.FieldValue()
		v, err := FromIPLD(ivalue)
		if err != nil {
			return Dict{}, err
		}
		// Append pair
		pairs = append(pairs, Pair{Key: k, Value: v})
	}
	return Dict{Pairs: pairs}, nil
}

// Create Predicate in XR from Predicate_IPLD
func fromIPLDToPredicate(n xipld.Predicate_IPLD) (Predicate, error) {
	// Get positional
	els := make([]Node, 0)
	li := n.FieldPositional().Iterator()
	for !li.Done() {
		_, enode := li.Next()
		n, err := FromIPLD(enode)
		if err != nil {
			return Predicate{}, err
		}
		// Append element
		els = append(els, n)
	}

	// Get named
	pairs := make([]Pair, 0)
	pi := n.FieldNamed().Iterator()
	for !pi.Done() {
		_, enode := pi.Next()
		// Get key and convert to xr.Node
		ikey := enode.FieldKey()
		k, err := FromIPLD(ikey)
		if err != nil {
			return Predicate{}, err
		}
		// Get value and convert to xr.Node
		ivalue := enode.FieldValue()
		v, err := FromIPLD(ivalue)
		if err != nil {
			return Predicate{}, err
		}
		// Append pair
		pairs = append(pairs, Pair{Key: k, Value: v})
	}
	t, err := n.FieldTag().AsString()
	if err != nil {
		return Predicate{}, err
	}
	return Predicate{Tag: t, Positional: els, Named: pairs}, nil
}
