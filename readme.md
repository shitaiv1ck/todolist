## Структура прокета
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
├─ migrations                 # Миграции
│  ├─ 000001_init.down.sql
│  └─ 000001_init.up.sql
└─ readme.md

```