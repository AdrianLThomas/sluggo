package assets

import (
	"bytes"
	"io"
	"sync"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const sampleRate = 44100

var (
	audioCtx     *audio.Context
	audioOnce    sync.Once
	musicPlayer  *audio.Player
	musicStarted bool
)

func context() *audio.Context {
	audioOnce.Do(func() {
		audioCtx = audio.NewContext(sampleRate)
	})
	return audioCtx
}

func decodeOgg(name string) (*vorbis.Stream, error) {
	f, err := assets.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return vorbis.Decode(context(), bytes.NewReader(data))
}

// PlaySFX plays a one-shot sound effect by file name, e.g. PlaySFX("sfx_eat.ogg").
func PlaySFX(name string) {
	stream, err := decodeOgg(name)
	if err != nil {
		return
	}

	p, err := context().NewPlayer(stream)
	if err != nil {
		return
	}

	p.Play()
}

// PlayMusic starts the background music loop. It is idempotent: subsequent
// calls are no-ops.
func PlayMusic() {
	if musicStarted {
		return
	}

	stream, err := decodeOgg("music.ogg")
	if err != nil {
		return
	}

	loop := audio.NewInfiniteLoop(stream, stream.Length())
	p, err := context().NewPlayer(loop)
	if err != nil {
		return
	}

	p.SetVolume(0.6)
	p.Play()
	musicPlayer = p
	musicStarted = true
}
