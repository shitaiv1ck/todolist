# API ENDPOINTS

| Endpoint                              | Описание                              |
| ------------------------------------- | ------------------------------------- |
| `POST` /api/users                     | Создание нового пользователя          |
| `GET` /api/users/me                   | Получение текущей сессии (необходимо пройти аутентификацию) |
| `POST` /api/sessions                  | Создание новой сессии                 | 
| `DELETE` /api/protected/sessions      | Удаление текущей сессии               | 
| `POST` /api/protected/tasks           | Создание новой задача                 |
| `GET` /api/tasks            | Получение всех задач пользователя (ID пользователя берется из текущей сессии) | 
| `PATCH` /api/protected/tasks{id}      | Изменение задачи по ID задачи (ID пользователя берется из текущей сессии) | 
| `DELETE` /api/protected/tasks{id}     | Удаление задачи по ID задачи (ID пользователя берется из текущей сессии) |
| `GET` /api/statistics                 | Получение статистики по задачам пользователя (ID пользователя берется из текущей сессии) |  


Для паттернов группы /api/protected/ дополнительно проверятеся аутентификация + csrf токен, переданный в запросе под заголовком **X-CSRF-Token**

## JSON в теле запроса и ответа

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
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
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
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
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
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
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
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
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
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
}
```

**`GET` /api/tasks:**

Request body:
```JSON
{}
```

Response body:
```JSON
200 OK

[
    {
        "id": 5,
        "user_id": 4,
        "title": "test",
        "description": "some description",
        "completed": true,
        "created_at": "2026-06-04T16:52:13.874257Z",
        "completed_at": "2026-06-04T16:52:30.153441Z"
    },
    {
        "id": 6,
        "user_id": 4,
        "title": "test again",
        "description": "some description",
        "completed": false,
        "created_at": "2026-06-04T16:59:21.126815Z",
        "completed_at": null
    }
]
```
```JSON
401 Unautorized

{
    "message": "failed to authentication",
    "error": "invalid cookie's values"
}
```
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
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
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
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
```JSON
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
}
```

**`GET` /api/statistics:**

Request body:
```JSON
{}
```

Response body:
```JSON
200 OK

{
    "total_tasks": 2,
    "total_completed": 1,
    "total_completed_rate": 50,
    "avg_completed_time": "3h22m11.001571s"
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
500 Internal Server Error

{
    "message": "failed to create user",
    "error": "some error"
}
```
