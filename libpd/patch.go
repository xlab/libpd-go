package libpd

import (
	"unsafe"

	"github.com/xlab/libpd-go/core"
)

type Patch interface {
	Name() string
	Dir() string
	DollarZero() int
	Close()
}

type pdPatch struct {
	name       string
	dir        string
	dollarZero int

	fileHandle     unsafe.Pointer
	instanceHandle int
}

func (p *pdPatch) Name() string {
	return p.name
}

func (p *pdPatch) Dir() string {
	return p.dir
}

func (p *pdPatch) DollarZero() int {
	return p.dollarZero
}

func (p *pdPatch) Close() {
	if p.fileHandle == nil {
		return
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(p.instanceHandle)
	orPanic(err)

	core.CloseFile(p.fileHandle)

	p.dollarZero = 0
	p.fileHandle = nil
}

func (i *Instance) OpenPatch(name, dir string) Patch {
	patch := &pdPatch{
		name: name, dir: dir,
		instanceHandle: i.handle,
	}
	pdMux.Lock()
	defer pdMux.Unlock()
	err := switchInstance(patch.instanceHandle)
	orPanic(err)

	handle := core.OpenFile(name+"\x00", dir+"\x00")
	if handle == nil {
		return nil
	}
	patch.fileHandle = handle
	patch.dollarZero = int(core.GetDollarZero(handle))
	i.patches[patch.dollarZero] = patch
	return patch
}

func (i *Instance) ClosePatch(handle int) {
	i.patchesMux.Lock()
	defer i.patchesMux.Unlock()
	if patch, ok := i.patches[handle]; ok {
		patch.Close()
		delete(i.patches, handle)
	}
}

func (i *Instance) CloseAllPatches() {
	i.patchesMux.Lock()
	defer i.patchesMux.Unlock()
	for handle, patch := range i.patches {
		patch.Close()
		delete(i.patches, handle)
	}
}
