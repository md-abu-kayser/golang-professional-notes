# 11 — API Development

> 🚧 **Status:** Syllabus ready — 9 lessons outlined, full content in progress.

Designing and building APIs in Go across four major styles — REST, gRPC, and GraphQL — plus the cross-cutting concerns (validation, auth, docs) every real API needs.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Rest API Design Principles](./01-rest-api-design-principles.md) | Resource modeling, versioning strategy, and idempotency in REST API design. |
| 02 | [Building a Rest API with Net HTTP](./02-building-a-rest-api-with-net-http.md) | A complete REST API built with only the standard library. |
| 03 | [Building Apis with Echo Gin](./03-building-apis-with-echo-gin.md) | The same API rebuilt with a framework, to compare ergonomics directly. |
| 04 | [Validation with Go Playground Validator](./04-validation-with-go-playground-validator.md) | Struct-tag-based request validation with `go-playground/validator`. |
| 05 | [Authentication JWT and Middleware](./05-authentication-jwt-and-middleware.md) | A complete JWT auth flow — see the working [jwt-auth-middleware.go](../18-code-snippets/jwt-auth-middleware.go) snippet. |
| 06 | [OAuth2 and Google Login](./06-oauth2-and-google-login.md) | Implementing the OAuth2 authorization code flow, using Google login as a worked example. |
| 07 | [gRPC and Protobuf](./07-grpc-and-protobuf.md) | Defining services and messages with Protocol Buffers and generating gRPC code. |
| 08 | [GraphQL with GraphQL Go](./08-graphql-with-graphql-go.md) | Building a GraphQL API in Go with `graphql-go`. |
| 09 | [OpenAPI and Swagger Documentation](./09-openapi-and-swagger-documentation.md) | Generating and maintaining OpenAPI/Swagger specs alongside your handlers. |

**Previous module:** [10 — Databases](../10-databases/README.md) · **Next module:** [12 — Microservices & Distributed Systems](../12-microservices-and-distributed-systems/README.md)
