# REST API TODOLIST APPLICATION

Реализация backend-части будущего веб-приложения по принципам REST API на **Golang** + **PostgreSQL**

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
│  │     │  ├─ decode.go
│  │     │  └─ pathvalue.go
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
| **tasks**          | `id` (GENERATED)               | Задача (хранит информацию о созданной задаче + ID пользователя) |

## API ENDPOINTS

| Endpoint                              | Описание                              |
| ------------------------------------- | ------------------------------------- |
| `POST` /api/users                     | Создание нового пользователя          |
| `GET` /api/users/me                   | Получение текущей сессии (необходимо пройти аутентификацию)             |
| `POST` /api/sessions                  | Создание новой сессии                 | 
| `DELETE` /api/protected/sessions      | Удаление текущей сессии               | 
| `POST` /api/protected/tasks           | Создание новой задача                 | 
| `PATCH` /api/protected/tasks{id}      | Изменение задачи по ID задачи (ID пользователя берется из текущей сессии) | 
| `DELETE` /api/protected/tasks{id}     | Удаление задачи по ID задачи (ID пользователя берется из текущей сессии) | 

Для паттернов группы /api/protected/ дополнительно проверятеся аутентификация + csrf токен, переданный в запросе под заголовком **X-CSRF-Token**

### JSON в теле запроса и ответа

**`POST` /api/users:**

Request body:
```JSON
{
    "username": "some username",  
    "password": "some password"  
}
```

Response body:
```JSON
201 Created

{
    "id": 3,
    "username": "some username"
}
```
```JSON
400 Bad Request

{
    "message": "failed to decode and validate",
    "error": "wrong username or password: invalid argument"
}
```
```JSON
409 Conflict

{
    "message": "failed to create user",
    "error": "failed to save user: already exists"
}
```

**`GET` /api/users/me:**

Request body:
```JSON
{}
```

Response body:
```JSON
200 OK

{}
```
```JSON
401 Unautorized

{
    "message": "failed to authentication",
    "error": "invalid cookie's values"
}
```

**`POST` /api/sessions:**

Request body:
```JSON
{
    "username": "some username",  # required, min=8, max=100, not null
    "password": "some password"   # required, min=8, max=100, not null
}
```

Response body:
```JSON
201 Created

{}
```
```JSON
401 Unautorized

{
    "message": "failed to authentication",
    "error": "invalid username or password"
}
```

**`DELETE` /api/protected/sessions:**

Request body:
```JSON
{}
```

Response body:
```JSON
204 No Content

{}
```
```JSON
401 Unautorized

{
    "message": "failed to authentication",
    "error": "invalid cookie's values"
}
```

**`POST` /api/protected/tasks:**

Request body:
```JSON
{
    "title": "some title",             # required, max=100, not null
    "description": "some description"  # optional, max=1000
}
```

Response body:
```JSON
201 Created

{
    "id": 13,
    "user_id": 3,
    "title": "some title",
    "description": "some description",
    "completed": false,
    "created_at": "2026-06-03T18:27:51.524537Z",
    "completed_at": null
}
```
```JSON
400 Bad Request

{
    "message": "failed to decode and validate",
    "error": "wrong username or password: invalid argument"
}
```
```JSON
401 Unautorized

{
    "message": "failed to authentication",
    "error": "invalid cookie's values"
}
```
```JSON
409 Conflict

{
    "message": "failed to create task",
    "error": "failed to save task: already exists"
}
```

**`PATCH` /api/protected/tasks{id}:**

Request body:
```JSON
{
    "title": "patched title",             # optional, max=100, not null
    "description": "patched description", # optional, max=1000
    "completed": true                     # optional, not null
}
```

Response body:
```JSON
200 OK

{
    "id": 13,
    "user_id": 3,
    "title": "patched title",
    "description": "patched description",
    "completed": true,
    "created_at": "2026-06-03T18:27:51.524537Z",
    "completed_at": "2026-06-03T21:30:06.369811Z"
}
```
```JSON
400 Bad Request

{
    "message": "failed to patch task",
    "error": "failed to apply patches: completed can't be null: invalid argument"
}
```
```JSON
401 Unautorized

{
    "message": "failed to authentication",
    "error": "invalid cookie's values"
}
```
```JSON
404 Not Found

{
    "message": "failed to patch task",
    "error": "not found"
}
```

**`DELETE` /api/protected/tasks{id}:**

Request body:
```JSON
{}
```

Response body:
```JSON
204 No Content

{}
```
```JSON
404 Not Found

{
    "message": "failed to delete task",
    "error": "task doesn't exist: not found"
}
```
```JSON
401 Unautorized

{
    "message": "failed to authentication",
    "error": "invalid cookie's values"
}
```

