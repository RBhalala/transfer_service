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
### 3. Create .env file and Configure it by adding below values
```bash
DATABASE_URL=postgres://root:root@localhost:5432/transfer_service?sslmode=disable
PORT=8100
```
#### Note: Ensure that you provide your PostgreSQL credentials in the .env file. If you're using the provided docker-compose setup, the default username and password are both root.

### 4. Run application using ()
```bash
go run cmd/main.go
```
#### Note: The default server port is set to 8080. You can change this value in the .env file to match your machine or deployment environment. Make sure the selected port is not in use by another application.

