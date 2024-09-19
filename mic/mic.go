package mic

import (
	"fmt"
	"os"

	"github.com/gen2brain/malgo"
)

func Example1() {
	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, func(message string) {
		fmt.Printf("LOG <%v>\n", message)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = ctx.Uninit()
		ctx.Free()
	}()

	deviceConfig := malgo.DefaultDeviceConfig(malgo.Duplex)
	deviceConfig.Capture.Format = malgo.FormatS16
	deviceConfig.Capture.Channels = 1
	deviceConfig.Playback.Format = malgo.FormatS16
	deviceConfig.Playback.Channels = 1
	deviceConfig.SampleRate = 44100
	deviceConfig.Alsa.NoMMap = 1

	//var playbackSampleCount uint32
	//var capturedSampleCount uint32

	recvcounter := 0
	sendcounter := 0
	sizeInBytes := uint32(malgo.SampleSizeInBytes(deviceConfig.Capture.Format))
	buffersize := 10

	//	buffer1 := make(chan []byte, 1)
	buffer := make([][]byte, buffersize)

	for i := range buffer {
		buffer[i] = make([]byte, 1102*int(sizeInBytes))
	}

	//pCapturedSamples := make([]byte, 22040)
	onRecvFrames := func(pSample2, pSample []byte, framecount uint32) {
		copy(buffer[recvcounter%buffersize], pSample)
		recvcounter++
		//buffer1 <- pSample

	}

	fmt.Println("Recording...")
	captureCallbacks := malgo.DeviceCallbacks{
		Data: onRecvFrames,
	}
	device, err := malgo.InitDevice(ctx.Context, deviceConfig, captureCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	onSendFrames := func(pSample, pSample2 []byte, framecount uint32) {
		//	copy(pSample, <-buffer1)

		if recvcounter > sendcounter {
			copy(pSample, buffer[sendcounter%buffersize])
			sendcounter++

		} else {
			sendcounter = 0
		}

	}

	fmt.Println("Playing...")
	playbackCallbacks := malgo.DeviceCallbacks{
		Data: onSendFrames,
	}

	device2, err := malgo.InitDevice(ctx.Context, deviceConfig, playbackCallbacks)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = device2.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Press Enter to quit...")
	fmt.Scanln()

	device.Uninit()
	device2.Uninit()

}
