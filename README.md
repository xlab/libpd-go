libpd-go ![libpd](http://barangulesen.com/puredata/1.png)
========

The package provides Go bindings for for libpd — Pure Data embeddable audio synthesis library.<br />
All the binding code has automatically been generated with rules defined in [core.yml](/core.yml).

Before using you must install [libpd](https://github.com/libpd/libpd) library. Don't worry, it installs fine and
that's the fastest way to start unsing PD as an embedded DSP.

### Usage

```
$ go get github.com/xlab/libpd-go
```

### Demo

There is a simple Pure Data player implemented in Go that can read patches, including extras, and play them via [portaudio-go](https://github.com/xlab/portaudio-go). So you will need to get portaudio installed first.

```bash
$ brew install portaudio
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
