FROM golang:1.23

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./backend/

COPY backend/.env .env

RUN go build -o miniapp_demo ./backend/cmd/main.go

CMD ["./miniapp_demo"]
