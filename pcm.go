// Package audio holds the audio value shared by capture, speech-to-text, text-to-speech and
// playback.
package audio

// PCM is uncompressed audio: interleaved float32 samples in [-1, 1].
type PCM struct {
	SampleRate int       // samples per second per channel, e.g. 16000
	Channels   int       // 1 = mono, 2 = stereo; Samples holds channels interleaved
	Samples    []float32 // len(Samples) == frames × Channels
}

// Frames is the number of sample instants: len(Samples) / Channels. It is 0 when Channels is 0.
func (p PCM) Frames() int {
	if p.Channels <= 0 {
		return 0
	}
	return len(p.Samples) / p.Channels
}

// DurationMS is the length of the audio in milliseconds, 0 when SampleRate or Channels is 0.
func (p PCM) DurationMS() int64 {
	if p.SampleRate <= 0 {
		return 0
	}
	frames := p.Frames()
	if frames == 0 {
		return 0
	}
	return int64(frames) * 1000 / int64(p.SampleRate)
}
