FROM golang:1.25.1

WORKDIR /app

COPY ./ ./
RUN go mod tidy
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o main ./cmd/

EXPOSE 8080

CMD ["./main"]