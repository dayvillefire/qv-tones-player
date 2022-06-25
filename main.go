package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jbuchbinder/beep"
	"github.com/jbuchbinder/beep/generators"
	"github.com/jbuchbinder/beep/speaker"
)

func usage() {
	fmt.Printf("usage: %s station [...station]\n", os.Args[0])
	fmt.Println("station is the numeric station identifier")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	for _, a := range os.Args[1:] {
		if len(a) == 1 {
			a = "0" + a
		}
		if t, ok := toneMap[a]; ok {
			err := playTones(t[0], t[1])
			if err != nil {
				panic(err)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func playTones(f1, f2 float64) error {
	sr := beep.SampleRate(48000)
	speaker.Init(sr, 4800)

	sine1, err := generators.SineTone(sr, f1)
	if err != nil {
		return err
	}
	sine2, err := generators.SineTone(sr, f2)
	if err != nil {
		return err
	}

	// Play 2 seconds of each tone
	two := sr.N(2 * time.Second)

	ch := make(chan struct{})
	sounds := []beep.Streamer{
		//beep.Callback(print("sine")),
		beep.Take(two, sine1),
		beep.Take(two, sine2),
		beep.Callback(func() {
			ch <- struct{}{}
		}),
	}
	speaker.Play(beep.Seq(sounds...))

	<-ch

	return nil
}

// a simple clousure to wrap fmt.Println
func print(s string) func() {
	return func() {
		fmt.Println(s)
	}
}
