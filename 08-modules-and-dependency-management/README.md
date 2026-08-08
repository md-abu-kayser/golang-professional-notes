# 08 — Modules & Dependency Management

> 🚧 **Status:** Syllabus ready — 5 lessons outlined, full content in progress.

A deeper dive into the module system introduced in [01-getting-started/05](../01-getting-started/05-modules-and-packages-intro.md): versioning, vendoring, local overrides, private modules, and multi-module workspaces.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Go Mod Init Tidy Vendor](./01-go-mod-init-tidy-vendor.md) | The full `go.mod` lifecycle, including vendoring dependencies for reproducible, offline builds. |
| 02 | [Semantic Versioning and Module Proxy](./02-semantic-versioning-and-module-proxy.md) | How semantic versioning rules interact with the module proxy's version resolution. |
| 03 | [Replace and Exclude Directives](./03-replace-and-exclude-directives.md) | Using `replace` for local development overrides and `exclude` to block bad versions. |
| 04 | [Private Modules and Authentication](./04-private-modules-and-authentication.md) | Configuring `GOPRIVATE` and authentication for private module hosts. |
| 05 | [Workspaces Multi Module Development](./05-workspaces-multi-module-development.md) | Developing across multiple local modules simultaneously with `go.work`. |

**Previous module:** [07 — Tooling & Code Quality](../07-tooling-and-code-quality/README.md) · **Next module:** [09 — Web Development](../09-web-development/README.md)
