package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jbuchbinder/beep"
	"github.com/jbuchbinder/beep/generators"
	"github.com/jbuchbinder/beep/speaker"
)

var (
	list = flag.Bool("list", false, "List all supported stations")
)

func usage() {
	fmt.Printf("usage: %s [-list] station [...station]\n", os.Args[0])
	fmt.Println("station is the numeric station identifier")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if *list {
		listStations()
		return
	}

	if len(args) < 1 {
		usage()
		return
	}

	for _, a := range args {
		if len(a) == 1 {
			a = "0" + a
		}
		if t, ok := toneMap[a]; ok {
			fmt.Printf(" -> %s\n", t.name)
			err := playTones(t.tone1, t.tone2)
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

func listStations() {
	fmt.Printf("Supported station IDs and tones:\n\n")
	for k, v := range toneMap {
		fmt.Printf("%3s : %35s [%0.1f,%0.1f]\n", k, v.name, v.tone1, v.tone2)
	}
}
