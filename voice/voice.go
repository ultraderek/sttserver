package voice

import (
	"github.com/amitybell/piper"
	jenny "github.com/amitybell/piper-voice-jenny"
)

func Example() {
	tts, err := piper.New("", jenny.Asset)
	if err != nil {
		panic(err)
	}
	wav, err := tts.Synthesize("hello world")
	if err != nil {
		panic(err)
	}
	_ = wav
}
