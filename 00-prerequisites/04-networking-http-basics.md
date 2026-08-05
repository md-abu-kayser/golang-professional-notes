# Networking & HTTP Basics

## Overview

Almost every real Go program you'll build in this repo — REST APIs, microservices, websocket chat — communicates over HTTP. This lesson covers the networking fundamentals underneath `net/http` before you start calling it.

## Why It Matters

`net/http` (covered in [05-standard-library](../05-standard-library/README.md) and [09-web-development](../09-web-development/README.md)) will make far more sense once you know what a request/response cycle, a status code, and a header actually represent — rather than treating them as magic strings.

## Core Concepts

### The client-server model

A client (browser, mobile app, another service) sends a **request** to a server; the server sends back a **response**. Go programs act as both — as a server handling incoming requests, and as a client calling other APIs.

### TCP/IP, in one paragraph

Data travels the internet in packets, routed between machines using **IP addresses**. **TCP** sits on top of IP and guarantees packets arrive in order and without loss — it's what makes a "connection" reliable. HTTP is built on top of TCP.

### DNS

Humans use domain names (`api.example.com`); machines route by IP address. **DNS** is the lookup service that translates one into the other before a connection is even opened.

### Anatomy of an HTTP request

```http
GET /users/42 HTTP/1.1
Host: api.example.com
Authorization: Bearer eyJhbGciOi...
Accept: application/json
```

- **Method** (`GET`, `POST`, `PUT`, `PATCH`, `DELETE`) — the intended action.
- **Path** (`/users/42`) — the resource being addressed.
- **Headers** — metadata (auth tokens, content type, caching hints).
- **Body** (for `POST`/`PUT`/`PATCH`) — the payload being sent.

### Anatomy of an HTTP response

```http
HTTP/1.1 200 OK
Content-Type: application/json

{"id": 42, "name": "Ada"}
```

### Status code families

| Range | Meaning | Examples |
|---|---|---|
| 2xx | Success | `200 OK`, `201 Created`, `204 No Content` |
| 3xx | Redirection | `301 Moved Permanently`, `304 Not Modified` |
| 4xx | Client error | `400 Bad Request`, `401 Unauthorized`, `404 Not Found` |
| 5xx | Server error | `500 Internal Server Error`, `503 Service Unavailable` |

## Common Pitfalls

- Returning `200 OK` for every response, including errors — clients (and monitoring tools) rely on accurate status codes.
- Confusing `401 Unauthorized` (not authenticated) with `403 Forbidden` (authenticated but not allowed).
- Forgetting that HTTP is stateless — each request stands alone unless you explicitly carry state (tokens, cookies, sessions).

## Key Takeaways

- HTTP is a request/response protocol built on TCP/IP; DNS resolves domain names to IP addresses first.
- A request has a method, path, headers, and optional body; a response has a status code, headers, and optional body.
- Status codes are grouped by first digit: 2xx success, 3xx redirect, 4xx client error, 5xx server error.
- HTTP is stateless by design — Go servers must explicitly manage anything that needs to persist across requests.
