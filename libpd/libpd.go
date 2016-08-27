package libpd

import (
	"fmt"
	"sync"

	"github.com/xlab/libpd-go/core"
)

var (
	pdIdx     int
	pdCurrent int

	pdInit bool
	pdMap  = make(map[int]*core.PdInstance)
	pdMux  = new(sync.RWMutex)
)

type Instance struct {
	handle   int
	initDone bool

	printHook   core.PrintHook
	bangHook    core.BangHook
	floatHook   core.FloatHook
	symbolHook  core.SymbolHook
	listHook    core.ListHook
	messageHook core.MessageHook

	noteOnHook         core.NoteOnHook
	controlChangeHook  core.ControlChangeHook
	programChangeHook  core.ProgramChangeHook
	pitchbendHook      core.PitchbendHook
	aftertouchHook     core.AftertouchHook
	polyAftertouchHook core.PolyAftertouchHook
	midiByteHook       core.MIDIByteHook
}

func (i *Instance) Destroy() {
	if i.handle > 0 {
		pdMux.Lock()
		if pdCurrent == i.handle {
			// switch to another
			for handle, ref := range pdMap {
				pdCurrent = handle
				core.PdSetInstance(ref)
				break
			}
			if pdCurrent == i.handle {
				pdCurrent = 0 // just reset
			}
		}
		ref, ok := pdMap[i.handle]
		if !ok { // gone already
			i.handle = 0
			pdMux.Unlock()
			return
		}
		core.PdInstanceFree(ref)
		delete(pdMap, i.handle)
		i.handle = 0
		pdMux.Unlock()
	}
}

func NewInstance() *Instance {
	var handle int

	pdMux.Lock()
	pdIdx++
	handle = pdIdx
	ref := core.PdInstanceNew()
	pdMap[pdIdx] = ref
	pdMux.Unlock()

	return &Instance{
		handle: handle,
	}
}

func switchInstance(handle int) error {
	// get a real instance reference
	ref, ok := pdMap[handle]
	if !ok {
		err := fmt.Errorf("internal error: switchInstance: handle %d not registered", handle)
		return err
	}
	// switch libpd to use that instance
	pdCurrent = handle
	core.PdSetInstance(ref)
	return nil
}

func (i *Instance) Init(inCh, outCh int, sampleRate int) error {
	pdMux.Lock()
	defer pdMux.Unlock()
	if err := switchInstance(i.handle); err != nil {
		return err
	}

	if i.printHook != nil {
		core.SetPrintHook(i.printHook)
	}
	if i.bangHook != nil {
		core.SetBangHook(i.bangHook)
	}
	if i.floatHook != nil {
		core.SetFloatHook(i.floatHook)
	}
	if i.symbolHook != nil {
		core.SetSymbolHook(i.symbolHook)
	}
	if i.listHook != nil {
		core.SetListHook(i.listHook)
	}
	if i.messageHook != nil {
		core.SetMessageHook(i.messageHook)
	}
	if i.noteOnHook != nil {
		core.SetNoteOnHook(i.noteOnHook)
	}
	if i.controlChangeHook != nil {
		core.SetControlChangeHook(i.controlChangeHook)
	}
	if i.programChangeHook != nil {
		core.SetProgramChangeHook(i.programChangeHook)
	}
	if i.pitchbendHook != nil {
		core.SetPitchbendHook(i.pitchbendHook)
	}
	if i.aftertouchHook != nil {
		core.SetAftertouchHook(i.aftertouchHook)
	}
	if i.polyAftertouchHook != nil {
		core.SetPolyAftertouchHook(i.polyAftertouchHook)
	}
	if i.midiByteHook != nil {
		core.SetMIDIByteHook(i.midiByteHook)
	}
	if !pdInit {
		core.Init()
		pdInit = true
	}
	if ret := core.InitAudio(int32(inCh), int32(outCh), int32(sampleRate)); ret != 0 {
		err := fmt.Errorf("pd InitAudio failed with ret: %d", ret)
		return err
	}
	i.initDone = true
	return nil
}

func (i *Instance) SetPrintHook(fn func(recv string)) {
	i.printHook = fn
}

func (i *Instance) SetBangHook(fn func(recv string)) {
	i.bangHook = fn
}

func (i *Instance) SetFloatHook(fn func(recv string, x float32)) {
	i.floatHook = fn
}

func (i *Instance) SetSymbolHook(fn func(recv string, sym string)) {
	i.symbolHook = fn
}

func (i *Instance) SetListHook(fn func(recv string, argv ...Atom)) {
	i.listHook = func(recv string, argc int32, argv []core.Atom) {
		args := convertAtomList(argv)
		fn(recv, args...)
	}
}

func (i *Instance) SetMessageHook(fn func(recv string, msg string, argv ...Atom)) {
	i.messageHook = func(recv string, msg string, argc int32, argv []core.Atom) {
		args := convertAtomList(argv)
		fn(recv, msg, args...)
	}
}

func (i *Instance) SetNoteOnHook(fn func(channel, pitch, velocity int)) {
	i.noteOnHook = func(channel int32, pitch int32, velocity int32) {
		fn(int(channel), int(pitch), int(velocity))
	}
}

func (i *Instance) SetControlChangeHook(fn func(channel, controller, value int)) {
	i.controlChangeHook = func(channel int32, controller int32, value int32) {
		fn(int(channel), int(controller), int(value))
	}
}

func (i *Instance) SetProgramChangeHook(fn func(channel, value int)) {
	i.programChangeHook = func(channel int32, value int32) {
		fn(int(channel), int(value))
	}
}

func (i *Instance) SetPitchbendHook(fn func(channel, value int)) {
	i.pitchbendHook = func(channel int32, value int32) {
		fn(int(channel), int(value))
	}
}

func (i *Instance) SetAftertouchHook(fn func(channel, value int)) {
	i.aftertouchHook = func(channel int32, value int32) {
		fn(int(channel), int(value))
	}
}

func (i *Instance) SetPolyAftertouchHook(fn func(channel, pitch, value int)) {
	i.polyAftertouchHook = func(channel, pitch, value int32) {
		fn(int(channel), int(pitch), int(value))
	}
}

func (i *Instance) SetMIDIByteHook(fn func(port int, v byte)) {
	i.midiByteHook = func(port int32, v int32) {
		fn(int(port), byte(v))
	}
}
