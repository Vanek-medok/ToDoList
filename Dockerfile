FROM golang:1.27.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/API

EXPOSE 8080

CMD ["./app"]


