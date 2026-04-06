# Libro

Libro is a personal book management backend built with Go, Fiber, GORM, MySQL, and Redis.

## Overview

Libro follows a layered modular architecture mirroring the reference project style:
- `apiSchema` for request/response contracts per module
- `controllers` as Fiber delivery layer
- `services` for business logic
- `repositories` for persistence access and initialization
- `models` for module data models and support structures
- `middleware` for auth concerns
- `pkg` for shared technical utilities
- `statics` for config/constants/custom errors/messages
- `migrations` for raw SQL schema changes
- `tests` for module-based tests
- `template` for scaffolding conventions

## Folder Structure

```
/apiSchema
/controllers
/middleware
/migrations
/models
/pkg
/repositories
/services
/statics
/template
/tests
main.go
```

## Environment Setup

1. Copy values from `dev.env` and adjust as needed.
2. Ensure MySQL and Redis are running.
3. Create database (default: `libro`).

## MySQL Setup

- Host, credentials, and DB are loaded from `dev.env`.
- GORM auto-migration runs on startup.
- Raw SQL files in `/migrations` are provided for external migration tools.

## Redis Setup

- Redis is used for refresh token storage/invalidation.

## Migration Usage

Use your preferred migration runner against files in `/migrations`.

## Run

```bash
go mod tidy
go run main.go
```

## Test

```bash
go test ./...
```

## API Summary

Public:
- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/refresh`

Protected:
- `POST /auth/logout`
- `GET /auth/me`
- `GET/POST/PUT/DELETE /books...`
- `GET /reading/current`
- `PATCH /reading/books/:id/progress`
- `GET/POST/PUT/DELETE /wishlist...`
- `POST/PUT/DELETE /wishlist/:id/links...`
- `GET /user/profile`
- `PUT /user/profile`
- `PATCH /user/password`

Monitoring:
- `GET /health`
- `GET /main/dashboard-summary`

## Authentication Notes

- Passwords are hashed with bcrypt.
- Access and refresh tokens are JWT.
- Refresh tokens are stored in Redis and invalidated on logout.
