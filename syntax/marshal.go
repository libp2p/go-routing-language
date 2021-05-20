package syntax

import (
	"bytes"

	cbor "github.com/ipld/go-ipld-prime/codec/dagcbor"
	json "github.com/ipld/go-ipld-prime/codec/dagjson"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

// MarshalJSON syntactic representation
func MarshalJSON(n Node) ([]byte, error) {
	in, err := n.toNode_IPLD()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Encode(in, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON syntactic representation
func UnmarshalJSON(r []byte) (Node, error) {
	n := xipld.Type.Node_IPLD.NewBuilder()
	err := json.Decode(n, bytes.NewReader(r))
	if err != nil {
		return nil, err
	}
	return FromIPLD(n.Build())
}

// MarshalCBOR serializes syntactic nodes in CBOR using its IPLD capabilities
func MarshalCBOR(n Node) ([]byte, error) {
	in, err := n.toNode_IPLD()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = cbor.Encode(in, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalCBOR deserializes syntactic nodes from CBOR using its IPLD capabilities
func UnmarshalCBOR(r []byte) (Node, error) {
	n := xipld.Type.Node_IPLD.NewBuilder()
	err := cbor.Decode(n, bytes.NewReader(r))
	if err != nil {
		return nil, err
	}
	return FromIPLD(n.Build())
}
