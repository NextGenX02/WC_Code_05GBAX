package utils

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"log"
	"os"
	"path"
	"time"
)

// DoMusic Shiku favorite features!
func DoMusic() {
	// get Current Path
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	_, fileINfoError := os.Stat(path.Join(wd, "shiku.wav"))
	if os.IsNotExist(fileINfoError) {
		return
	}
	// hemm...~, good song!
	file, _ := os.OpenFile(path.Join(wd, "shiku.wav"), os.O_RDONLY, os.ModePerm)

	// Decode the WAV file
	streamer, format, decodeError := wav.Decode(file)
	if decodeError != nil {
		log.Fatalln(decodeError)
	}
	defer streamer.Close()
	// Create Bool channel
	isDone := make(chan bool)

	speakerInitError := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if speakerInitError != nil {
		log.Fatalln(speakerInitError)
	}
	loop := beep.Loop(-1, streamer)

	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		isDone <- true
	})))

	// Since is running at separate thread we use the channel to tell if the audio is done playing
	<-isDone
}
