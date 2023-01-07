migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go-chat?sslmode=disable" -verbase up
migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/go-chat?sslmode=disable" -verbose up

sql library: https://pkg.go.dev/database/sql
lib/pq: https://github.com/lib/pq
golang-migrate: https://github.com/golang-migrate/mig...
golang-jwt: https://github.com/golang-jwt/jwt
bcrypt: https://pkg.go.dev/golang.org/x/crypt...
gin: https://github.com/gin-gonic/gin