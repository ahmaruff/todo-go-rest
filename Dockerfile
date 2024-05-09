FROM golang:1.21 AS builder

ENV TODO_SERVER_ADDRESS 0.0.0.0
ENV TODO_SERVER_PORT 80
ENV HOST 0.0.0.0
ENV PORT 80

WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /app .

FROM debian:bullseye-slim
COPY --from=builder /app /app
EXPOSE 80

ENTRYPOINT ["/app"]