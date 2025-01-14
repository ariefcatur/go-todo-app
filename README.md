# ToDo App

To Do App API in Go language with Gin and MYSQL 

## Table of Contents
- [Setup](#setup)
- [API Documentation](#api-documentation)

## Setup

### Prerequisites
- Go 1.22.5+
- Gin 1.10.0+
- Mysql 1.5.7+

### Installation

1. Clone the repository
```bash
git clone [https://github.com/ariefcatur/go-todo-app.git]
cd go-todo-app
```

2. Install dependencies
```bash
go mod tidy
```

3. Configure environment variables
```bash
cp .env.example .env
# Edit .env with your credentials
```

5. Start the server
```bash
go run main.go
```

## API Documentation

### Register

```
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: RAHASIA' \
--data '{
    "username": "testuser",
    "password": "pass123"
}'
```

Response:
```json
{
    "message": "User created successfully",
    "userId": 4
}
```

### Login

```
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: RAHASIA' \
--data '{
    "username": "testuser",
    "password": "pass123"
}'
```

Response:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5c*********",
    "user": "testuser"
}
```

### ToDo Operations

#### New Task

```
curl --location 'http://localhost:8080/api/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Belajar GO",
    "description": "Membuat aplikasi Chat",
    "priority": "high"
}'
```

Response:
```json
{
    "ID": 3,
    "CreatedAt": "2025-01-13T16:49:16.835+07:00",
    "UpdatedAt": "2025-01-13T16:49:16.835+07:00",
    "DeletedAt": null,
    "title": "Belajar GO",
    "description": "Membuat aplikasi Chat",
    "status": "pending",
    "priority": "high",
    "user_id": 4
}
```

#### Get Task

```
curl --location 'http://localhost:8080/api/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork'
```

Response:
```json
[
    {
        "ID": 2,
        "CreatedAt": "2025-01-13T16:41:11.731+07:00",
        "UpdatedAt": "2025-01-13T16:41:11.731+07:00",
        "DeletedAt": null,
        "title": "Belajar SWIFT",
        "description": "Membuat aplikasi Chat",
        "status": "pending",
        "priority": "high",
        "user_id": 4
    },
    {
        "ID": 3,
        "CreatedAt": "2025-01-13T16:49:16.835+07:00",
        "UpdatedAt": "2025-01-13T16:49:16.835+07:00",
        "DeletedAt": null,
        "title": "Belajar GO",
        "description": "Membuat aplikasi Chat",
        "status": "pending",
        "priority": "high",
        "user_id": 4
    }
]
```

#### Filter Task

```
curl --location 'http://localhost:8080/api/tasks?status=pending&priority=high' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork'
```

Response:
```json
[
    {
        "ID": 1,
        "CreatedAt": "2025-01-13T16:40:15.814+07:00",
        "UpdatedAt": "2025-01-13T16:40:15.814+07:00",
        "DeletedAt": null,
        "title": "Belajar Go",
        "description": "Membuat aplikasi todo",
        "status": "pending",
        "priority": "high",
        "user_id": 4
    },
    {
        "ID": 2,
        "CreatedAt": "2025-01-13T16:41:11.731+07:00",
        "UpdatedAt": "2025-01-13T16:41:11.731+07:00",
        "DeletedAt": null,
        "title": "Belajar SWIFT",
        "description": "Membuat aplikasi Chat",
        "status": "pending",
        "priority": "high",
        "user_id": 4
    }
]
```

#### Update Task

```
curl --location --request PUT 'http://localhost:8080/api/tasks/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork' \
--header 'Content-Type: application/json' \
--data '{
    "status": "completed"
}'
```

Response:
```json
{
    "ID": 1,
    "CreatedAt": "2025-01-13T16:40:15.814+07:00",
    "UpdatedAt": "2025-01-13T16:47:15.734+07:00",
    "DeletedAt": null,
    "title": "Belajar Go",
    "description": "Membuat aplikasi todo",
    "status": "completed",
    "priority": "high",
    "user_id": 4
}
```

#### Delete Task

```
curl --location --request DELETE 'http://localhost:8080/api/tasks/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork'
```

Response:
```json
{
    "message": "Task deleted successfully"
}
```
