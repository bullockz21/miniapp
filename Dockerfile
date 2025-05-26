FROM golang
ENV GOPROXY=https://goproxy.io,https://proxy.golang.org,https://gocenter.io,direct

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./backend/
COPY ./.env ./

RUN cd backend && make build

# RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["./backend/build/miniapp"]

