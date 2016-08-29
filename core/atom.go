package core

/*
#cgo LDFLAGS: -lpd
#include "z_libpd.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Atom as declared in core/m_pd.h:183
type Atom struct {
	AType  AtomType
	ref    *C.t_atom
	allocs interface{}
}

func (a *Atom) IsFloat() bool {
	return C.libpd_is_float(a.ref) > 0
}

func (a *Atom) IsSymbol() bool {
	return C.libpd_is_symbol(a.ref) > 0
}

func (a *Atom) Float() float32 {
	return float32(C.libpd_get_float(a.ref))
}

func (a *Atom) Symbol() string {
	cptr := C.libpd_get_symbol(a.ref)
	return C.GoString(cptr)
}

func (a *Atom) Next() *Atom {
	return &Atom{
		ref: C.libpd_next_atom(a.ref),
	}
}

func (a *Atom) SetFloat(v float32) {
	C.libpd_set_float(a.ref, C.float(v))
}

func NewAtomRef(ref unsafe.Pointer) *Atom {
	if ref == nil {
		return nil
	}
	return &Atom{
		ref: (*C.t_atom)(unsafe.Pointer(ref)),
	}
}

func (x *Atom) PassRef() (*C.t_atom, *cgoAllocMap) {
	return x.ref, nil
}
