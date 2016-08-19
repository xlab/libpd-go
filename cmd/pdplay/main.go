package main

import (
	"log"
	"os"

	"github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	"github.com/xlab/portaudio-go/portaudio"
)

const (
	frameSize     = 64
	bitDepth      = 16
	audioChannels = 2
	sampleRate    = 44100
	sampleFormat  = portaudio.PaFloat32
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
	pdPlayer := InitPureDataPlayer(*filename, *patchdir,
		audioChannels, sampleRate, frameSize, debug)
	if pdPlayer == nil {
		closer.Close()
	}
	closer.Bind(func() {
		pdPlayer.StopDSP()
		pdPlayer.Close()
	})

	var stream *portaudio.Stream
	if err := portaudio.OpenDefaultStream(&stream, 0, audioChannels, sampleFormat, sampleRate,
		frameSize, pdPlayer.StreamCallback, nil); paError(err) {
		log.Fatalln("PortAudio error:", err)
	}
	closer.Bind(func() {
		if err := portaudio.CloseStream(stream); paError(err) {
			log.Println("[WARN] PortAudio error:", err)
		}
	})

	pdPlayer.StartDSP()
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
