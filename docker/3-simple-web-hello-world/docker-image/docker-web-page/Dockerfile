FROM golang:1.13-alpine

COPY . /app
RUN go build -o /app/app /app/src/main.go
CMD "/app/app"
