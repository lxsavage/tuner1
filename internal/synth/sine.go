package synth

import (
	"math"
	"sync"
)

type sine_synth struct {
	current_sample_rate float64 // Hz
	freq                float64 // Hz
	mu_freq             sync.RWMutex
	phase               float64
}

func NewSineSynth(sample_rate, freq float64) Synth {
	return &sine_synth{
		current_sample_rate: sample_rate,
		freq:                freq,
	}
}

func (ss *sine_synth) SetWaveFrequency(new_freq float64) {
	ss.mu_freq.Lock()
	ss.freq = new_freq
	ss.mu_freq.Unlock()
}

func (ss *sine_synth) SynthesizeWave(samples [][2]float64) (int, bool) {
	ss.mu_freq.RLock()
	f := ss.freq
	ss.mu_freq.RUnlock()

	for i := range samples {
		v := math.Sin(2 * math.Pi * ss.phase)
		samples[i][0] = v
		samples[i][1] = v

		ss.phase += f / ss.current_sample_rate
		if ss.phase >= 1 {
			ss.phase -= 1
		}
	}

	return len(samples), true
}

func (ss *sine_synth) GetSampleRate() float64 {
	return ss.current_sample_rate
}
