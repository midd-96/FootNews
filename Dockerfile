FROM golang:alpine AS builder
WORKDIR /footnews
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main main.go


FROM alpine
WORKDIR /footnews
COPY --from=builder /footnews/main .
COPY . .
COPY .env .

EXPOSE 8080