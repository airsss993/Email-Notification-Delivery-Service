# Email Notification Delivery Service

A scalable service for sending email using templates and a task queue. Built with Go, Gin, Redis, and PostgreSQL.

📄 [Документация на русском](README_RU.md)

## Features

- Email notification delivery (SMTP)
- Template storage and rendering (Go templates)
- Redis-backed task queue
- Worker pool for concurrent processing
- Basic REST API for enqueuing tasks and managing templates

## Requirements

- Go 1.21+
- Redis
- PostgreSQL
- SMTP account (e.g. Yandex, Gmail)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/notification-service.git
cd notification-service
```

### 2. Configure environment variables

Create a `.env` file in the project root.

Refer to `.env.example`:

```env
DB_URL=postgres://<user>@localhost:5432/notificationservice?sslmode=disable
SMTP_HOST=smtp.yandex.ru
SMTP_PORT=587
SMTP_USER=<your_smtp_user>
SMTP_PASS=<your_smtp_password>
SMTP_EMAIL=<your_email>
```

Make sure the `.env` file is loaded at startup. If you use `github.com/joho/godotenv`, call `godotenv.Load()` early in
`main.go`.

### 3. Start PostgreSQL and Redis

Ensure PostgreSQL and Redis are running locally and accessible using the connection details in your `.env`.

### 4. Run the application

```bash
cd cmd && go run main.go
```

The service will be available at `http://localhost:8080`.

## API Reference

### Enqueue a Notification Task

**POST** `/send`

```json
{
  "template_id": 1,
  "to": "user@example.com",
  "params": {
    "name": "John"
  }
}
```

### Manage Templates

* `POST /templates` — create new template
* `GET /templates/:id` — retrieve template  
  *Other endpoints (update, delete) may be added in future versions.*