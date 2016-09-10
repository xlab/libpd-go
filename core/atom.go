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
type Atom C.t_atom

func (a *Atom) IsFloat() bool {
	return C.libpd_is_float((*C.t_atom)(a)) > 0
}

func (a *Atom) IsSymbol() bool {
	return C.libpd_is_symbol((*C.t_atom)(a)) > 0
}

func (a *Atom) Float() float32 {
	return float32(C.libpd_get_float((*C.t_atom)(a)))
}

func (a *Atom) Symbol() string {
	cptr := C.libpd_get_symbol((*C.t_atom)(a))
	return C.GoString(cptr)
}

func (a *Atom) Next() *Atom {
	atom := (*C.t_atom)(a)
	nextAtom := C.libpd_next_atom(atom)
	return (*Atom)(unsafe.Pointer(nextAtom))
}

func (a *Atom) SetFloat(v float32) {
	C.libpd_set_float((*C.t_atom)(a), C.float(v))
}

func NewAtomRef(ref unsafe.Pointer) *Atom {
	if ref == nil {
		return nil
	}
	return (*Atom)(ref)
}

func (x *Atom) PassRef() (*C.t_atom, *cgoAllocMap) {
	return (*C.t_atom)(unsafe.Pointer(x)), nil
}
