# go-auth
## Setup
go mod tidy
go run main.go

## APIs
POST /signup  - body: {"email":"","password":""}
POST /login   - body: {"email":"","password":""}
GET  /profile - Header: Authorization: Bearer <token>
GET  /users   - Header: Authorization: Bearer <token> (admin only)