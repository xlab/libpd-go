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
	AW             Word
	AType          AtomType
	ref9a368c09    *C.t_atom
	allocs9a368c09 interface{}
}

const wordSize = unsafe.Sizeof(uintptr(0))

type Word [wordSize]byte

func (w *Word) Float32() float32 {
	return (*[1]float32)(unsafe.Pointer(w))[0]
}

func (w *Word) Pointer() unsafe.Pointer {
	return unsafe.Pointer((*[1]uintptr)(unsafe.Pointer(w))[0])
}

// !! The code below has been automatically generated !!

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x Atom) PassValue() (C.t_atom, *cgoAllocMap) {
	if x.ref9a368c09 != nil {
		return *x.ref9a368c09, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *Atom) Deref() {
	if x.ref9a368c09 == nil {
		return
	}
	x.AType = (AtomType)(x.ref9a368c09.a_type)
	x.AW = *(*Word)(unsafe.Pointer(&x.ref9a368c09.a_w))
}

// Ref returns a reference.
func (x *Atom) Ref() *C.t_atom {
	if x == nil {
		return nil
	}
	return x.ref9a368c09
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *Atom) Free() {
	if x != nil && x.allocs9a368c09 != nil {
		x.allocs9a368c09.(*cgoAllocMap).Free()
		x.ref9a368c09 = nil
	}
}

// NewAtomRef initialises a new struct holding the reference to the originaitng C struct.
func NewAtomRef(ref unsafe.Pointer) *Atom {
	if ref == nil {
		return nil
	}
	obj := new(Atom)
	obj.ref9a368c09 = (*C.t_atom)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *Atom) PassRef() (*C.t_atom, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9a368c09 != nil {
		return x.ref9a368c09, nil
	}
	mem9a368c09 := allocAtomMemory(1)
	ref9a368c09 := (*C.t_atom)(mem9a368c09)
	allocs9a368c09 := new(cgoAllocMap)
	var ca_type_allocs *cgoAllocMap
	ref9a368c09.a_type, ca_type_allocs = (C.t_atomtype)(x.AType), cgoAllocsUnknown
	allocs9a368c09.Borrow(ca_type_allocs)

	var ca_w_allocs *cgoAllocMap
	ref9a368c09.a_w, ca_w_allocs = *(*C.union_word)(unsafe.Pointer(&x.AW)), cgoAllocsUnknown
	allocs9a368c09.Borrow(ca_w_allocs)

	x.ref9a368c09 = ref9a368c09
	x.allocs9a368c09 = allocs9a368c09
	return ref9a368c09, allocs9a368c09

}
