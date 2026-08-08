# 09 — Web Development

> 🚧 **Status:** Syllabus ready — 9 lessons outlined, full content in progress.

Building HTTP services in Go, from the standard library alone through to popular frameworks, plus the production concerns (middleware, websockets, graceful shutdown) that separate a demo from a real service.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Net HTTP Server Basics](./01-net-http-server-basics.md) | Standing up an HTTP server using only `net/http` — no framework required. |
| 02 | [Routing and Handlers](./02-routing-and-handlers.md) | `ServeMux` routing patterns, including Go 1.22+'s built-in path parameters. |
| 03 | [Middleware Pattern](./03-middleware-pattern.md) | Composing cross-cutting concerns as middleware — see the working [middleware-chain.go](../18-code-snippets/middleware-chain.go) snippet. |
| 04 | [Third Party Routers Gorilla Mux Chi](./04-third-party-routers-gorilla-mux-chi.md) | When a router library like `gorilla/mux` or `chi` earns its keep over the standard mux. |
| 05 | [Echo and Gin Frameworks](./05-echo-and-gin-frameworks.md) | Full-featured web frameworks — `Echo` and `Gin` — and what they trade off against the standard library. |
| 06 | [Fiber and Performance](./06-fiber-and-performance.md) | `Fiber`'s `fasthttp` foundation and its performance characteristics/tradeoffs. |
| 07 | [Template Rendering and Htmx](./07-template-rendering-and-htmx.md) | Server-rendered UI with `html/template`, paired with `htmx` for interactivity without a JS framework. |
| 08 | [WebSockets with Gorilla WebSocket](./08-websockets-with-gorilla-websocket.md) | Real-time, bidirectional connections with `gorilla/websocket`. |
| 09 | [Graceful Shutdown](./09-graceful-shutdown.md) | Production-safe shutdown handling — see the working [graceful-shutdown.go](../18-code-snippets/graceful-shutdown.go) snippet. |

**Previous module:** [08 — Modules & Dependency Management](../08-modules-and-dependency-management/README.md) · **Next module:** [10 — Databases](../10-databases/README.md)
