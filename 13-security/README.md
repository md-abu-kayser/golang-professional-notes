# 13 — Security

> 🚧 **Status:** Syllabus ready — 6 lessons outlined, full content in progress.

Practical, defensive security for Go services: input handling, HTTP hardening, rate limiting, secrets management, and automated vulnerability scanning.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Input Validation and Sanitization](./01-input-validation-and-sanitization.md) | Preventing injection attacks through strict, explicit input validation. |
| 02 | [Secure Headers and CSRF](./02-secure-headers-and-csrf.md) | HTTP security headers and CSRF protection for server-rendered apps. |
| 03 | [Rate Limiting](./03-rate-limiting.md) | Token-bucket and sliding-window rate limiting strategies. |
| 04 | [Secrets Management Vault Env](./04-secrets-management-vault-env.md) | Keeping secrets out of source control with environment variables and tools like Vault. |
| 05 | [Static Analysis and Gosec](./05-static-analysis-and-gosec.md) | Automated security scanning of Go source with `gosec`. |
| 06 | [Dependency Vulnerability Scanning](./06-dependency-vulnerability-scanning.md) | Catching vulnerable dependencies with `govulncheck` and supply-chain hygiene practices. |

**Previous module:** [12 — Microservices & Distributed Systems](../12-microservices-and-distributed-systems/README.md) · **Next module:** [14 — Performance Optimization](../14-performance-optimization/README.md)
