# Go Todo App (Gin + PostgreSQL)

A simple Todo API built with **Go**, **Gin**, **GORM**, and **PostgreSQL**.  
Supports user registration, login (JWT authentication), and CRUD operations for tasks.

---

## 🚀 Features
- User registration & login with hashed password (bcrypt)
- JWT-based authentication (middleware protected routes)
- CRUD tasks (with `status` and `priority`)
- Dockerized setup (App + PostgreSQL + Adminer)
- Unit tests (SQLite in-memory)
- Integration tests with Docker + Postgres

---

## 📦 Requirements
- Go 1.22+ (or newer)
- Docker & Docker Compose (for quick setup)
- Make (optional)

---

## ⚙️ Environment Variables

Copy `.env.example` → `.env`, then configure:

```env
PORT=8080
GIN_MODE=release

JWT_SECRET=supersecret
JWT_EXP_MIN=60

# PostgreSQL connection string (using Docker)
DB_DSN=host=db user=app password=app dbname=todo port=5432 sslmode=disable TimeZone=Asia/Jakarta

# For local (without Docker):
# DB_DSN=host=127.0.0.1 user=app password=app dbname=todo port=5432 sslmode=disable TimeZone=Asia/Jakarta
```

> **Note:** the app only reads `DB_DSN`, so make sure it follows the PostgreSQL DSN format.

---

## ▶️ Run with Docker

Build & start:

```bash
docker compose up --build
```

Access:
- API → [http://localhost:8080](http://localhost:8080)  
- Adminer (DB GUI) → [http://localhost:8081](http://localhost:8081)  

Adminer login:
- **System:** PostgreSQL  
- **Server:** db  
- **Username:** app  
- **Password:** app  
- **Database:** todo  

---

## 🌱 Seed Database

### Option A (via API - recommended)
Register a new user:
```bash
curl -X POST http://localhost:8080/register   -H "Content-Type: application/json"   -d '{"username":"demo","email":"demo@example.com","password":"pass12345"}'
```

Then login to get a token:
```bash
curl -X POST http://localhost:8080/login   -H "Content-Type: application/json"   -d '{"identity":"demo@example.com","password":"pass12345"}'
```

### Option B (via SQL)
1. Generate a bcrypt hash for the password:
   ```bash
   go run ./scripts/bcrypt_hash.go pass12345
   ```
   (this script will print the bcrypt hash to the terminal)

2. Edit `scripts/seed.sql`, replace `<PASTE_BCRYPT_HASH_HERE>` with the hash.

3. Run the seed script:
   ```bash
   docker compose exec -T db psql -U app -d todo -f scripts/seed.sql
   ```

---

## 🔑 API Endpoints

### Auth
- `POST /register` → register new user
- `POST /login` → login and get JWT token

### Tasks (require Bearer token)
- `GET /api/tasks?status=pending|completed&priority=low|medium|high`
- `POST /api/tasks` → create new task
- `PUT /api/tasks/:id` → update task
- `DELETE /api/tasks/:id` → delete task

---

## 🧪 Testing

### Unit Test (SQLite in-memory)
```bash
go test ./...
```

### Integration Test (Postgres via Docker)
1. Start database:
   ```bash
   docker compose up -d db
   ```
2. Run integration tests (if configured for real Postgres).

---

## 🛠 Development (without Docker)
```bash
go mod tidy
go run main.go
```

Server will run on [http://localhost:8080](http://localhost:8080).

---

## 📂 Project Structure

```
.
├── config/             # app config & DB connection
├── controllers/        # handlers (auth & tasks)
├── helpers/            # helper utilities (JWT, response, email, etc)
├── middlewares/        # JWT auth & security middleware
├── models/             # GORM models
├── scripts/            # seed.sql & utility scripts
├── main.go             # entry point
└── README.md
```

---

## 🧹 Future Improvements
- Refresh token implementation
- Pagination for task list
- CI/CD workflow (GitHub Actions)
