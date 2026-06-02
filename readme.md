# REST API TODOLIST APPLICATION

Реализация backend-части будущего веб-приложения по принципам REST API на **Golang** + **PostgresSQL**

## Структура проекта
```
todolist
├─ cmd
│  └─ todolist
│     └─ main.go               # Точка входа
├─ docker-compose.yaml         # Docker образы
├─ go.mod                      # Зависимости
├─ go.sum
├─ internal
│  ├─ core                     # Общий код
│  │  ├─ domains               # Доменные сущности
│  │  │  ├─ session.go
│  │  │  ├─ task.go
│  │  │  └─ user.go
│  │  ├─ errors                # Ошибки
│  │  │  └─ errors.go
│  │  ├─ logger                # Логгер
│  │  │  ├─ config.go
│  │  │  └─ logger.go
│  │  ├─ server                # Сервер
│  │  │  ├─ config.go
│  │  │  └─ server.go
│  │  ├─ store                 # Хранилище
│  │  │  └─ postgres
│  │  │     ├─ config.go
│  │  │     └─ postgres.go
│  │  └─ transport             # Работа с сетью
│  │     ├─ middleware
│  │     │  └─ middleware.go
│  │     ├─ request
│  │     │  └─ decode.go
│  │     └─ response
│  │        ├─ dto.go
│  │        └─ response.go
│  └─ features                 # Особенности
│     ├─ sessions              # Сессии
│     │  ├─ repository
│     │  │  └─ repository.go
│     │  ├─ service
│     │  │  └─ service.go
│     │  └─ transport
│     │     ├─ dto.go
│     │     └─ transport.go
│     ├─ tasks                 # Задачи
│     │  ├─ repository
│     │  │  └─ repository.go
│     │  ├─ service
│     │  │  └─ service.go
│     │  └─ transport
│     │     ├─ dto.go
│     │     └─ transport.go
│     └─ users                 # Пользователи
│        ├─ repository
│        │  └─ repository.go
│        ├─ service
│        │  └─ service.go
│        └─ transport
│           ├─ dto.go
│           └─ transport.go
├─ Makefile                   # Cборка проекта
├─ .env                       # Переменные окружения
├─ migrations                 # Миграции
│  ├─ 000001_init.down.sql
│  └─ 000001_init.up.sql
├─ out
│  └─ pgdata                  # Файлы БД
└─ readme.md

```

## Схема БД (3 таблицы)

| Таблица            | Первичный ключ                 | Описание                                                    |
| -----------------  | ------------------------------ | ----------------------------------------------------------- |
| **users**          | `id` (GENERATED)               | Пользователь (хранит информацию о пользователе)             |
| **sessions**       | `session_token` (VARCHAR(255)) | Сессия пользователя (хранит файлы cookie + ID пользователя) |
| **tasks**          | `id` (GENERATED)               | Задача (хранит информацию о созданной задаче)               |

## API ENDPOINTS

| Endpoint                         | Описание                              |
| -------------------------------- | ------------------------------------- |
| `POST` /api/users                | Создание нового пользователя          |
| `GET` /api/users/me              | Получение текущей сессии              |
| `POST` /api/sessions             | Создание новой сессии                 | 
| `DELETE` /api/protected/sessions | Удаление текущей сессии               | 
| `POST` /api/protected/tasks      | Создание новой задача                 | 
| `PATCH` /api/protected/tasks{id} | Изменение задачи по ID задачи (ID пользователя берется из текущей сессии) | 
