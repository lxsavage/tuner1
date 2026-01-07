package synth

import (
	"sync"
)

type sawtooth_synth struct {
	current_sample_rate float64 // Hz
	freq                float64 // Hz
	mu_freq             sync.RWMutex
	phase               float64
}

func NewSawtoothSynth(sample_rate float64, freq float64) Synth {
	return &sawtooth_synth{
		current_sample_rate: sample_rate,
		freq:                freq,
	}
}

func (ss *sawtooth_synth) SetWaveFrequency(new_freq float64) {
	ss.mu_freq.Lock()
	ss.freq = new_freq
	ss.mu_freq.Unlock()
}

func (ss *sawtooth_synth) SynthesizeWave(samples [][2]float64) (int, bool) {
	ss.mu_freq.RLock()
	freq := ss.freq
	ss.mu_freq.RUnlock()

	for i := range samples {
		// Sawtooth wave: linear ramp from -1 to 1 over one period
		v := 2.0*ss.phase - 1.0

		samples[i][0] = v
		samples[i][1] = v

		ss.phase += freq / ss.current_sample_rate
		if ss.phase >= 1 {
			ss.phase -= 1
		}
	}

	return len(samples), true
}

func (ss *sawtooth_synth) GetSampleRate() float64 {
	return ss.current_sample_rate
}
