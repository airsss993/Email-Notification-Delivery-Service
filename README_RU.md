# Сервис Рассылки Уведомлений по Email

Масштабируемый сервис для отправки email-сообщений с использованием шаблонов и очереди задач. Реализован на Go, Gin, Redis и PostgreSQL.

## Возможности

- Отправка email-уведомлений через SMTP
- Хранение и рендеринг шаблонов (Go templates)
- Очередь задач на базе Redis
- Пул воркеров для параллельной обработки
- REST API для постановки задач и управления шаблонами

## Требования

- Go 1.21+
- Redis
- PostgreSQL
- SMTP-аккаунт (например, Яндекс или Gmail)

## Быстрый старт

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/yourusername/notification-service.git
cd notification-service
```

### 2. Настройте переменные окружения

Создайте файл `.env` в корне проекта.

Пример содержимого `.env.example`:

```env
DB_URL=postgres://<user>@localhost:5432/notificationservice?sslmode=disable
SMTP_HOST=smtp.yandex.ru
SMTP_PORT=587
SMTP_USER=<your_smtp_user>
SMTP_PASS=<your_smtp_password>
SMTP_EMAIL=<your_email>
```

Убедитесь, что файл `.env` загружается при старте. Если вы используете `github.com/joho/godotenv`, добавьте `godotenv.Load()` в `main.go`.

### 3. Запустите PostgreSQL и Redis

Redis и PostgreSQL должны быть запущены и доступны по адресам, указанным в `.env`.

### 4. Запустите приложение

```bash
cd cmd && go run main.go
```

Приложение будет доступно по адресу: `http://localhost:8080`.

## API

### Поставить задачу на отправку

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

### Управление шаблонами

* `POST /templates` — создать шаблон  
* `GET /templates/:id` — получить шаблон  
*Другие методы (обновление, удаление) могут быть добавлены позже.*