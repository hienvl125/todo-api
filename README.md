### Dependecies
- Gin
- Sqlx
- Wire
- Envconfig

### Development
#### 1. Prepare envinronment variables
```
# local.env
DSN="postgres://postgres:password@127.0.0.1/todo_db_development?sslmode=disable"
GIN_PORT=":7070"
```
#### 2. Export envinroment variables
```
export $(cat local.env | xargs)
```
#### 3. Start a go server
```
go run cmd/api/main.go
```
and visit http://localhost:7070/todos
