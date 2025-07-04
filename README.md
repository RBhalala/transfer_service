# Transfer Service

A Go application to process financial transfers between accounts with a PostgreSQL backend.

---

## Features

- Create accounts with initial balances
- Query account balances
- Submit transactions between accounts
- HTTP API implemented with [chi router](https://github.com/go-chi/chi)
- PostgreSQL to store accounts and transaction logs
- Maintains data integrity with transactional updates

---

## Requirements

- Go 1.20+
- Docker (for easy Postgres setup)
- PostgreSQL 12+

---

## Setup & Run

### 1. Clone the repo

```bash
git clone https://github.com/rbhalala/transfer_service.git
cd transfer_service
```
### 2. Run PostgreSQL locally with Docker Compose and apply migrations
```bash
docker-compose up migrate
```


