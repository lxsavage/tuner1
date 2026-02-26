// Package synth provides audio synthesis interfaces and implementations for generating waveform signals for use with the "beep" library.
package synth

type Synth interface {
	// Sets the frequency of the wave. Works while the wave is playing.
	SetWaveFrequency(freq float64)

	// Generate a wave in place in samples. Returns the amount of samples
	// after the operation, as well as if the operation was successful.
	SynthesizeWave(samples [][2]float64) (int, bool)

	// Returns the sample rate of the synth in Hz.
	GetSampleRate() float64
}

func NewSynth(waveType string, sampleRate float64) Synth {
	var waveSynth Synth
	switch waveType {
	case "square":
		waveSynth = NewSquareSynth(sampleRate, 0)
	case "sawtooth":
		waveSynth = NewSawtoothSynth(sampleRate, 0)
	case "sine":
		fallthrough
	default:
		waveSynth = NewSineSynth(sampleRate, 0)
	}

	return waveSynth
}
