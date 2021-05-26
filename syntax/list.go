package syntax

import (
	"io"

	"github.com/ipld/go-ipld-prime"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

// List is a set of (uniquely) elements.
type List struct {
	Elements Nodes
}

func (s List) Copy() List {
	e := make(Nodes, len(s.Elements))
	copy(e, s.Elements)
	return List{
		Elements: e,
	}
}

func (s List) Len() int {
	return len(s.Elements)
}

func (s List) WritePretty(w io.Writer) error {
	if _, err := w.Write([]byte{'['}); err != nil {
		return err
	}
	u := IndentWriter(w)
	if _, err := u.Write([]byte{'\n'}); err != nil {
		return err
	}
	for i, p := range s.Elements {
		if err := p.WritePretty(u); err != nil {
			return err
		}
		if i+1 == len(s.Elements) {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		} else {
			if _, err := u.Write([]byte("\n")); err != nil {
				return err
			}
		}
	}
	if _, err := w.Write([]byte{']'}); err != nil {
		return err
	}
	return nil
}

func IsEqualList(x, y List) bool {
	return AreSameNodes(x.Elements, y.Elements)
}

// ToIPLD converts xr.Node into its corresponding IPLD Node type
func (s List) ToIPLD() (ipld.Node, error) {
	// NOTE: Consider adding multierr throughout this whole function
	// Initialize Dict
	sbuild := xipld.Type.List_IPLD.NewBuilder()
	ma, err := sbuild.BeginMap(-1)
	if err != nil {
		return nil, err
	}
	// Build elements
	elemsIPLD, err := s.Elements.ToIPLD()
	if err != nil {
		return nil, err
	}
	// Assign elements to set
	psasm, err := ma.AssembleEntry("Elements")
	if err != nil {
		return nil, err
	}
	err = psasm.AssignNode(elemsIPLD)
	if err != nil {
		return nil, err
	}
	// Finish elements building
	if err := ma.Finish(); err != nil {
		return nil, err
	}
	return sbuild.Build(), nil
}

// toNode_IPLD convert into IPLD Node of dynamic type NODE_IPLD
func (s List) toNode_IPLD() (ipld.Node, error) {
	t := xipld.Type.Node_IPLD.NewBuilder()
	ma, err := t.BeginMap(-1)
	asm, err := ma.AssembleEntry("List_IPLD")
	if err != nil {
		return nil, err
	}
	nd, err := s.ToIPLD()
	if err != nil {
		return nil, err
	}
	err = asm.AssignNode(nd)
	if err != nil {
		return nil, err
	}
	if err := ma.Finish(); err != nil {
		return nil, err
	}
	return t.Build(), nil
}
