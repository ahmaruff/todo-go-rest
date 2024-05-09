FROM golang:1.21 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=builder /app /app


ENV TODO_SERVER_ADDRESS 0.0.0.0
ENV TODO_SERVER_PORT 80
ENV HOST 0.0.0.0
ENV PORT 80

EXPOSE 4000
EXPOSE 80

ENTRYPOINT ["/app"]