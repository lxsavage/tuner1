package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func promptOpt() string {
	fmt.Print("\n> ")

	reader := bufio.NewScanner(os.Stdin)
	if !reader.Scan() {
		log.Fatalf("Reached EOF")
	}

	text := strings.TrimSpace(reader.Text())
	return text
}

func ui(tunings []Note, a4 float64) {
	//#region Audio
	// TODO: Clean up sine wave streaming/playing code here
	const sampleRate = 44100
	sr := beep.SampleRate(sampleRate)
	speaker.Init(sr, sr.N(time.Second/10))

	var (
		mu   sync.RWMutex
		freq = 0.0 // off by default
		pos  = 0
	)

	// This function generates the actual sound wave
	streamer := beep.StreamerFunc(func(samples [][2]float64) (int, bool) {
		mu.RLock()
		f := freq
		mu.RUnlock()

		for i := range samples {
			v := math.Sin(2 * math.Pi * float64(pos) * f / sampleRate)
			samples[i][0] = v
			samples[i][1] = v
			pos++
		}
		return len(samples), true
	})

	speaker.Play(streamer)
	//#endregion

	// Setup printing
	mode := 0

	for {
		// Move writing cursor back to top
		fmt.Print("\033[H\033[2J")

		fmt.Println("[ q] exit")
		if mode == 0 {
			fmt.Println("[!0] off")
		} else {
			fmt.Println("[ 0] off")
		}

		for i := range tunings {
			note := tunings[i]

			fmt.Print("[")
			if mode-1 == i {
				fmt.Print("!")
			} else {
				fmt.Print(" ")
			}

			fmt.Printf("%d] %s\n", i+1, note)
		}

		//#region Command processing
		opt := promptOpt()

		if opt == "q" {
			break
		}

		opt_int, err := strconv.Atoi(opt)
		if err != nil {
			continue
		}
		//#endregion

		if opt_int >= 0 && opt_int <= len(tunings) {
			mode = opt_int
			var newFreq float64
			if mode == 0 {
				newFreq = 0.0
			} else {
				var err error
				newFreq, err = pitchOf(tunings[mode-1], a4)
				if err != nil {
					log.Fatal(err)
				}
			}

			mu.Lock()
			freq = newFreq
			mu.Unlock()
		}
	}
}
