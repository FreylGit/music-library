# Онлайн Библиотека Песен 🎶

Данный проект представляет собой сервис для управления библиотекой песен. Реализованы REST-методы для работы с песнями и интеграция с внешним API, а также поддержка PostgreSQL для хранения данных.

---

## 📋 Функциональность

- Получение списка песен с фильтрацией и пагинацией.
- Получение текста песни с пагинацией по куплетам.
- Добавление новой песни с запросом в API для обогащения информации.
- Изменение данных песни.
- Удаление песни.
- Миграции для инициализации структуры базы данных.
- Логирование (debug и info).
- Использование переменных окружения из `local.env`.

---

## 🚀 Запуск проекта

### 1. Подготовка окружения
1. Убедитесь, что на вашей машине установлен Docker и Docker Compose.

### 2. Запуск контейнеров
Для запуска всех необходимых сервисов выполните команду:
```bash
docker compose --env-file local.env up -d
```

### 3. Запуск приложения
Приложение запускается командой:
```bash
 go run ./cmd/main.go
```

---
## 🌐 Интеграция с внешним API

Для обогащения данных песни при добавлении используется внешний API. Чтобы указать хост с вашим Swagger, отредактируйте функцию в файле `/internal/services/song/add`:
```go
    const host = "http://localhost:8080" // Укажите ваш хост здесь
```
---
## 💻 Используемые технологии

- **Go**
- **PostgreSQL**
- **Docker & Docker Compose**
- **Swagger**

---

## 🗂️ Структура проекта

- `/cmd/main.go` — точка входа приложения.
- `/internal` — бизнес-логика и сервисы.
- `/migrations` — SQL-файлы миграций.

---

## 🚨 Примечание

Для интеграции с внешним API (`/info`) необходимо, чтобы данный сервис был доступен. В противном случае добавление новой песни не будет обогащено информацией.