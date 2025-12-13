---

# Notes App

A backend-focused **Notes Application** built with **Go (Gin)** and **PostgreSQL**, implementing authentication, secure password handling, and CRUD operations for notes.
Designed with clear separation of concerns (controller → service → database).

---

## Tech Stack

* **Language:** Go
* **Framework:** Gin
* **Database:** PostgreSQL
* **ORM:** GORM
* **Auth:** bcrypt (password hashing)
* **Infra:** Docker (Postgres)
* **Tools:** Air (hot reload), Postman

---

## Features

* User Registration
* User Login (bcrypt password verification)
* Secure password storage (hashed, never returned)
* Notes CRUD (per authenticated user)
* Email uniqueness enforced at DB level
* Clean controller–service architecture
* Proper error handling (no silent failures)

---

## Project Structure

```
.
├── controller/        # HTTP handlers (Gin)
├── service/           # Business logic
├── models/            # GORM models
├── dto/               # Request/response DTOs
├── validator/         # Input validation
├── db/                # Database initialization
├── middleware/        # Auth / logging middleware
├── routes/            # Route registration
├── docker-compose.yml
├── .env
└── main.go
```

---

## Environment Variables

Create a `.env` file:

```env
DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
PORT=8000
```

---

## Running the App

### 1. Start PostgreSQL (Docker)

```bash
docker-compose up -d
```

---

### 2. Run the server

```bash
go run main.go
```

Or with hot reload:

```bash
air
```

Server will start on:

```
http://localhost:8000
```

---

## API Endpoints

### Auth

#### Register

```http
POST /auth/register
```

Request body:

```json
{
  "name": "Harsh Kharwar",
  "email": "harsh@example.com",
  "password": "StrongPass123"
}
```

---

#### Login

```http
POST /auth/login
```

Request body:

```json
{
  "email": "harsh@example.com",
  "password": "StrongPass123"
}
```

---

### Notes (example)

```http
GET    /notes
POST   /notes
PUT    /notes/:id
DELETE /notes/:id
```

> Notes are scoped per authenticated user.

---

## Security Considerations

* Passwords are **hashed using bcrypt**
* No password or hash is ever returned in responses
* Same error message for invalid email/password (prevents enumeration)
* Unique email enforced at database level

---

## Known Improvements (Planned)

* JWT access & refresh tokens
* Auth middleware
* Rate limiting on login
* DTO-based validation everywhere
* Pagination for notes
* Unit tests for services

---
