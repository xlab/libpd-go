package libpd

import (
	"fmt"
	"log"
	"sync"
	"unsafe"

	"github.com/xlab/libpd-go/core"
)

var (
	pdIdx     int
	pdCurrent int

	pdInit bool
	pdMap  = make(map[int]*core.PdInstance)
	pdMux  = new(sync.Mutex)

	pdBindMap    = make(map[string]unsafe.Pointer)
	pdBindMapMux = new(sync.RWMutex)
)

type Instance struct {
	handle     int
	initDone   bool
	msgMux     *sync.Mutex
	patches    map[int]Patch
	patchesMux *sync.RWMutex

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
		handle:     handle,
		msgMux:     new(sync.Mutex),
		patches:    make(map[int]Patch, 8),
		patchesMux: new(sync.RWMutex),
	}
}

func switchInstance(handle int) error {
	if pdCurrent == handle {
		return nil // no-op
	}
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
	err := switchInstance(i.handle)
	orPanic(err)

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
		err := fmt.Errorf("audio init failed: %d", ret)
		return err
	}
	i.initDone = true
	return nil
}

// SendMessage sends a message with no args, this opeation
// blocks until instance is busy processing another message.
func (i *Instance) SendMessage(recv, msg string) bool {
	if !i.initDone {
		return false
	}
	i.msgMux.Lock()
	defer i.msgMux.Unlock()

	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	ret := core.Message(recv+"\x00", msg+"\x00", 0, nil)
	return ret == 0
}

// NewMessageQueue creates a message stream, returns a channel for message's arguments.
// The channel will be automatically closed when number of arguments sent hits the maxArgs limit,
// however if there is less args than that, you must close the channel explicitly.
//
// For example: maxArgs=1, you send 1 value of type float32, the channel is closed and the message is sent.
// Another example: maxArgs=2, you send 1 value of type string, you must close the channel, otherwise the message
// session will stuck awaiting and will prevent you from creating another one.
func (i *Instance) NewMessageQueue(recv, msg string, maxArgs int) (chan<- Atom, bool) {
	if !i.initDone {
		return nil, false
	}
	if maxArgs == 0 {
		return nil, i.SendMessage(recv, msg)
	}
	pdMux.Lock()
	if err := switchInstance(i.handle); err != nil {
		pdMux.Unlock()
		orPanic(err)
	}
	ret := core.StartMessage(int32(maxArgs))
	pdMux.Unlock()
	if ret != 0 {
		return nil, false
	}

	ch := make(chan Atom, maxArgs)
	go func() {
		i.msgMux.Lock()
		defer i.msgMux.Unlock()
		var processed int
		for atom := range ch {
			processQueuedAtom(i.handle, atom)
			processed++
			if processed >= maxArgs {
				close(ch)
			}
		}
		pdMux.Lock()
		defer pdMux.Unlock()
		err := switchInstance(i.handle)
		orPanic(err)
		core.FinishMessage(recv+"\x00", msg+"\x00")
	}()
	return ch, true
}

func processQueuedAtom(handle int, atom Atom) {
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(handle)
	orPanic(err)

	switch v := atom.(type) {
	case float32:
		core.AddFloat(v)
	case string:
		core.AddSymbol(v + "\x00")
	case int:
		core.AddFloat(float32(v))
	case int8:
		core.AddFloat(float32(v))
	case int16:
		core.AddFloat(float32(v))
	case int32:
		core.AddFloat(float32(v))
	case int64:
		core.AddFloat(float32(v))
	case uint:
		core.AddFloat(float32(v))
	case uint8:
		core.AddFloat(float32(v))
	case uint16:
		core.AddFloat(float32(v))
	case uint32:
		core.AddFloat(float32(v))
	case uint64:
		core.AddFloat(float32(v))
	case float64:
		core.AddFloat(float32(v))
	default:
		// skip unknown arg types
	}
}

func orPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func (i *Instance) ProcessRaw(inBuffer, outBuffer []float32) bool {
	if !i.initDone {
		return false
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	ret := core.ProcessRaw(inBuffer, outBuffer)
	return ret == 0
}

func (i *Instance) ProcessInt16(ticks int, inBuffer, outBuffer []int16) bool {
	if !i.initDone {
		return false
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	ret := core.ProcessShort(int32(ticks), inBuffer, outBuffer)
	return ret == 0
}

func (i *Instance) ProcessFloat32(ticks int, inBuffer, outBuffer []float32) bool {
	if !i.initDone {
		return false
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	ret := core.ProcessFloat(int32(ticks), inBuffer, outBuffer)
	return ret == 0
}

func (i *Instance) ProcessFloat64(ticks int, inBuffer, outBuffer []float64) bool {
	if !i.initDone {
		return false
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	ret := core.ProcessDouble(int32(ticks), inBuffer, outBuffer)
	return ret == 0
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
	i.listHook = func(recv string, argc int32, argv *core.Atom) {
		args := convertAtomList(argv, int(argc))
		fn(recv, args...)
	}
}

func (i *Instance) SetMessageHook(fn func(recv string, msg string, argv ...Atom)) {
	i.messageHook = func(recv string, msg string, argc int32, argv *core.Atom) {
		args := convertAtomList(argv, int(argc))
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

func (i *Instance) AddToSearchPath(path string) {
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	core.AddToSearchPath(path + "\x00")
}

func (i *Instance) ClearSearchPath() {
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	core.ClearSearchPath()
}

// Bind subscribes to messages sent to the given symbol.
// The call Bind("foo") adds an object to the patch that behaves much like [r foo],
// with the output being passed on to the various message hooks of libpd.
// The call to Bind() should take place after the call to Init().
func (i *Instance) Bind(sym string) {
	if !i.initDone {
		return
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	pdBindMapMux.Lock()
	if _, ok := pdBindMap[sym]; ok {
		// bound already
		log.Println("libpd: symbol", sym, "is already bound (global).")
		pdBindMapMux.Unlock()
		return
	}
	pdBindMap[sym] = core.Bind(sym + "\x00")
	pdBindMapMux.Unlock()
}

// Unbind deletes the receiver objects created by Bind().
func (i *Instance) Unbind(sym string) {
	if !i.initDone {
		return
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(i.handle)
	orPanic(err)

	pdBindMapMux.Lock()
	ptr, ok := pdBindMap[sym]
	if !ok {
		pdBindMapMux.Unlock()
		return
	}
	core.Unbind(ptr)
	delete(pdBindMap, sym)
	pdBindMapMux.Unlock()
}
