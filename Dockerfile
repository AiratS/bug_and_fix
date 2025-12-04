FROM golang:1.25.5-alpine3.23

WORKDIR /app

COPY go.mod go.sum ./

COPY . ./

RUN go mod download
RUN go build -o app ./cmd/app/main.go

EXPOSE 3000

CMD ["./app"]