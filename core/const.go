// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sat, 27 Aug 2016 22:58:56 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package core

/*
#cgo LDFLAGS: -lpd
#include "z_libpd.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// AtomType as declared in core/m_pd.h:175
type AtomType int32

// AtomType enumeration from core/m_pd.h:175
const (
	AtomNull     AtomType = iota
	AtomFloat    AtomType = 1
	AtomSymbol   AtomType = 2
	AtomPointer  AtomType = 3
	AtomSemi     AtomType = 4
	AtomComma    AtomType = 5
	AtomDeffloat AtomType = 6
	AtomDefsym   AtomType = 7
	AtomDollar   AtomType = 8
	AtomDollsym  AtomType = 9
	AtomGimme    AtomType = 10
	AtomCant     AtomType = 11
)
