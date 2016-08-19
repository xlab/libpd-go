package main

import (
	"log"
	"unsafe"

	libpd "github.com/xlab/libpd-go/core"
	"github.com/xlab/portaudio-go/portaudio"
)

type PureDataPlayer struct {
	frameSize  int
	channels   int
	fileHandle unsafe.Pointer
	outbuf     []float32
}

func InitPureDataPlayer(name string, dir string,
	channels, sampleRate, frameSize int, debug bool) *PureDataPlayer {
	if debug {
		libpd.SetPrintHook(func(msg string) {
			log.Print("> ", msg)
		})
	}
	if libpd.Init() != 0 {
		log.Println("[ERR] libpd.Init")
		return nil
	}
	if libpd.InitAudio(0, int32(channels), int32(sampleRate)) != 0 {
		log.Println("[ERR] libpd.InitAudio")
		return nil
	}
	fileHandle := libpd.OpenFile(name+"\x00", dir+"\x00")
	if fileHandle == nil {
		log.Println("[ERR] libpd.OpenFile")
		return nil
	}
	p := &PureDataPlayer{
		frameSize:  frameSize,
		channels:   channels,
		fileHandle: fileHandle,
		outbuf:     make([]float32, frameSize*channels),
	}
	return p
}

func (p *PureDataPlayer) Close() {
	libpd.CloseFile(p.fileHandle)
}

func (p *PureDataPlayer) StartDSP() {
	libpd.StartMessage(1)
	libpd.AddFloat(1.0)
	libpd.FinishMessage("pd\x00", "dsp\x00")
}

func (p *PureDataPlayer) StopDSP() {
	libpd.StartMessage(1)
	libpd.AddFloat(0.0)
	libpd.FinishMessage("pd\x00", "dsp\x00")
}

const (
	statusContinue = int32(portaudio.PaContinue)
	statusComplete = int32(portaudio.PaComplete)
	statusAbort    = int32(portaudio.PaAbort)
)

func (p *PureDataPlayer) StreamCallback(_ unsafe.Pointer, output unsafe.Pointer, sampleCount uint,
	_ *portaudio.StreamCallbackTimeInfo, _ portaudio.StreamCallbackFlags, _ unsafe.Pointer) int32 {

	if libpd.ProcessFloat(1, nil, p.outbuf) != 0 {
		log.Println("[ERR] libpd.ProcessFloat")
		return statusComplete
	}

	out := (*(*[1 << 32]float32)(output))[:sampleCount*audioChannels]
	copy(out, p.outbuf)

	return statusContinue
}
