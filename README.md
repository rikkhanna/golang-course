# Go Curriculum

A hands-on, comprehensive Go course. One Go module, one section per topic.

## How to run things

From the repo root:

```bash
go run ./01-foundations/01-hello   # run a single example
go build ./...                      # compile everything
go test ./...                       # run all exercise tests
go vet ./...                        # static checks
go test -race ./...                 # data-race detector (concurrency sections)
gofmt -l .                          # list unformatted files (want: empty)
```

## Sections

| #  | Topic | Status |
|----|-------|--------|
| 01 | Foundations — toolchain, variables, types, control flow, functions | ✅ done |
| 02 | Core data structures — arrays, slices, maps, strings, pointers | ✅ done |
| 03 | Error handling — errors, wrapping, panic/recover | ✅ done |
| 04 | Structs, methods, interfaces | pending |
| 05 | Concurrency — goroutines, channels, select, sync, race detector | pending |
| 06 | Generics — type parameters, constraints | pending |
| 07 | Packages / modules | pending |
| 08 | Testing & tooling | pending |
| 09 | Standard library — net/http, encoding/json, io, database/sql | pending |
| 10 | Advanced — reflection, unsafe, cgo, build tags, memory model | pending |
| 11 | Ecosystem — idioms, cobra/gin/zap | pending |
| —  | Capstone project | pending |

## Layout convention

- Each **concept** is its own directory with a `main.go` (so `go run ./NN-x/concept` works).
- Each **exercise** lives under `NN-x/exercises/` as a stub file + a `_test.go` you make pass.
