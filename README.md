libpd-go <img alt="puredata-logo" src="http://barangulesen.com/puredata/1.png" width="50px"/> [![GoDoc](https://godoc.org/github.com/xlab/libpd-go/libpd?status.svg)](https://godoc.org/github.com/xlab/libpd-go/libpd)
========

This project provides Go bindings for Pure Data wrapper z_libpd.h — a Pure Data embeddable audio synthesis library.<br />
All the binding code has automatically been generated with rules defined in [core.yml](/core.yml). There is also a high-level 
Go package **libpd** implemented over the core, it introduces threadsafe access to PD, some idiomatic helpers and allows to run multiple instances of PD.

Before start you must install [libpd](https://github.com/libpd/libpd) library. Don't worry, it installs fine and
that's the fastest way to begin using PD as an embedded DSP.

### Usage

```
$ go get github.com/xlab/libpd-go/libpd
```

### Demo

There is a minimal Pure Data player implemented in Go that can read patches, including extras, and play them via [portaudio-go](https://github.com/xlab/portaudio-go). It's about 100 lines of code. You will need to get [PortAudio](http://www.portaudio.com) installed first.

```bash
$ brew install portaudio
$ go get github.com/xlab/libpd-go/cmd/pdplay

$ pdplay -h

Usage: pdplay [OPTIONS] PATCHDIR

A minimal PureData player implemented in Go.

Arguments:
  PATCHDIR=""   Path to the patch dir.

Options:
  --name="main.pd"   Name of the main file.

$ pdplay $GOPATH/src/github.com/xlab/libpd-go/assets/patch01

# an atmospheric meditation synthesis sound goes...
# by freezemode - http://soundcloud.com/freezemode
```

<img src="https://cl.ly/322o0A1W0s10/pdscr.png" width="500"/>

Try this patch for example too: [PerotinusRandom on pdpatchrepo.info](http://pdpatchrepo.info/patches/patch/6) (also a precompiled version with **freeverb~** extra for OS X 64-bit located here: [PerotinusRandom.zip](http://dl.xlab.is/music/pd/PerotinusRandom.zip)). Awesome chorus!

### Rebuilding the package

You will need to get the [cgogen](https://git.io/cgogen) tool installed first.

```
$ git clone https://github.com/xlab/libpd-go && cd libpd-go
$ make clean
$ make
```
