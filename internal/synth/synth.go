package synth

type Synth interface {
	// Sets the frequency of the wave. Works while the wave is playing.
	SetWaveFrequency(new_freq float64)

	// Generate a wave in place in samples. Returns the amount of samples
	// after the operation, as well as if the operation was successful.
	SynthesizeWave(samples [][2]float64) (int, bool)
}
