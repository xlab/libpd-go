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
import "unsafe"

// Init function as declared in core/z_libpd.h:21
func Init() int32 {
	__ret := C.libpd_init()
	__v := (int32)(__ret)
	return __v
}

// ClearSearchPath function as declared in core/z_libpd.h:22
func ClearSearchPath() {
	C.libpd_clear_search_path()
}

// AddToSearchPath function as declared in core/z_libpd.h:23
func AddToSearchPath(sym string) {
	csym, _ := unpackPCharString(sym)
	C.libpd_add_to_search_path(csym)
}

// OpenFile function as declared in core/z_libpd.h:25
func OpenFile(basename string, dirname string) unsafe.Pointer {
	cbasename, _ := unpackPCharString(basename)
	cdirname, _ := unpackPCharString(dirname)
	__ret := C.libpd_openfile(cbasename, cdirname)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// CloseFile function as declared in core/z_libpd.h:26
func CloseFile(p unsafe.Pointer) {
	cp, _ := (unsafe.Pointer)(unsafe.Pointer(p)), cgoAllocsUnknown
	C.libpd_closefile(cp)
}

// GetDollarZero function as declared in core/z_libpd.h:27
func GetDollarZero(p unsafe.Pointer) int32 {
	cp, _ := (unsafe.Pointer)(unsafe.Pointer(p)), cgoAllocsUnknown
	__ret := C.libpd_getdollarzero(cp)
	__v := (int32)(__ret)
	return __v
}

// BlockSize function as declared in core/z_libpd.h:29
func BlockSize() int32 {
	__ret := C.libpd_blocksize()
	__v := (int32)(__ret)
	return __v
}

// InitAudio function as declared in core/z_libpd.h:30
func InitAudio(inchans int32, outchans int32, samplerate int32) int32 {
	cinchans, _ := (C.int)(inchans), cgoAllocsUnknown
	coutchans, _ := (C.int)(outchans), cgoAllocsUnknown
	csamplerate, _ := (C.int)(samplerate), cgoAllocsUnknown
	__ret := C.libpd_init_audio(cinchans, coutchans, csamplerate)
	__v := (int32)(__ret)
	return __v
}

// ProcessRaw function as declared in core/z_libpd.h:31
func ProcessRaw(inbuffer []float32, outbuffer []float32) int32 {
	cinbuffer, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inbuffer)).Data)), cgoAllocsUnknown
	coutbuffer, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outbuffer)).Data)), cgoAllocsUnknown
	__ret := C.libpd_process_raw(cinbuffer, coutbuffer)
	__v := (int32)(__ret)
	return __v
}

// ProcessShort function as declared in core/z_libpd.h:32
func ProcessShort(ticks int32, inbuffer []int16, outbuffer []int16) int32 {
	cticks, _ := (C.int)(ticks), cgoAllocsUnknown
	cinbuffer, _ := (*C.short)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inbuffer)).Data)), cgoAllocsUnknown
	coutbuffer, _ := (*C.short)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outbuffer)).Data)), cgoAllocsUnknown
	__ret := C.libpd_process_short(cticks, cinbuffer, coutbuffer)
	__v := (int32)(__ret)
	return __v
}

// ProcessFloat function as declared in core/z_libpd.h:34
func ProcessFloat(ticks int32, inbuffer []float32, outbuffer []float32) int32 {
	cticks, _ := (C.int)(ticks), cgoAllocsUnknown
	cinbuffer, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inbuffer)).Data)), cgoAllocsUnknown
	coutbuffer, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outbuffer)).Data)), cgoAllocsUnknown
	__ret := C.libpd_process_float(cticks, cinbuffer, coutbuffer)
	__v := (int32)(__ret)
	return __v
}

// ProcessDouble function as declared in core/z_libpd.h:36
func ProcessDouble(ticks int32, inbuffer []float64, outbuffer []float64) int32 {
	cticks, _ := (C.int)(ticks), cgoAllocsUnknown
	cinbuffer, _ := (*C.double)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&inbuffer)).Data)), cgoAllocsUnknown
	coutbuffer, _ := (*C.double)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outbuffer)).Data)), cgoAllocsUnknown
	__ret := C.libpd_process_double(cticks, cinbuffer, coutbuffer)
	__v := (int32)(__ret)
	return __v
}

// ArraySize function as declared in core/z_libpd.h:39
func ArraySize(name string) int32 {
	cname, _ := unpackPCharString(name)
	__ret := C.libpd_arraysize(cname)
	__v := (int32)(__ret)
	return __v
}

// ReadArray function as declared in core/z_libpd.h:41
func ReadArray(dest []float32, src string, offset int32, n int32) int32 {
	cdest, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&dest)).Data)), cgoAllocsUnknown
	csrc, _ := unpackPCharString(src)
	coffset, _ := (C.int)(offset), cgoAllocsUnknown
	cn, _ := (C.int)(n), cgoAllocsUnknown
	__ret := C.libpd_read_array(cdest, csrc, coffset, cn)
	__v := (int32)(__ret)
	return __v
}

// WriteArray function as declared in core/z_libpd.h:42
func WriteArray(dest string, offset int32, src []float32, n int32) int32 {
	cdest, _ := unpackPCharString(dest)
	coffset, _ := (C.int)(offset), cgoAllocsUnknown
	csrc, _ := (*C.float)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&src)).Data)), cgoAllocsUnknown
	cn, _ := (C.int)(n), cgoAllocsUnknown
	__ret := C.libpd_write_array(cdest, coffset, csrc, cn)
	__v := (int32)(__ret)
	return __v
}

// Bang function as declared in core/z_libpd.h:44
func Bang(recv string) int32 {
	crecv, _ := unpackPCharString(recv)
	__ret := C.libpd_bang(crecv)
	__v := (int32)(__ret)
	return __v
}

// Float function as declared in core/z_libpd.h:45
func Float(recv string, x float32) int32 {
	crecv, _ := unpackPCharString(recv)
	cx, _ := (C.float)(x), cgoAllocsUnknown
	__ret := C.libpd_float(crecv, cx)
	__v := (int32)(__ret)
	return __v
}

// Symbol function as declared in core/z_libpd.h:46
func Symbol(recv string, sym string) int32 {
	crecv, _ := unpackPCharString(recv)
	csym, _ := unpackPCharString(sym)
	__ret := C.libpd_symbol(crecv, csym)
	__v := (int32)(__ret)
	return __v
}

// SetFloat function as declared in core/z_libpd.h:48
func SetFloat(v []Atom, x float32) {
	cv, _ := unpackArgSAtom(v)
	cx, _ := (C.float)(x), cgoAllocsUnknown
	C.libpd_set_float(cv, cx)
	packSAtom(v, cv)
}

// SetSymbol function as declared in core/z_libpd.h:49
func SetSymbol(v []Atom, sym string) {
	cv, _ := unpackArgSAtom(v)
	csym, _ := unpackPCharString(sym)
	C.libpd_set_symbol(cv, csym)
	packSAtom(v, cv)
}

// List function as declared in core/z_libpd.h:50
func List(recv string, argc int32, argv []Atom) int32 {
	crecv, _ := unpackPCharString(recv)
	cargc, _ := (C.int)(argc), cgoAllocsUnknown
	cargv, _ := unpackArgSAtom(argv)
	__ret := C.libpd_list(crecv, cargc, cargv)
	packSAtom(argv, cargv)
	__v := (int32)(__ret)
	return __v
}

// Message function as declared in core/z_libpd.h:51
func Message(recv string, msg string, argc int32, argv []Atom) int32 {
	crecv, _ := unpackPCharString(recv)
	cmsg, _ := unpackPCharString(msg)
	cargc, _ := (C.int)(argc), cgoAllocsUnknown
	cargv, _ := unpackArgSAtom(argv)
	__ret := C.libpd_message(crecv, cmsg, cargc, cargv)
	packSAtom(argv, cargv)
	__v := (int32)(__ret)
	return __v
}

// StartMessage function as declared in core/z_libpd.h:53
func StartMessage(maxLength int32) int32 {
	cmaxLength, _ := (C.int)(maxLength), cgoAllocsUnknown
	__ret := C.libpd_start_message(cmaxLength)
	__v := (int32)(__ret)
	return __v
}

// AddFloat function as declared in core/z_libpd.h:54
func AddFloat(x float32) {
	cx, _ := (C.float)(x), cgoAllocsUnknown
	C.libpd_add_float(cx)
}

// AddSymbol function as declared in core/z_libpd.h:55
func AddSymbol(sym string) {
	csym, _ := unpackPCharString(sym)
	C.libpd_add_symbol(csym)
}

// FinishList function as declared in core/z_libpd.h:56
func FinishList(recv string) int32 {
	crecv, _ := unpackPCharString(recv)
	__ret := C.libpd_finish_list(crecv)
	__v := (int32)(__ret)
	return __v
}

// FinishMessage function as declared in core/z_libpd.h:57
func FinishMessage(recv string, msg string) int32 {
	crecv, _ := unpackPCharString(recv)
	cmsg, _ := unpackPCharString(msg)
	__ret := C.libpd_finish_message(crecv, cmsg)
	__v := (int32)(__ret)
	return __v
}

// Exists function as declared in core/z_libpd.h:59
func Exists(sym string) int32 {
	csym, _ := unpackPCharString(sym)
	__ret := C.libpd_exists(csym)
	__v := (int32)(__ret)
	return __v
}

// Bind function as declared in core/z_libpd.h:60
func Bind(sym string) unsafe.Pointer {
	csym, _ := unpackPCharString(sym)
	__ret := C.libpd_bind(csym)
	__v := *(*unsafe.Pointer)(unsafe.Pointer(&__ret))
	return __v
}

// Unbind function as declared in core/z_libpd.h:61
func Unbind(p unsafe.Pointer) {
	cp, _ := (unsafe.Pointer)(unsafe.Pointer(p)), cgoAllocsUnknown
	C.libpd_unbind(cp)
}

// IsFloat function as declared in core/z_libpd.h:63
func IsFloat(a []Atom) int32 {
	ca, _ := unpackArgSAtom(a)
	__ret := C.libpd_is_float(ca)
	packSAtom(a, ca)
	__v := (int32)(__ret)
	return __v
}

// IsSymbol function as declared in core/z_libpd.h:64
func IsSymbol(a []Atom) int32 {
	ca, _ := unpackArgSAtom(a)
	__ret := C.libpd_is_symbol(ca)
	packSAtom(a, ca)
	__v := (int32)(__ret)
	return __v
}

// GetFloat function as declared in core/z_libpd.h:65
func GetFloat(a []Atom) float32 {
	ca, _ := unpackArgSAtom(a)
	__ret := C.libpd_get_float(ca)
	packSAtom(a, ca)
	__v := (float32)(__ret)
	return __v
}

// GetSymbol function as declared in core/z_libpd.h:66
func GetSymbol(a []Atom) *byte {
	ca, _ := unpackArgSAtom(a)
	__ret := C.libpd_get_symbol(ca)
	packSAtom(a, ca)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// NextAtom function as declared in core/z_libpd.h:67
func NextAtom(a []Atom) *Atom {
	ca, _ := unpackArgSAtom(a)
	__ret := C.libpd_next_atom(ca)
	packSAtom(a, ca)
	__v := NewAtomRef(unsafe.Pointer(__ret))
	return __v
}

// SetPrintHook function as declared in core/z_libpd.h:77
func SetPrintHook(hook PrintHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_printhook(chook)
}

// SetBangHook function as declared in core/z_libpd.h:78
func SetBangHook(hook BangHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_banghook(chook)
}

// SetFloatHook function as declared in core/z_libpd.h:79
func SetFloatHook(hook FloatHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_floathook(chook)
}

// SetSymbolHook function as declared in core/z_libpd.h:80
func SetSymbolHook(hook SymbolHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_symbolhook(chook)
}

// SetListHook function as declared in core/z_libpd.h:81
func SetListHook(hook ListHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_listhook(chook)
}

// SetMessageHook function as declared in core/z_libpd.h:82
func SetMessageHook(hook MessageHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_messagehook(chook)
}

// NoteOn function as declared in core/z_libpd.h:84
func NoteOn(channel int32, pitch int32, velocity int32) int32 {
	cchannel, _ := (C.int)(channel), cgoAllocsUnknown
	cpitch, _ := (C.int)(pitch), cgoAllocsUnknown
	cvelocity, _ := (C.int)(velocity), cgoAllocsUnknown
	__ret := C.libpd_noteon(cchannel, cpitch, cvelocity)
	__v := (int32)(__ret)
	return __v
}

// ControlChange function as declared in core/z_libpd.h:85
func ControlChange(channel int32, controller int32, value int32) int32 {
	cchannel, _ := (C.int)(channel), cgoAllocsUnknown
	ccontroller, _ := (C.int)(controller), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.libpd_controlchange(cchannel, ccontroller, cvalue)
	__v := (int32)(__ret)
	return __v
}

// ProgramChange function as declared in core/z_libpd.h:86
func ProgramChange(channel int32, value int32) int32 {
	cchannel, _ := (C.int)(channel), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.libpd_programchange(cchannel, cvalue)
	__v := (int32)(__ret)
	return __v
}

// Pitchbend function as declared in core/z_libpd.h:87
func Pitchbend(channel int32, value int32) int32 {
	cchannel, _ := (C.int)(channel), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.libpd_pitchbend(cchannel, cvalue)
	__v := (int32)(__ret)
	return __v
}

// Aftertouch function as declared in core/z_libpd.h:88
func Aftertouch(channel int32, value int32) int32 {
	cchannel, _ := (C.int)(channel), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.libpd_aftertouch(cchannel, cvalue)
	__v := (int32)(__ret)
	return __v
}

// PolyAftertouch function as declared in core/z_libpd.h:89
func PolyAftertouch(channel int32, pitch int32, value int32) int32 {
	cchannel, _ := (C.int)(channel), cgoAllocsUnknown
	cpitch, _ := (C.int)(pitch), cgoAllocsUnknown
	cvalue, _ := (C.int)(value), cgoAllocsUnknown
	__ret := C.libpd_polyaftertouch(cchannel, cpitch, cvalue)
	__v := (int32)(__ret)
	return __v
}

// MIDIByte function as declared in core/z_libpd.h:90
func MIDIByte(port int32, byte int32) int32 {
	cport, _ := (C.int)(port), cgoAllocsUnknown
	cbyte, _ := (C.int)(byte), cgoAllocsUnknown
	__ret := C.libpd_midibyte(cport, cbyte)
	__v := (int32)(__ret)
	return __v
}

// SysEx function as declared in core/z_libpd.h:91
func SysEx(port int32, byte int32) int32 {
	cport, _ := (C.int)(port), cgoAllocsUnknown
	cbyte, _ := (C.int)(byte), cgoAllocsUnknown
	__ret := C.libpd_sysex(cport, cbyte)
	__v := (int32)(__ret)
	return __v
}

// SysRealtime function as declared in core/z_libpd.h:92
func SysRealtime(port int32, byte int32) int32 {
	cport, _ := (C.int)(port), cgoAllocsUnknown
	cbyte, _ := (C.int)(byte), cgoAllocsUnknown
	__ret := C.libpd_sysrealtime(cport, cbyte)
	__v := (int32)(__ret)
	return __v
}

// SetNoteOnHook function as declared in core/z_libpd.h:103
func SetNoteOnHook(hook NoteOnHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_noteonhook(chook)
}

// SetControlChangeHook function as declared in core/z_libpd.h:104
func SetControlChangeHook(hook ControlChangeHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_controlchangehook(chook)
}

// SetProgramChangeHook function as declared in core/z_libpd.h:105
func SetProgramChangeHook(hook ProgramChangeHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_programchangehook(chook)
}

// SetPitchbendHook function as declared in core/z_libpd.h:106
func SetPitchbendHook(hook PitchbendHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_pitchbendhook(chook)
}

// SetAftertouchHook function as declared in core/z_libpd.h:107
func SetAftertouchHook(hook AftertouchHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_aftertouchhook(chook)
}

// SetPolyAftertouchHook function as declared in core/z_libpd.h:108
func SetPolyAftertouchHook(hook PolyAftertouchHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_polyaftertouchhook(chook)
}

// SetMIDIByteHook function as declared in core/z_libpd.h:109
func SetMIDIByteHook(hook MIDIByteHook) {
	chook, _ := hook.PassValue()
	C.libpd_set_midibytehook(chook)
}

// PdInstanceNew function as declared in core/m_pd.h:801
func PdInstanceNew() *PdInstance {
	__ret := C.pdinstance_new()
	__v := *(**PdInstance)(unsafe.Pointer(&__ret))
	return __v
}

// PdSetInstance function as declared in core/m_pd.h:802
func PdSetInstance(x *PdInstance) {
	cx, _ := (*C.struct__pdinstance)(unsafe.Pointer(x)), cgoAllocsUnknown
	C.pd_setinstance(cx)
}

// PdInstanceFree function as declared in core/m_pd.h:803
func PdInstanceFree(x *PdInstance) {
	cx, _ := (*C.struct__pdinstance)(unsafe.Pointer(x)), cgoAllocsUnknown
	C.pdinstance_free(cx)
}
