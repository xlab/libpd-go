package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	"github.com/xlab/libpd-go/libpd"
	"github.com/xlab/portaudio-go/portaudio"
)

const (
	bitDepth      = 16
	audioChannels = 2
	sampleRate    = 44100
	sampleFormat  = portaudio.PaFloat32

	frameSize = 64
	ticks     = 4 // 256 frames
)

var (
	app      = cli.App("pdplay", "A minimal PureData player implemented in Go.")
	filename = app.StringOpt("name", "main.pd", "Name of the main file.")
	patchdir = app.StringArg("PATCHDIR", "", "Path to the patch dir.")
)

func main() {
	log.SetFlags(0)
	app.Action = appRun
	app.Run(os.Args)
}

func appRun() {
	defer closer.Close()

	if err := portaudio.Initialize(); paError(err) {
		log.Fatalln("PortAudio init error:", err)
	}
	closer.Bind(func() {
		if err := portaudio.Terminate(); paError(err) {
			log.Println("PortAudio term error:", err)
		}
	})

	const debug = true
	player, err := NewPlayer(*filename, *patchdir, audioChannels, sampleRate, debug)
	if err != nil {
		closer.Fatalln(err)
	}
	closer.Bind(func() {
		player.StopDSP()
		player.Close()
	})

	var stream *portaudio.Stream
	if err := portaudio.OpenDefaultStream(&stream, 0, audioChannels, sampleFormat, sampleRate,
		frameSize*ticks, player.StreamCallback, nil); paError(err) {
		log.Fatalln("PortAudio error:", err)
	}
	closer.Bind(func() {
		if err := portaudio.CloseStream(stream); paError(err) {
			log.Println("[WARN] PortAudio error:", err)
		}
	})

	player.StartDSP()
	if err := portaudio.StartStream(stream); paError(err) {
		log.Fatalln("PortAudio error:", err)
	}
	closer.Bind(func() {
		if err := portaudio.StopStream(stream); paError(err) {
			log.Fatalln("[WARN] PortAudio error:", err)
		}
	})

	closer.Hold()
}

func paError(err portaudio.Error) bool {
	return portaudio.ErrorCode(err) != portaudio.PaNoError
}

type Player struct {
	dsp *libpd.Instance
	buf []float32
}

func NewPlayer(patchName, patchDir string, channels, sampleRate int, debug bool) (*Player, error) {
	instance := libpd.NewInstance()
	if debug {
		instance.SetPrintHook(func(msg string) {
			log.Print("> ", msg)
		})
	}
	if err := instance.Init(0, channels, sampleRate); err != nil {
		err = fmt.Errorf("instance.Init failed: %v", err)
		return nil, err
	}
	patch := instance.OpenPatch(patchName, patchDir)
	if patch == nil {
		err := fmt.Errorf("failed opening patch %s", patchName)
		return nil, err
	}
	if debug {
		path := filepath.Join(patch.Dir(), patch.Name())
		log.Printf("playing %s ($0=%d)\n", path, patch.DollarZero())
	}
	player := &Player{
		dsp: instance,
		buf: make([]float32, frameSize*ticks*channels),
	}
	return player, nil
}

func (p *Player) Close() {
	if p.dsp != nil {
		p.dsp.CloseAllPatches()
		p.dsp = nil
	}
}

func (p *Player) StartDSP() {
	q, _ := p.dsp.NewMessageQueue("pd", "dsp", 1)
	q <- 1
}

func (p *Player) StopDSP() {
	q, _ := p.dsp.NewMessageQueue("pd", "dsp", 1)
	q <- 0
}

var messageFailedErr = errors.New("failed to send message")

const (
	statusContinue = int32(portaudio.PaContinue)
	statusComplete = int32(portaudio.PaComplete)
	statusAbort    = int32(portaudio.PaAbort)
)

func (p *Player) StreamCallback(_ unsafe.Pointer, output unsafe.Pointer, sampleCount uint,
	_ *portaudio.StreamCallbackTimeInfo, _ portaudio.StreamCallbackFlags, _ unsafe.Pointer) int32 {

	if !p.dsp.ProcessFloat32(ticks, nil, p.buf) {
		return statusAbort
	}
	out := (*(*[1 << 32]float32)(output))[:sampleCount*audioChannels]
	copy(out, p.buf)

	return statusContinue
}
