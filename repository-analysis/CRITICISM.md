# Review findings

## Remaining Issues

- **File reopening can corrupt existing logs.** `logger/file.go:OpenLoggerFile` opens with `O_WRONLY|O_CREATE`, without append or truncation. Reusing a filename (including within the default one-second naming interval) starts writing at offset zero and can leave stale trailing bytes. Use exclusive creation with collision handling, or explicit append semantics. Reinitialization also replaces the file handle without closing it.
- **Output failures are silent.** `writeLogJSONL` drops marshal failures and ignores write errors; terminal writes also ignore errors. Unsupported nested values, cycles, and disk failures can lose records without feedback. Return errors or provide an error callback. `InitializeConfig` instead exits the process on file-open failure, which limits library integration.
- **Formatting changes argument types.** `LogBool` and the reader convert every argument to a string before `fmt.Sprintf`; `%d` and other typed verbs produce formatting diagnostics. Document the restricted contract or preserve argument types in formatting. File sanitization is shallow and does not strip preexisting ANSI sequences, so claims of universally ANSI-free storage would be too strong.
- **Default filling prevents explicit zero values.** `logger/default-value.go:IsZeroOrEmpty` treats `Critical` (0) as missing and recursively treats a pointer to false as empty. The latter can prevent disabling TID after it has been enabled because initialization uses current `Cfg` as defaults. Use explicit optional fields or distinguish unset values from valid zero values.
- **Reader limits and misleading options.** `cmd/log-reader/main.go` buffers every line even without tail, uses the scanner's default token limit, and parses empty time bounds despite help text allowing them. Large JSONL records can be written but fail to replay. Stream ordinary reads, use a bounded tail buffer, support optional bounds, and choose an explicit record-size policy. Time-format overrides also leave the fixed default bounds unchanged, requiring users to supply compatible bounds.
- **Build configuration diverges.** `build/Dockerfile` uses Go 1.25.3 while `go.mod` requires 1.25.5. Builds can require automatic toolchain downloading. Align the versions for predictable builds.
- **Verification and concurrent configuration are incomplete.** No automated tests or CI files were found. Output mutexes do not protect mutable `Cfg`, file-handle replacement, or the colorizer registry. Configuration during logging can race; document initialization-only use or synchronize state. Prioritize tests for file collisions, formatting, default values, and reader boundaries.

## Simplification Opportunities

- Consolidate the overlapping multiline ANSI helpers in `palette/ansi-format.go` and `palette/colorizer.go`; they implement similar wrapping with different newline handling. One internal implementation would reduce behavior drift while retaining the public API.
