---
PLAN: "feat: audio.PCM — the audio value shared by media, stt and tts"
TAG: v0.1.0
EXECUTOR: jules
REVIEWER: none
STATUS: running
SESSION: 11126279561605079163
---

> This plan is dispatched via the CodeJob workflow. See skill: agents-workflow.
>
> **Phase 4a** of
> [`AGENT_ECOSYSTEM_MASTER_PLAN.md`](https://github.com/webtyp/agent/blob/main/docs/AGENT_ECOSYSTEM_MASTER_PLAN.md).
> Independent of phases 1–3. `webtyp/stt` and `webtyp/tts` (phase 4b) wait for this tag.

# Plan — `webtyp.com/audio`: one type for a piece of sound

## 0. Context

Voice is planned for version 2 of the webtyp agent: speech-to-text (`webtyp/stt`) turns the
microphone into text, and text-to-speech (`webtyp/tts`) turns the answer into sound. The
microphone itself is `webtyp/media`. The three need one shared answer to "what does a piece
of audio look like in Go". Without it, each would declare its own, and every hand-off would
need a conversion.

This plan creates only that value type. Resampling, voice-activity detection and playback are
version 2 work and get their own plans.

**PCM** (pulse-code modulation) is uncompressed audio: a list of numbers, each the amplitude
of the sound at one instant. The Web Audio API delivers and plays microphone audio as 32-bit
floats between -1 and 1, so that is the representation used here, and no conversion is
needed at the browser edge.

## Development rules (inline)

- **Contract library:** value types only in this plan. No I/O, no browser calls.
- Every file compiles under `GOOS=js GOARCH=wasm` and TinyGo.
- **Never import:** `fmt`/`errors`/`strings`/`strconv` (use `webtyp.com/fmt`), `context`
  (use `webtyp.com/context`), `encoding/json`, `time` (use `webtyp.com/time`), `map[K]V`,
  `os`, `log`, `net/http`.
- Flat layout, max 500 lines per file, tests with `testing` only. Do **not** run `gopush`/`codejob`.

## Design gate (api-design — five answers)

1. **Prior art.** The **Web Audio API** `AudioBuffer` has a `sampleRate`, a
   `numberOfChannels`, and `Float32Array` data in [-1, 1]. **go-audio `audio.FloatBuffer`**
   has a `Format{SampleRate, NumChannels}` plus `[]float64` data. **PortAudio** and **miniaudio**
   use interleaved `float32` frames with a sample rate and a channel count. All three agree on
   the same three facts. This type uses them with the browser's own sample type (`float32`), so
   crossing into Web Audio costs no conversion.
2. **Novice-name test.** `audio.PCM{SampleRate: 16000, Channels: 1, Samples: s}` reads as "PCM
   audio, 16 kHz, mono". `PCM` is the standard term every audio API uses.
3. **Complexity ledger.** Concepts +1 / −0. Files +1 / −0. Ways to do the same thing +0: there
   is no other audio type in the ecosystem today.
4. **Where it belongs.** It is the value that crosses between `media` (capture), `stt`
   (consumes), `tts` (produces) and playback. It belongs to none of them, so it gets its own
   repository, like `llm` between the agent and a model.
5. **What it deletes.** Nothing. This is new capability.

## Stage 1 — the type

**`pcm.go`**

```go
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
func (p PCM) Frames() int

// DurationMS is the length of the audio in milliseconds, 0 when SampleRate or Channels is 0.
func (p PCM) DurationMS() int64
```

`DurationMS` returns `int64(p.Frames()) * 1000 / int64(p.SampleRate)`. It exists because both
`stt` (timestamps) and `tts` (queueing playback) need it. Without it, each would write the
same arithmetic.

## Stage 2 — tests (package `audio_test`)

**File:** `pcm_test.go`

| Test | Asserts |
|---|---|
| `TestFrames_Mono` | 16000 samples, 1 channel → 16000 |
| `TestFrames_StereoInterleaved` | 8 samples, 2 channels → 4 |
| `TestFrames_ZeroChannels` | → 0, no panic |
| `TestDurationMS_OneSecond` | 16000 frames at 16000 Hz → 1000 |
| `TestDurationMS_ZeroRate` | → 0, no panic |

## Stages

| Stage | Files | Acceptance |
|---|---|---|
| 1 | `pcm.go`, `go.mod` | builds for host, wasm, TinyGo; `go.mod` has no requirements |
| 2 | `pcm_test.go` | passes under `gotest` and `gotest -tinygo` |
