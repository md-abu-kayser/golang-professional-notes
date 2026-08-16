# 10 — Databases

> 🚧 **Status:** Syllabus ready — 8 lessons outlined, full content in progress.

Talking to relational and NoSQL databases from Go: the standard `database/sql` abstraction, popular drivers, ORMs, code generators, migrations, and caching.

## Lessons

| # | Lesson | Covers |
|---|---|---|
| 01 | [Database SQL and Connection Pooling](./01-database-sql-and-connection-pooling.md) | The `database/sql` package and tuning its built-in connection pool. |
| 02 | [PostgreSQL with Pgx](./02-postgresql-with-pgx.md) | Working with PostgreSQL via the high-performance `pgx` driver. |
| 03 | [MySQL and SQLite](./03-mysql-and-sqlite.md) | Driver setup and usage for MySQL and embedded SQLite. |
| 04 | [ORM GORM Basics](./04-orm-gorm-basics.md) | GORM fundamentals, and the tradeoffs of using an ORM in Go at all. |
| 05 | [Sqlc and Query Building](./05-sqlc-and-query-building.md) | Generating type-safe Go code directly from hand-written SQL with `sqlc`. |
| 06 | [Migrations Golang Migrate](./06-migrations-golang-migrate.md) | Managing schema changes over time with `golang-migrate`. |
| 07 | [Redis and Caching](./07-redis-and-caching.md) | Caching patterns with `go-redis`, from simple key-value caching to rate limiting. |
| 08 | [NoSQL and MongoDB Driver](./08-nosql-and-mongodb-driver.md) | Working with MongoDB using its official Go driver. |

**Previous module:** [09 — Web Development](../09-web-development/README.md) · **Next module:** [11 — API Development](../11-api-development/README.md)
