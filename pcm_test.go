package audio_test

import (
	"testing"

	"webtyp.com/audio"
)

func TestFrames_Mono(t *testing.T) {
	p := audio.PCM{
		SampleRate: 16000,
		Channels:   1,
		Samples:    make([]float32, 16000),
	}
	if got := p.Frames(); got != 16000 {
		t.Errorf("Frames() = %d, want 16000", got)
	}
}

func TestFrames_StereoInterleaved(t *testing.T) {
	p := audio.PCM{
		SampleRate: 16000,
		Channels:   2,
		Samples:    make([]float32, 8),
	}
	if got := p.Frames(); got != 4 {
		t.Errorf("Frames() = %d, want 4", got)
	}
}

func TestFrames_ZeroChannels(t *testing.T) {
	p := audio.PCM{
		SampleRate: 16000,
		Channels:   0,
		Samples:    make([]float32, 8),
	}
	if got := p.Frames(); got != 0 {
		t.Errorf("Frames() = %d, want 0", got)
	}
}

func TestDurationMS_OneSecond(t *testing.T) {
	p := audio.PCM{
		SampleRate: 16000,
		Channels:   1,
		Samples:    make([]float32, 16000),
	}
	if got := p.DurationMS(); got != 1000 {
		t.Errorf("DurationMS() = %d, want 1000", got)
	}
}

func TestDurationMS_ZeroRate(t *testing.T) {
	p := audio.PCM{
		SampleRate: 0,
		Channels:   1,
		Samples:    make([]float32, 16000),
	}
	if got := p.DurationMS(); got != 0 {
		t.Errorf("DurationMS() = %d, want 0", got)
	}
}

func TestDurationMS_ZeroChannels(t *testing.T) {
	p := audio.PCM{
		SampleRate: 16000,
		Channels:   0,
		Samples:    make([]float32, 16000),
	}
	if got := p.DurationMS(); got != 0 {
		t.Errorf("DurationMS() = %d, want 0", got)
	}
}
