# audio
<img src="docs/img/badges.svg">

The audio value shared by the webtyp voice pieces. `audio.PCM` is a piece of uncompressed
sound: float32 samples in [-1, 1], with a sample rate and a channel count. That is the shape
the browser's Web Audio API uses, so no conversion is needed at the edge.

> Voice ships in version 2 of the agent; version 1 is text only. This type is declared now so
> the voice pipeline is wired from the start.

## Getting started

```go
speech := audio.PCM{SampleRate: 16000, Channels: 1, Samples: samples}
ms := speech.DurationMS()
```

| I want to… | Use |
|---|---|
| pass audio from the microphone to speech-to-text | `audio.PCM` (`webtyp/media` → `webtyp/stt`) |
| play synthesized speech | `audio.PCM` (`webtyp/tts` → playback) |
| know how long a clip is | `PCM.DurationMS()` |

## Documentation

- [Architecture](docs/ARCHITECTURE.md): where this type sits in the voice pipeline, and what version 2 adds.
- [Agent guide](AGENTS.md): rules for anyone changing this library.
