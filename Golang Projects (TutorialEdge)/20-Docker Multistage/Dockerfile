FROM golang:1.12.1 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]

# docker build -t go-multi-stage .
# docker run -it -p 8080:8080 go-multi-stage