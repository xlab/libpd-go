package libpd

import "github.com/xlab/libpd-go/core"

type Atom interface{}

var UnknownAtom Atom = struct{}{}

func convertAtomList(atomList *core.Atom) []Atom {
	list := make([]Atom, 0, 10)
	for atomList != nil {
		switch {
		case atomList.IsFloat():
			list = append(list, atomList.Float())
		case atomList.IsSymbol():
			list = append(list, atomList.Symbol())
		default:
			list = append(list, UnknownAtom)
		}
		atomList = atomList.Next()
	}
	return list
}
