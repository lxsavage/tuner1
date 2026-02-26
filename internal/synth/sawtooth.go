package synth

import (
	"sync"
)

type sawtoothSynth struct {
	currentSampleRate float64 // Hz
	freq              float64 // Hz
	freqMu            sync.RWMutex
	phase             float64
}

func NewSawtoothSynth(sampleRate, freq float64) Synth {
	return &sawtoothSynth{
		currentSampleRate: sampleRate,
		freq:              freq,
	}
}

func (ss *sawtoothSynth) SetWaveFrequency(freq float64) {
	ss.freqMu.Lock()
	ss.freq = freq
	ss.freqMu.Unlock()
}

func (ss *sawtoothSynth) SynthesizeWave(samples [][2]float64) (int, bool) {
	ss.freqMu.RLock()
	freq := ss.freq
	ss.freqMu.RUnlock()

	for i := range samples {
		// Sawtooth wave: linear ramp from -1 to 1 over one period
		v := 2.0*ss.phase - 1.0

		samples[i][0] = v
		samples[i][1] = v

		ss.phase += freq / ss.currentSampleRate
		if ss.phase >= 1 {
			ss.phase -= 1
		}
	}

	return len(samples), true
}

func (ss *sawtoothSynth) GetSampleRate() float64 {
	return ss.currentSampleRate
}
