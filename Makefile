GIN = github.com/gin-gonic/gin
JWT = github.com/golang-jwt/jwt/v5
LOGRUS = github.com/sirupsen/logrus github.com/sirupsen/logrus@v1.9.3
CLEANENV = github.com/ilyakaznacheev/cleanenv
POSTGRESQL = github.com/jackc/pgx github.com/jackc/pgx/v5/pgxpool
ENV = github.com/joho/godotenv
SWAG := github.com/swaggo/swag/cmd/swag
GIN_SWAG := github.com/swaggo/gin-swagger github.com/swaggo/files

GOOSE = github.com/pressly/goose
EXEC :=  build/miniapp
SRC := cmd/app/main.go

all: clean build run

init:
	go mod init miniapp

get:
	go get -u $(GIN) $(JWT) $(LOGRUS) $(CLEANENV) $(POSTGRESQL) $(ENV) $(SWAG) $(GIN_SWAG)

build:
	mkdir -p build
	go build -o $(EXEC) $(SRC)

run:
	$(EXEC)

swag:
	swag fmt
	swag init -g main.go -d cmd/app,internal/handlers/,internal/dto

goose_set_path:
	export PATH=$(PATH):home/$(USER)/go/bin/

goose_up:
	# goose -dir migrations postgres "postgresql://postgres:$(PSQLPASS)@postgres:5432/postgres?sslmode=disable" up
	fish -c 'goose -dir migrations postgres "postgresql://postgres:$(PSQLPASS)@0.0.0.0:5432/postgres?sslmode=disable" up'

goose_down:
	goose -dir migrations postgres "postgresql://postgres:$(PSQLPASS)@0.0.0.0:5432/postgres?sslmode=disable" down

.PHONY: all init get build goose_set_path goose_up goose_down

test:
	echo "oops"

clean:
	rm -rf $(EXEC)