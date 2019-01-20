package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/acmujica/ding/audio"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

//go:generate go-bindata -pkg audio -o audio/audio.go audio/

func main() {
	// Default measurement is seconds parse for alternative
	isHours := flag.Bool("h", false, "Sets the measurement to seconds")
	isMinutes := flag.Bool("m", false, "Sets the measurement to minutes")
	wantsLouder := flag.Bool("l", false, "Adds a louder alarm")
	flag.Parse()
	numberFlt, err := strconv.ParseFloat(os.Args[len(os.Args)-1], 64)
	if err != nil {
		fmt.Println("Could not read value for timer.")
	}

	// Create wav file
	data, err := audio.Asset("audio/hand-bell.wav")
	if err != nil {
		fmt.Println("Could not open bell file.")
	}

	if *wantsLouder {
		data, err = audio.Asset("audio/alarm.wav")
		if err != nil {
			fmt.Println("Could not open bell file.")
		}
	}

	f, _ := ioutil.TempFile("", "audio")
	defer f.Close()
	filePath := f.Name()
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	// Run Timer
	var waitTime time.Duration
	switch {
	case *isHours:
		waitTime = time.Duration(numberFlt) * time.Hour
	case *isMinutes:
		waitTime = time.Duration(numberFlt) * time.Minute
	default:
		waitTime = time.Duration(numberFlt) * time.Second
	}

	timer := time.NewTimer(waitTime)
	<-timer.C

	// Play wav file
	f, _ = os.Open(filePath)
	s, format, _ := wav.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	<-done

	// Cleanup
	os.Remove(filePath)
}
