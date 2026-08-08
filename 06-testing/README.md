# 06 — Testing

> 🚧 **Status:** Syllabus ready — 9 lessons outlined, full content in progress.

Go treats testing as a first-class part of the toolchain (`go test`), not a bolted-on framework. This module covers the idioms — table-driven tests especially — that make Go test suites fast to write and fast to run.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Testing Package Basics](./01-testing-package-basics.md) | The `testing` package, `*testing.T`, and Go's `TestXxx` naming convention. |
| 02 | [Table Driven Tests](./02-table-driven-tests.md) | The idiomatic table-driven test pattern for covering many cases with one test function. |
| 03 | [Subtests and Testmain](./03-subtests-and-testmain.md) | Organizing related cases with `t.Run` subtests, and shared setup/teardown via `TestMain`. |
| 04 | [Mocking with Interfaces](./04-mocking-with-interfaces.md) | Designing code for testability by depending on small interfaces instead of concrete types. |
| 05 | [Testify and Gomock](./05-testify-and-gomock.md) | Reducing boilerplate with the `testify` assertion library and generating mocks with `gomock`. |
| 06 | [Httptest for HTTP Handlers](./06-httptest-for-http-handlers.md) | Testing HTTP handlers in isolation with `net/http/httptest`, no real network required. |
| 07 | [Golden Files and Test Fixtures](./07-golden-files-and-test-fixtures.md) | The golden-file pattern for comparing large or complex expected output. |
| 08 | [Code Coverage and Coverprofile](./08-code-coverage-and-coverprofile.md) | Generating and reading coverage reports with `go test -cover` and `-coverprofile`. |
| 09 | [Fuzzing Basics](./09-fuzzing-basics.md) | Go's built-in fuzz testing (`go test -fuzz`) for discovering edge cases automatically. |

**Previous module:** [05 — Standard Library](../05-standard-library/README.md) · **Next module:** [07 — Tooling & Code Quality](../07-tooling-and-code-quality/README.md)
