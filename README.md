# To Do App

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
--data-raw '{
    "username": "testuser1",
    "email": "test1@example.com",
    "password": "pass123"
}'
```


### Login

```
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--header 'X-API-Key: RAHASIA' \
--data-raw '{
    
    "identity": "test@example.com",
    
    "password": "pass123"
}'
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


#### Get Task

```
curl --location 'http://localhost:8080/api/tasks' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork'
```


#### Filter Task

```
curl --location 'http://localhost:8080/api/tasks?status=pending&priority=high' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork'
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


#### Delete Task

```
curl --location --request DELETE 'http://localhost:8080/api/tasks/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY4MjU2NDQsInVzZXJfaWQiOjR9.sDRJq8PQfajl-6Lmc-VzIkS-X7WAo3DnPbdFw5bsork'
```

