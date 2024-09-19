package voice

import (
	"bytes"
	"fmt"
	"time"

	"github.com/amitybell/piper"
	jenny "github.com/amitybell/piper-voice-jenny"
	"github.com/gopxl/beep"         // Import the beep package for playing audio
	"github.com/gopxl/beep/speaker" // Import the speaker package to output the audio
	"github.com/gopxl/beep/wav"     // Import the wav package to decode the wav audio
)

func Example(input string) {
	tts, err := piper.New("", jenny.Asset)
	if err != nil {
		fmt.Println("error 0")
		panic(err)
	}
	wavBytes, err := tts.Synthesize(input)

	if err != nil {
		fmt.Println("error 1")
		panic(err)
	}

	r := bytes.NewReader(wavBytes)

	streamer, format, err := wav.Decode(r)
	if err != nil {
		fmt.Println("error 2")
		panic(err) // If there's an error, stop the program
	}

	// Initialize the speaker with the sample rate from the WAV data
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Create a done channel to signal when the audio has finished playing
	done := make(chan bool)

	// Play the audio and signal the done channel when finished
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true // Signal the done channel
	})))

	// Wait for the audio to finish playing before allowing the program to exit
	<-done
}
