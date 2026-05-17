# go-auth — Golang Authentication Service

A backend authentication service built with Go, featuring JWT-based authentication, role-based access control, and SQLite database.

## Tech Stack

- **Go** — core language
- **Gin** — HTTP web framework
- **GORM** — ORM for database interaction
- **SQLite** — database (zero config, single file)
- **golang-jwt/jwt** — JWT token generation and validation
- **bcrypt** — password hashing

## Project Structure

```
go-auth/
├── main.go              # Entry point, route registration
├── config/
│   └── database.go      # DB connection and auto migration
├── models/
│   └── user.go          # User struct and DB schema
├── controllers/
│   ├── auth.go          # Signup and Login handlers
│   └── user.go          # Profile and Users list handlers
├── middleware/
│   └── auth.go          # JWT validation middleware
└── utils/
    └── jwt.go           # Token generation and validation helpers
```

## Setup & Run

**Prerequisites:** Go 1.21+

```bash
# Clone the repo
git clone https://github.com/harsh-21-chauhan/go-auth.git
cd go-auth

# Install dependencies
go mod tidy

# Run the server
go run main.go
# Server starts at http://localhost:8080
```

## API Endpoints

### POST /signup
Register a new user.

**Request body:**
```json
{
  "email": "user@test.com",
  "password": "secret123"
}
```

**Response:**
```json
{
  "message": "user created",
  "user": { "ID": 1, "email": "user@test.com", "role": "user" }
}
```

---

### POST /login
Authenticate and receive a JWT token.

**Request body:**
```json
{
  "email": "user@test.com",
  "password": "secret123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### GET /profile
Returns the authenticated user's own profile.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "user": { "ID": 1, "email": "user@test.com", "role": "user" }
}
```

---

### GET /users
Returns all users. **Admin role only.**

**Headers:**
```
Authorization: Bearer <admin_token>
```

**Response:**
```json
{
  "count": 2,
  "users": [ ... ]
}
```

## Roles

| Role  | /profile | /users |
|-------|----------|--------|
| user  | ✅       | ❌ 403 |
| admin | ✅       | ✅     |

## Creating an Admin User

Signup normally, then update the role directly in the database:

```bash
sqlite3 auth.db "UPDATE users SET role='admin' WHERE email='admin@test.com';"
```

## Auth Flow

1. User signs up — password is hashed with bcrypt and stored
2. User logs in — bcrypt compares input with stored hash
3. On success, a JWT is generated containing `user_id`, `role`, and `exp`
4. Client sends JWT as `Authorization: Bearer <token>` on protected routes
5. `AuthMiddleware` validates the token and sets `userID` and `role` on the request context
6. Handlers use the role for access control (admin-only routes return 403 for regular users)
