# Структура проекта
```
todolist
├─ api.md
├─ architecrure.md
├─ cmd
│  └─ todolist
│     └─ main.go
├─ docker-compose.yaml
├─ docs
├─ go.mod
├─ go.sum
├─ internal
│  ├─ core
│  │  ├─ domains
│  │  │  ├─ nullable.go
│  │  │  ├─ session.go
│  │  │  ├─ task.go
│  │  │  └─ user.go
│  │  ├─ errors
│  │  │  └─ errors.go
│  │  ├─ logger
│  │  │  ├─ config.go
│  │  │  └─ logger.go
│  │  ├─ server
│  │  │  ├─ config.go
│  │  │  └─ server.go
│  │  ├─ store
│  │  │  └─ postgres
│  │  │     ├─ config.go
│  │  │     └─ postgres.go
│  │  └─ transport
│  │     ├─ middleware
│  │     │  └─ middleware.go
│  │     ├─ request
│  │     │  ├─ decode.go
│  │     │  └─ pathvalue.go
│  │     └─ response
│  │        ├─ dto.go
│  │        └─ response.go
│  └─ features
│     ├─ sessions
│     │  ├─ repository
│     │  │  └─ repository.go
│     │  ├─ service
│     │  │  └─ service.go
│     │  └─ transport
│     │     ├─ dto.go
│     │     └─ transport.go
│     ├─ tasks
│     │  ├─ repository
│     │  │  └─ repository.go
│     │  ├─ service
│     │  │  └─ service.go
│     │  └─ transport
│     │     ├─ dto.go
│     │     └─ transport.go
│     ├─ users
│     │  ├─ repository
│     │  │  └─ repository.go
│     │  ├─ service
│     │  │  └─ service.go
│     │  └─ transport
│     │     ├─ dto.go
│     │     └─ transport.go
│     └─ web
│        ├─ repository
│        │  └─ repository.go
│        ├─ service
│        │  └─ service.go
│        └─ transport
│           └─ transport.go
├─ Makefile
├─ migrations
│  ├─ 000001_init.down.sql
│  └─ 000001_init.up.sql
├─ public
│  └─ index.html
└─ readme.md

```

## Схема БД (3 таблицы)

| Таблица            | Первичный ключ                 | Описание                                                    |
| -----------------  | ------------------------------ | ----------------------------------------------------------- |
| **users**          | `id` (GENERATED)               | Пользователь (хранит информацию о пользователе)             |
| **sessions**       | `session_token` (VARCHAR(255)) | Сессия пользователя (хранит файлы cookie + ID пользователя) |
| **tasks**          | `id` (GENERATED)               | Задача (хранит информацию о созданной задаче + ID пользователя) |