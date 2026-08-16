# 12 — Microservices & Distributed Systems

> 🚧 **Status:** Syllabus ready — 7 lessons outlined, full content in progress.

Moving from a single service to many: how to draw service boundaries, communicate reliably between them, and keep a distributed system observable and resilient.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Microservices Architecture](./01-microservices-architecture.md) | Drawing service boundaries and choosing synchronous vs. asynchronous communication. |
| 02 | [Service Discovery with Consul Etcd](./02-service-discovery-with-consul-etcd.md) | Dynamic service registration and discovery with Consul or etcd. |
| 03 | [Message Queues RabbitMQ Kafka](./03-message-queues-rabbitmq-kafka.md) | Asynchronous messaging patterns using RabbitMQ and Kafka. |
| 04 | [Event Driven Patterns](./04-event-driven-patterns.md) | Event sourcing, pub/sub, and reasoning about eventual consistency. |
| 05 | [Distributed Tracing Opentelemetry](./05-distributed-tracing-opentelemetry.md) | Tracing a single request as it crosses multiple services with OpenTelemetry. |
| 06 | [Circuit Breakers and Resiliency](./06-circuit-breakers-and-resiliency.md) | Preventing cascading failures with circuit breakers and related resiliency patterns — pairs with [retry-with-backoff.go](../18-code-snippets/retry-with-backoff.go). |
| 07 | [gRPC Microservices Example](./07-grpc-microservices-example.md) | A worked, multi-service example wiring everything in this module together over gRPC. |

**Previous module:** [11 — API Development](../11-api-development/README.md) · **Next module:** [13 — Security](../13-security/README.md)
