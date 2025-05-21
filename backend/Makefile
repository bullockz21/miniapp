GIN := github.com/gin-gonic/gin 
JWT := github.com/golang-jwt/jwt/v5
LOGRUS := github.com/sirupsen/logrus github.com/sirupsen/logrus@v1.9.3
CLEANENV := github.com/ilyakaznacheev/cleanenv
POSTGRESQL := github.com/jackc/pgx github.com/jackc/pgx/v5/pgxpool

all:
	go run cmd/app/main.go

init:
	go mod init miniapp

get:
	go get $(GIN)
	go get $(JWT)
	go get $(LOGRUS)
	go get $(CLEANENV)
	go get $(POSTGRESQL)
