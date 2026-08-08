# 05 — Standard Library

> 🚧 **Status:** Syllabus ready — 10 lessons outlined, full content in progress.

Go ships with an unusually capable standard library — this module tours the packages you'll reach for daily, before pulling in any third-party dependency.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Fmt and Io](./01-fmt-and-io.md) | Formatting verbs in `fmt`, and the `io.Reader`/`io.Writer` interfaces that unify I/O across the standard library. |
| 02 | [Strings Strconv and Bytes](./02-strings-strconv-and-bytes.md) | String manipulation with `strings`, type conversions with `strconv`, and working with raw `[]byte`. |
| 03 | [Os and File Handling](./03-os-and-file-handling.md) | Reading and writing files, command-line arguments, and environment variables via `os`. |
| 04 | [Log and Logrus Zap](./04-log-and-logrus-zap.md) | The standard `log` package, and structured logging with `logrus` and `zap` for production services. |
| 05 | [Encoding JSON and XML](./05-encoding-json-and-xml.md) | Struct tags, `Marshal`/`Unmarshal`, and writing custom (un)marshalers. |
| 06 | [Net HTTP Client and Server](./06-net-http-client-and-server.md) | Building HTTP servers and making outbound HTTP calls with `net/http` alone. |
| 07 | [HTML Template](./07-html-template.md) | Safe, auto-escaping server-rendered HTML with `html/template`. |
| 08 | [Sort and Container Packages](./08-sort-and-container-packages.md) | Custom sorting with `sort.Slice`, plus `container/heap` and `container/list`. |
| 09 | [Math Rand and Crypto Rand](./09-math-rand-and-crypto-rand.md) | The difference between `math/rand` (fast, predictable) and `crypto/rand` (secure) — and when each is appropriate. |
| 10 | [Testing and Benchmarking](./10-testing-and-benchmarking.md) | A standard-library-level preview of `testing` — the full deep dive lives in [06-testing](../06-testing/README.md). |

**Previous module:** [04 — Concurrency](../04-concurrency/README.md) · **Next module:** [06 — Testing](../06-testing/README.md)
