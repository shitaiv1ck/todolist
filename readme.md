# REST API TODOLIST APPLICATION

![Go Version](https://img.shields.io/badge/Go-1.25.5+-00ADD8?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-4169E1?logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)
![REST API](https://img.shields.io/badge/REST-API-blue)

Веб-приложение, реализованное по принципам REST API.

## Функционал

- Регистрация/аутентификация пользователя
- Создание/редактирование/удаление задач
- Статистика всех задач пользователя

## Технологический стэк

- **GOLANG** - backend
- **HTML** + **CSS** + **JS** - frontend
- **PostgreSQL** - хранение данных
- **Migrate** - инструмент для управления миграциями баз данных
- **Makefile** + **Docker** - сборка проекта 

## Установка и запуск

### Требования

- [Migrate](https://github.com/golang-migrate/migrate)
- [Docker](https://www.docker.com/get-started/)

### Как запустить

1. **Создайте** переменные окружения в `.env` в корневой папке проекта:

```properties
POSTGRES_USER=some-username
POSTGRES_PASSWORD=some-password
POSTGRES_DB=some-db

POSTGRES_HOST=localhost
POSTGRES_PORT=5432

HTTP_ADDR=:8080

LOG_LEVEL=DEBUG
```

2. **Запустите**:

```bash
make env-up && \
make migrate-up && \
make app-deploy
```

3. **Откройте** в браузере ссылку http://localhost:8080/

## Демонстрация работы приложения

### Регистрация и аутентификация пользователя

![Registration and authentication](docs/regandauth.gif)

### Создание задачи

![create task](docs/createtask.gif)

### Редактирование задачи

![patch task](docs//patchtask.gif)

### Удаление задачи

![delete task](docs/deletetask.gif)

### Статистика по задачам

![statistics](docs/getstatistics.gif)

#  Документация

- **[Архитектура проекта](architecture.md)** — структура, схема базы данных.
- **[API Reference](api.md)** — описание всех эндпоинтов, примеры запросов/ответов.
