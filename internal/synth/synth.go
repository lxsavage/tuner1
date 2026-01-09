// Package synth provides audio synthesis interfaces and implementations for generating waveform signals for use with the "beep" library.
package synth

type Synth interface {
	// Sets the frequency of the wave. Works while the wave is playing.
	SetWaveFrequency(new_freq float64)

	// Generate a wave in place in samples. Returns the amount of samples
	// after the operation, as well as if the operation was successful.
	SynthesizeWave(samples [][2]float64) (int, bool)

	// Returns the sample rate of the synth in Hz.
	GetSampleRate() float64
}

func NewSynth(wave_type string, sample_rate float64) Synth {
	var wave_synth Synth
	switch wave_type {
	case "square":
		wave_synth = NewSquareSynth(sample_rate, 0)
	case "sawtooth":
		wave_synth = NewSawtoothSynth(sample_rate, 0)
	case "sine":
		fallthrough
	default:
		wave_synth = NewSineSynth(sample_rate, 0)
	}

	return wave_synth
}
