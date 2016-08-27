package libpd

import "github.com/xlab/libpd-go/core"

type Atom interface{}

type UnknownAtom struct{}
type NullAtom struct{}

func convertAtomList(atoms []core.Atom) []Atom {
	list := make([]Atom, len(atoms))
	for i := range atoms {
		atoms[i].Deref()
		switch atoms[i].AType {
		case core.AtomNull:
			list[i] = NullAtom{}
		case core.AtomFloat:
			list[i] = atoms[i].AW.Float32()
		case core.AtomSymbol:
			list[i] = atoms[i].AW.Symbol()
		case core.AtomPointer:
			list[i] = atoms[i].AW.Pointer()
		default:
			list[i] = UnknownAtom{}
		}
	}
	return list
}
