package syntax

import (
	"fmt"
	"io"

	"github.com/ipld/go-ipld-prime"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

type Node interface {
	WritePretty(w io.Writer) error   // Pretty writes the node
	ToIPLD() (ipld.Node, error)      // Converts xr.Node into its corresponding IPLD Node type
	toNode_IPLD() (ipld.Node, error) // Convert into IPLD Node of dynamic type NODE_IPLD
}

type Nodes []Node

func (ns Nodes) IndexOf(element Node) int {
	for i, p := range ns {
		if IsEqual(p, element) {
			return i
		}
	}
	return -1
}

func (ns Nodes) ToIPLD() (ipld.Node, error) {
	// Build elements
	lbuild := xipld.Type.Nodes_IPLD.NewBuilder()
	// NOTE: We can assign here directly the size of Pairs instead of -1
	la, err := lbuild.BeginList(-1)
	if err != nil {
		return nil, err
	}
	// For each pair
	for _, e := range ns {

		// Add element to the list of nodes
		n, err := e.toNode_IPLD()
		if err != nil {
			return nil, err
		}
		// la.AssembleValue is Node_IPLD Assembler. Need to assemble a node
		if err := la.AssembleValue().AssignNode(n); err != nil {
			return nil, fmt.Errorf("error assembling value: %s", err)
		}
	}
	// Finish list building
	if err := la.Finish(); err != nil {
		return nil, err
	}
	return lbuild.Build(), nil
}

// AreSameNodes compairs to lists of key/values for set-wise equality (order independent).
func AreSameNodes(x, y Nodes) bool {
	if len(x) != len(y) {
		return false
	}
	for _, x := range x {
		if i := y.IndexOf(x); i < 0 {
			return false
		}
	}
	return true
}
