FROM golang:latest

WORKDIR /app

RUN apt-get update && apt-get install -y postgresql-client

COPY go.mod go.sum ./

RUN go mod download

COPY wait-for-db.sh .  
COPY . .

RUN go build -o noteapp ./cmd/api

EXPOSE 8000

CMD ["/app/noteapp"]
