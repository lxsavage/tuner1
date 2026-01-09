package synth

import (
	"sync"
)

type square_synth struct {
	current_sample_rate float64 // Hz
	freq                float64 //Hz
	mu_freq             sync.RWMutex
	phase               float64
}

func NewSquareSynth(sample_rate, freq float64) Synth {
	return &square_synth{
		current_sample_rate: sample_rate,
		freq:                freq,
	}
}

func (ss *square_synth) SetWaveFrequency(new_freq float64) {
	ss.mu_freq.Lock()
	ss.freq = new_freq
	ss.mu_freq.Unlock()
}

func (ss *square_synth) SynthesizeWave(samples [][2]float64) (int, bool) {
	ss.mu_freq.RLock()
	freq := ss.freq
	ss.mu_freq.RUnlock()

	for i := range samples {
		v := 1.
		if ss.phase < .5 {
			v = -1.
		}

		samples[i][0] = float64(v)
		samples[i][1] = float64(v)

		ss.phase += freq / ss.current_sample_rate
		if ss.phase >= 1 {
			ss.phase -= 1
		}
	}

	return len(samples), true
}

func (ss *square_synth) GetSampleRate() float64 {
	return ss.current_sample_rate
}
