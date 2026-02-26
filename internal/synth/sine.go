package synth

import (
	"math"
	"sync"
)

type sineSynth struct {
	currentSampleRate float64 // Hz
	freq              float64 // Hz
	freqMu            sync.RWMutex
	pos               int
}

func NewSineSynth(sampleRate, freq float64) Synth {
	return &sineSynth{
		currentSampleRate: sampleRate,
		freq:              freq,
	}
}

func (ss *sineSynth) SetWaveFrequency(freq float64) {
	ss.freqMu.Lock()
	ss.freq = freq
	ss.freqMu.Unlock()
}

func (ss *sineSynth) SynthesizeWave(samples [][2]float64) (int, bool) {
	ss.freqMu.RLock()
	f := ss.freq
	ss.freqMu.RUnlock()

	for i := range samples {
		v := math.Sin(2 * math.Pi * float64(ss.pos) * f / ss.currentSampleRate)
		samples[i][0] = v
		samples[i][1] = v
		ss.pos++
	}

	return len(samples), true
}

func (ss *sineSynth) GetSampleRate() float64 {
	return ss.currentSampleRate
}
