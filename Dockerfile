FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app ./app
EXPOSE 8080
CMD ["./app"]
