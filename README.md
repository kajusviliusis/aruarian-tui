# aruarian-tui

productivity app that handles notes, todo lists, work timer

## Install
```bash
go install github.com/kajusviliusis/aruarian-tui/cmd/aruarian@latest
````

## Build

from the repository root:

```bash
go build -o aruarian ./cmd/app
./aruarian
```

move the binary to path for easier access:

```bash
sudo mv aruarian /usr/local/bin
```



## Layout

- `cmd/app` - application entrypoint
- `internal/app` - root application model and state handling
- `internal/menu` - menu UI
- `internal/notes` - notes launching helpers
- `internal/styles` - shared styling
- `internal/timer` - timer model
- `internal/todo` - todo model and persistence

## Other
- todos are stored in `~/.config/aruarian-tui/`
- notes are automatically launched in `~/notes`, which is created if it doesn't exist

