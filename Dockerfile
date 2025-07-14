FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Ativa o CGO e instala dependências do SQLite
ENV CGO_ENABLED=1
RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev

RUN go build -o app ./cmd/main.go

EXPOSE 3000

CMD ["./app"]