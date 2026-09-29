# Agent Guide — `webtyp/audio`

Constraints for agents working on this library. **Read this before any change.**
The current work order, when one exists, is [docs/PLAN.md](docs/PLAN.md). The ecosystem plan
is [`agent/docs/AGENT_ECOSYSTEM_MASTER_PLAN.md`](https://github.com/webtyp/agent/blob/main/docs/AGENT_ECOSYSTEM_MASTER_PLAN.md).

## What this library is

The audio value (`PCM`) shared by capture, speech-to-text, text-to-speech and playback. It
holds no model, no browser call and no dependency. A change that needs a browser API (the
microphone, `AudioContext`) belongs in the library that owns that API, not here.

## The builds that define "done"

```bash
go vet ./...
gotest
gotest -tinygo
GOOS=js GOARCH=wasm go build ./...
```

## Never import these

| Never | Use instead | Why |
|---|---|---|
| `fmt`, `errors`, `strings`, `strconv` | `webtyp.com/fmt` | isomorphic, small under TinyGo |
| `context` (stdlib) | `webtyp.com/context` | |
| `time` | `webtyp.com/time` | |
| `encoding/json` | nothing | reflection JSON costs ~1 MB of wasm |
| `map[K]V` | a slice | TinyGo's map runtime is a size tax |
| `os`, `log`, `net/http`, `syscall/js` | nothing | this is a value library |

## Common mistakes to avoid

- Switching `Samples` to `[]float64` or `[]int16`. Web Audio is float32, so any other type
  adds a conversion at every browser crossing.
- Adding a per-model field (such as "language"). Audio has no language, but a transcript does
  (`webtyp/stt`).
