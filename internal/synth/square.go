package synth

import (
	"sync"
)

type squareSynth struct {
	currentSampleRate float64 // Hz
	freq              float64 //Hz
	freqMu            sync.RWMutex
	phase             float64
}

func NewSquareSynth(sample_rate, freq float64) Synth {
	return &squareSynth{
		currentSampleRate: sample_rate,
		freq:              freq,
	}
}

func (ss *squareSynth) SetWaveFrequency(freq float64) {
	ss.freqMu.Lock()
	ss.freq = freq
	ss.freqMu.Unlock()
}

func (ss *squareSynth) SynthesizeWave(samples [][2]float64) (int, bool) {
	ss.freqMu.RLock()
	freq := ss.freq
	ss.freqMu.RUnlock()

	for i := range samples {
		v := 1.
		if ss.phase < .5 {
			v = -1.
		}

		samples[i][0] = float64(v)
		samples[i][1] = float64(v)

		ss.phase += freq / ss.currentSampleRate
		if ss.phase >= 1 {
			ss.phase -= 1
		}
	}

	return len(samples), true
}

func (ss *squareSynth) GetSampleRate() float64 {
	return ss.currentSampleRate
}
