FROM golang:1.20.4-alpine3.18
WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 8000
CMD ["/app/main"]