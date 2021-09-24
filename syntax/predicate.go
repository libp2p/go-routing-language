package syntax

import (
	"io"

	"github.com/ipld/go-ipld-prime"
	xipld "github.com/libp2p/go-routing-language/syntax/ipld"
)

// Predicate models a function invocation with named and positional arguments, corresponding to the syntax:
//   tag(a1, a2, ...; n1=v1, n2=v2, ...)
type Predicate struct {
	Tag        string
	Positional Nodes
	Named      Pairs // the keys in each pair must be unique wrt IsEqual
}

func (p Predicate) GetNamed(name string) Node {
	for _, p := range p.Named {
		s, ok := p.Key.(String)
		if ok && s.Value == name {
			return p.Value
		}
	}
	return nil
}

func (p Predicate) WritePretty(w io.Writer) error {
	if _, err := w.Write([]byte(p.Tag)); err != nil {
		return err
	}
	if _, err := w.Write([]byte{'('}); err != nil {
		return err
	}
	u := IndentWriter(w)
	if _, err := u.Write([]byte{'\n'}); err != nil {
		return err
	}
	for i, n := range p.Positional {
		if err := n.WritePretty(u); err != nil {
			return err
		}
		if i+1 == len(p.Positional) && len(p.Named) == 0 {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		} else {
			if _, err := u.Write([]byte("\n")); err != nil {
				return err
			}
		}
	}
	for i, a := range p.Named {
		if err := a.WritePretty(u, "="); err != nil {
			return err
		}
		if i+1 == len(p.Named) {
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		} else {
			if _, err := u.Write([]byte("\n")); err != nil {
				return err
			}
		}
	}
	if _, err := w.Write([]byte{')'}); err != nil {
		return err
	}
	return nil
}

func (p Predicate) ToIPLD() (ipld.Node, error) {
	// Initialize Predicate
	dbuild := xipld.Type.Predicate_IPLD.NewBuilder()
	ma, err := dbuild.BeginMap(-1)
	if err != nil {
		return nil, err
	}
	// Build tag
	tagAsm, err := ma.AssembleEntry("Tag")
	if err != nil {
		return nil, err
	}
	err = tagAsm.AssignString(p.Tag)
	if err != nil {
		return nil, err
	}
	// Build positional arguments
	posIPLD, err := p.Positional.ToIPLD()
	if err != nil {
		return nil, err
	}
	// Assign lists of positional arguments to predicate
	posAsm, err := ma.AssembleEntry("Positional")
	if err != nil {
		return nil, err
	}
	err = posAsm.AssignNode(posIPLD)
	if err != nil {
		return nil, err
	}
	// Build named arguments
	namedIPLD, err := p.Named.ToIPLD()
	if err != nil {
		return nil, err
	}
	// Assign lists of named arguments to predicate
	namedAsm, err := ma.AssembleEntry("Named")
	if err != nil {
		return nil, err
	}
	err = namedAsm.AssignNode(namedIPLD)
	if err != nil {
		return nil, err
	}
	// Finish building predicate
	if err := ma.Finish(); err != nil {
		return nil, err
	}
	return dbuild.Build(), nil
}

// toNode_IPLD convert into IPLD Node of dynamic type NODE_IPLD
func (p Predicate) toNode_IPLD() (ipld.Node, error) {
	t := xipld.Type.Node_IPLD.NewBuilder()
	ma, err := t.BeginMap(-1)
	if err != nil {
		return nil, err
	}
	asm, err := ma.AssembleEntry("Predicate_IPLD")
	if err != nil {
		return nil, err
	}
	nd, err := p.ToIPLD()
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

func IsEqualPredicate(x, y Predicate) bool {
	return x.Tag == y.Tag && AreSameNodes(x.Positional, y.Positional) && AreSamePairs(x.Named, y.Named)
}
